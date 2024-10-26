package main

import(
	//"fmt"
	"test/route"
	"test/sql"
)

func main(){
	sql.InitSql()
	r := route.Router()
	r.Run(":8080")
}