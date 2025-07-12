package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/B-Bridger/server/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 인증 middleware 구현
// 인증 성공 시, context에 userID 키에 UserID 값을 저장
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{Message: "토큰이 만료되었습니다", Detail: "token is empty", Status: 401})
			return
		}
		splitToken := strings.SplitN(tokenString, "Bearer ", 2)
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{Message: "토큰 형식이 올바르지 않습니다", Detail: "invalid token format", Status: 401})
			return
		}
		auth := strings.TrimSpace(splitToken[1])

		token, err := jwt.ParseWithClaims(auth, &model.BridgerClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, model.ErrorResponse{Message: "접근 권한이 없습니다", Detail: err.Error(), Status: 403})
			return
		}

		if claims, ok := token.Claims.(*model.BridgerClaims); ok && token.Valid {
			c.Set("userID", claims.UserID)
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, model.ErrorResponse{Message: "토큰이 유효하지 않습니다", Detail: "invalid claims", Status: 403})
			return
		}
		c.Next()
	}
}
