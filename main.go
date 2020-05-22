package main

import (
	"dat2/conf"
	"dat2/db"
	"dat2/router"
)

func main() {
	r := router.InitRouter()
	r.Run(":" + conf.GetConfig().Sys.Port)
	db.Close()
}
