package app

import (
	"github.com/dgrijalva/jwt-go"
	"go-programming-tour-book-exercise/chapter-2-blog-service/global"
	"go-programming-tour-book-exercise/chapter-2-blog-service/pkg/util"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			Issuer:    global.JWTSetting.Issuer,
			ExpiresAt: expireTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		valid := tokenClaims.Valid
		if ok && valid {
			return claims, nil
		}
	}

	return nil, err
}

// NOTE: must not store important info in token, because it can be decoded by others.
// payload, err := base64.StdEncoding.DecodeString("eyJhcHBfa2V5IjoiZGQ4MWZjZTU2ZDg3ZjhjMjhkMzg5ZmJiM2M2Y2IxZTkiLCJhcHBfc2VjcmV0IjoiN2M5NzI2NjMxNzBkNmJjMTg0ODRkMDViYzk4NzIyZjQiLCJleHAiOjE2MjE2NjM3MjMsImlzcyI6ImJsb2ctc2VydmljZSJ9")
// if err != nil {
// log.Fatalf("jwt decode err: %v", err)
// }
//
// fmt.Printf("payload: %v", string(payload))
// // payload: {"app_key":"dd81fce56d87f8c28d389fbb3c6cb1e9","app_secret":"7c972663170d6bc18484d05bc98722f4","exp":1621663723,"iss":"blog-service"}
