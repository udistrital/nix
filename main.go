package main

import (
	"nix/utilidades"

	"nix/web/security"
	"nix/web/tesoreriaweb"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {

	//carga inicial de parametros de configuracion
	utilidades.Init()

	//Inicializacion de framework Gin-gonic
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(Cors())

	//inicializacion de middleware de AUTH
	security.Init(r)

	//inicializacion de todos los routes del aplicativo
	tipoavanceweb.Init(r, security.MiddlewareAUTH)




	//inicia el servidor
	endless.ListenAndServe(":8000", r)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
