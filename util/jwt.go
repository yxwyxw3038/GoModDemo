package util

import (
    "github.com/dgrijalva/jwt-go"
    "time"
    "GoModDemo/setting"
)


type Claims struct {
    Username string `json:"username"`
    Password string `json:"password"`
    jwt.StandardClaims
}

func GenerateToken(username string, password string) (string, error) {
    var jwtSecret = []byte(setting.AppSetting.JwtSecret)
    nowTime := time.Now()
    expireTime := nowTime.Add(3 * time.Hour)
    claims := Claims{
        username,
        password,
        jwt.StandardClaims {
            ExpiresAt : expireTime.Unix(),
            Issuer : "go-backend-starter",
        },
    }
    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token, err := tokenClaims.SignedString(jwtSecret)

    return token, err
}

func ParseToken(token string) (*Claims, error) {
    var jwtSecret = []byte(setting.AppSetting.JwtSecret)
    tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            return claims, nil
        }
    }

    return nil, err
}