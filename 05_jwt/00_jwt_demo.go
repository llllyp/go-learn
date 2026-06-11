package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("12345678")

func main() {
	token, err := GenerateToken(1001, "java_dev")
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}

/**
 * 生成token
 */
func GenerateToken(userID int64, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(), // 过期时间
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

/**
 * 解析token
 */
func ParseToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})
}

/**
 * 读取 claims
 */
func GetClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := ParseToken(tokenStr)
	if err != nil || !token.Valid {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}

/**
 * 自定义 Claims
 */
type CustomClaims struct {
	UserID  int64  `json:"user_id"`
	Userame string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateCustomToken(UserID int64, username string) (string, error) {
	claims := CustomClaims{
		10086,
		"zhangsan",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "my-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(jwtSecret)
}

/**
 * 自定义 GIN 中间件, token验证
 */
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "missing token"})
		}

		strings.Replace(tokenStr, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "invaild token"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		ctx.Set("user_id", claims["user_id"])
		ctx.Next()
	}
}
