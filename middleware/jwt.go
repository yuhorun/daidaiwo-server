package middleware

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/common"
	"server/pkg/e"
	"server/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Uid int64 `json:"uid"`
	jwt.StandardClaims
}

func GenerateToken(id int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(360 * 24 * time.Hour)

	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
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

func ParseTokenFromCookie(c *gin.Context, cookieName string) (claim *Claims, err error) {
	value, err := c.Cookie(cookieName)
	if err != nil {
		return nil, err
	}

	claim, err = ParseToken(value)
	if err != nil {
		return nil, err
	}
	return
}

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {

		jwtString, err := c.Cookie("jwtstring")

		if err != nil {
			common.CJSON(c, http.StatusOK, e.ERROR_COOKIE_NOT_SET)
			c.Abort()
			return
		}

		claims, err := ParseToken(jwtString)
		if err != nil {
			common.CJSON(c, http.StatusOK, e.ERROR_AUTH_CHECK_JWT_FAIL)
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			common.CJSON(c, http.StatusOK, e.ERROR_AUTH_CHECK_JWT_TIMEOUT)
			c.Abort()
			return
		}

		cmap := make(map[string]interface{})
		cmap["uid"] = claims.Uid

		c.Keys = cmap //字典为指针赋值  浅拷贝

		c.Next()
	}
}
