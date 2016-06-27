package security

import(
"github.com/appleboy/gin-jwt"
"github.com/gin-gonic/gin"
"time"
)


var MiddlewareAUTH *jwt.GinJWTMiddleware

func Init(router *gin.Engine) {

	// the jwt middleware
	MiddlewareAUTH = &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour * 24,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {

		// gestion de autenticacion
		  if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
		    return userId, true
		  }

		  return userId, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
		  if userId == "admin" {
		    return true
		  }

		  return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
		  c.JSON(code, gin.H{
		    "code":    code,
		    "message": message,
		  })
		},
	}

	router.POST("/login", MiddlewareAUTH.LoginHandler)

}