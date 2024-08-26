package middleware

import (
    "net/http"
    "os"

    jwtmiddleware "github.com/auth0/go-jwt-middleware"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

// JWTMiddleware function to verify the token
func JWTMiddleware() gin.HandlerFunc {
    jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
        ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("AUTH0_CLIENT_SECRET")), nil
        },
        SigningMethod: jwt.SigningMethodHS256,
    })

    return func(c *gin.Context) {
        if err := jwtMiddleware.CheckJWT(c.Writer, c.Request); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        c.Next()
    }
}
