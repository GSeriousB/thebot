package apiclient

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"tradebot/bybit/app/service/logger"

	"github.com/spf13/viper"
)

type IApiClient interface {
	CreateJSONRequest(ctx context.Context, httpMethod, absoluteURL string, headers map[string]string, requestBody interface{}) (*http.Request, error)
	RestExecute(ctx context.Context, request *http.Request) (int, string, error)
}

type ApiClient struct{}

func NewApiClient() IApiClient {
	return &ApiClient{}
}

// RestExecute : An abstraction to call a rest API
func (f *ApiClient) RestExecute(ctx context.Context, request *http.Request) (int, string, error) {
	log := logger.Logger(ctx)
	restExecuteTimeoutInSeconds := viper.GetString("RestExecuteTimeoutInSeconds")
	restExecuteTimeoutInSecondsInt, _ := strconv.ParseInt(restExecuteTimeoutInSeconds, 6, 12)
	client := &http.Client{Timeout: time.Duration(restExecuteTimeoutInSecondsInt) * time.Second}
	requestURL := request.URL.String()
	requestMethod := request.Method

	log.Infow("Invoking API", "URL", requestURL, "Method", requestMethod)

	requestedTime := time.Now()
	response, err := client.Do(request)
	if err != nil || response == nil {
		log.Errorw("Error or no response while invoking API", "URL", requestURL, "Method", requestMethod, "Error", err)
		return http.StatusInternalServerError, "", err
	}

	responseTime := time.Since(requestedTime)
	log.Infow("Received response from Resource URL", "ResourceURL", requestURL, "Method", requestMethod, "Status", response.Status, "ResponseTime", responseTime)

	defer response.Body.Close()

	if response.StatusCode == 401 {
		log.Warnf("Unauthorized resource", request)
		return response.StatusCode, "", errors.New("unauthorized resource")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Warn("Unable to read response body")
		return response.StatusCode, "", errors.New("unable to read response body")
	}

	responseString := string(responseData)
	log.Debugw("Received response", "Body", responseString)
	return response.StatusCode, responseString, nil
}

// CreateJSONRequest : Creates a JSON HTTP request
func (f *ApiClient) CreateJSONRequest(ctx context.Context, httpMethod, absoluteURL string, headers map[string]string, requestBody interface{}) (*http.Request, error) {
	log := logger.Logger(ctx)
	u, urlerr := url.ParseRequestURI(absoluteURL)
	if urlerr != nil {
		log.Warnf("Invalid api request path. ResourcePath: %s", absoluteURL)
		return nil, urlerr
	}

	httpMethod = strings.ToUpper(httpMethod)

	requestJSONBody := ""
	if httpMethod != "GET" && requestBody != nil {
		jsonValue, merr := json.Marshal(requestBody)
		if merr != nil {
			log.Errorf("Unable to marshal request body. RequestURL: %s Method: %s", u.String(), httpMethod)
			return nil, errors.New("unable to marshal request body")
		}
		requestJSONBody = string(jsonValue)
	}

	request, requestError := http.NewRequest(httpMethod, u.String(), strings.NewReader(requestJSONBody))

	if requestError != nil {
		log.Errorf("Unable to create request. URL: %s ErrorMessage: %s", u.String(), requestError.Error())
		return nil, requestError
	}

	for k, v := range headers {
		request.Header.Add(k, v)
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Length", strconv.Itoa(len(requestJSONBody)))

	return request, nil
}
