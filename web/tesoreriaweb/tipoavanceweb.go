package tipoavanceweb

import (
	//"fmt"
	"nix/model/tesoreria"
	"nix/repository/tesoreriarepository"
	"strconv"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine, middleware *jwt.GinJWTMiddleware) {

	apiTipoavance := router.Group("/tesoreria")
	//apiTipoavance.Use(middleware.MiddlewareFunc())

	apiTipoavance.GET("/tipoavance", List)
	apiTipoavance.GET("/tipoavance/:idtipo", FindOne)
	apiTipoavance.POST("/tipoavance", Create)
	apiTipoavance.PUT("/tipoavance/:idtipo", Modify)
	apiTipoavance.DELETE("/tipoavance/:idtipo", Delete)

	tipoavancerepository.Init()

}

func List(c *gin.Context) {

	tiposavance, msg := tipoavancerepository.FindAll()
	//consulta tipos de avance
	if msg.Code != 0 {
		c.JSON(200, msg)
	} else {
		c.JSON(200, tiposavance)
	}

}

func FindOne(c *gin.Context) {

	tipoavanceid, _ := strconv.ParseInt(c.Params.ByName("idtipo"), 0, 64)
	//fmt.Println("IDW :",tipoavanceid)
	tipoavance, msg := tipoavancerepository.FindOne(tipoavanceid)

	if msg.Code != 0 {
		c.JSON(200, msg)
	} else {
		c.JSON(200, tipoavance)
	}
}

func Create(c *gin.Context) {

	var tipoavanceins tipoavance.Tipoavance
	c.Bind(&tipoavanceins)
	msg := tipoavancerepository.Create(tipoavanceins)
	c.JSON(200, msg)
	
	/*
	tiposavance, msg := tipoavancerepository.Create(tipoavanceins)
	//consulta tipos de avance
	if msg.Code != 0 {
		c.JSON(200, msg)
	} else {
		c.JSON(200, tiposavance)
	}*/

}

func Modify(c *gin.Context) {

	var tipoavanceupd tipoavance.Tipoavance
	c.Bind(&tipoavanceupd)
	msg := tipoavancerepository.Update(tipoavanceupd)

	c.JSON(200, msg)
}

func Delete(c *gin.Context) {
	tipoavanceid, _ := strconv.ParseInt(c.Params.ByName("idtipo"), 0, 64)
	msg := tipoavancerepository.Delete(tipoavanceid)
	c.JSON(200, msg)
}
