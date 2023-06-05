package auth

// //prefix can be changed w.r.t application requirements
// var AuthPrefix = "Bearer "
// var TokenKey = "Authorization"

// // authentication is a middleware that verify JWT token headers
// func Authentication(next gin.HandlerFunc, jwt jwt.IJwtService) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		log := logger.Logger(ctx)

// 		token, err := getHeaderToken(ctx)
// 		if err != nil {
// 			controller.ResponseMethod(ctx, http.StatusUnauthorized, err.Error())
// 			return
// 		}

// 		claims, valid := jwt.VerifyToken(ctx, token)
// 		if !valid {
// 			controller.ResponseMethod(ctx, http.StatusUnauthorized, "unauthorized")
// 			return
// 		}
// 		log.Info("pos", zap.String("pos ", claims.POS))
// 		ctx.Set("pos", claims)
// 		next(ctx)
// 	}
// }

// func getHeaderToken(ctx *gin.Context) (string, error) {
// 	header := string(ctx.GetHeader(TokenKey))
// 	return extractToken(header)
// }

// func extractToken(header string) (string, error) {
// 	if strings.HasPrefix(header, AuthPrefix) {
// 		return header[len(AuthPrefix):], nil
// 	}
// 	return "", errors.New("token not found")
// }
