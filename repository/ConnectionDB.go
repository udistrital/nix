package repository

import (
	//"fmt"
	"database/sql"

    "nix/utilidades"
	"nix/model/tesoreria"

	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
)

func GetConnectionDB() *gorp.DbMap {

 	//fmt.Println("db :",utilidades.Parametros.User+":"+utilidades.Parametros.Password+"@"+utilidades.Parametros.Host+"/"+utilidades.Parametros.DataBaseName)
 	db, err := sql.Open("postgres", "postgres://"+utilidades.Parametros.User+":"+utilidades.Parametros.Password+"@"+utilidades.Parametros.Host+"/"+utilidades.Parametros.DataBaseName+"?sslmode=disable")
//	db, err := sql.Open("postgres", "postgres://"+User+":"+Password+"@"+Host+"/"+DataBaseName+"?sslmode=disable")
	utilidades.CheckErr(err, "No se logro realizar la conexion a la base de datos")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(tipoavance.Tipoavance{},"Tipoavance").SetKeys(true, "IdTipo")
	utilidades.CheckErr(err, "Create tables failed")

	return dbmap
}
