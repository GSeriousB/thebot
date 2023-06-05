package jwt

// type IJwtService interface {
// 	CreateNewTokens(ctx context.Context, pos *dbmodels.POS) (*TokenDetails, error)
// 	VerifyToken(ctx context.Context, tokenString string) (*dbmodels.POS, bool)
// }

// type JwtService struct {
// }

// func NewJwtService() *JwtService {
// 	return &JwtService{}
// }

// type TokenDetails struct {
// 	AccessToken  string
// 	RefreshToken string
// 	AccessUuid   string
// 	RefreshUuid  string
// 	AtExpires    int64
// 	RtExpires    int64
// }

// func (j *JwtService) CreateNewTokens(ctx context.Context, pos *dbmodels.POS) (*TokenDetails, error) {
// 	log := logger.Logger(ctx)
// 	log.Infof("Creating token for ", pos)

// 	td := &TokenDetails{}
// 	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
// 	newuuid, err := uuid.NewUUID()
// 	if err != nil {
// 		log.Errorf("error while generating access id ", err)
// 		return nil, fmt.Errorf("unable to generate access uuid")
// 	}
// 	td.AccessUuid = newuuid.String()
// 	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
// 	td.RefreshUuid = td.AccessUuid + "++" + pos.ID.String()

// 	atClaims := jwt.MapClaims{}
// 	atClaims["authorized"] = true
// 	atClaims["access_uuid"] = td.AccessUuid
// 	atClaims["pos_id"] = pos.ID
// 	atClaims["pos"] = pos
// 	atClaims["exp"] = td.AtExpires
// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	td.AccessToken, err = at.SignedString([]byte(os.Getenv("tradebot_ACCESS_SECRET")))
// 	if err != nil {
// 		return nil, err
// 	}

// 	//Creating Refresh Token
// 	rtClaims := jwt.MapClaims{}
// 	atClaims["refresh_uuid"] = td.RefreshUuid
// 	atClaims["pos_id"] = pos.ID
// 	atClaims["pos"] = pos
// 	atClaims["exp"] = td.RtExpires

// 	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
// 	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("tradebot_ACCESS_SECRET")))
// 	if err != nil {
// 		log.Errorf("error while generating access id ", err)
// 		return nil, err
// 	}
// 	return td, nil
// }

// func (j *JwtService) VerifyToken(ctx context.Context, tokenString string) (*dbmodels.POS, bool) {

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("JUBI_ACCESS_SECRET")), nil
// 	})
// 	if err != nil {
// 		return nil, false
// 	}

// 	var pos dbmodels.POS
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		pos.ID = claims["pos_id"].(uuid.UUID)
// 		pos.POS = claims["pos"].(string)
// 	}
// 	return nil, false

// }
