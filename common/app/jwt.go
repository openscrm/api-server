package app

//
//import (
//	"time"
//
//	"github.com/go-programming-tour-book/blog-service/pkg/util"
//
//	"github.com/dgrijalva/jwt-go"
//	"github.com/go-programming-tour-book/blog-service/global"
//)
//
//type Claims struct {
//	AppKey    string `json:"app_key"`
//	AppSecret string `json:"app_secret"`
//	jwt.StandardClaims
//}
//
//func GetJWTSecret() []byte {
//	return []byte(global.JWTSetting.Secret)
//}
//
//func GenerateToken(appKey, appSecret string) (string, error) {
//	nowTime := time.Now()
//	expireTime := nowTime.Add(global.JWTSetting.Expire)
//	claims := Claims{
//		AppKey:    util.EncodeMD5(appKey),
//		AppSecret: util.EncodeMD5(appSecret),
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expireTime.Seconds(),
//			Issuer:    global.JWTSetting.Issuer,
//		},
//	}
//
//	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	token, err := tokenClaims.SignedString(GetJWTSecret())
//	return token, err
//}
//
//func ParseToken(token string) (*Claims, error) {
//	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return GetJWTSecret(), nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	if tokenClaims != nil {
//		claims, ok := tokenClaims.Claims.(*Claims)
//		if ok && tokenClaims.Valid {
//			return claims, nil
//		}
//	}
//
//	return nil, err
//}
