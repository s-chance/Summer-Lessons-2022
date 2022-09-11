package main

import (
	"demo/config"
	"demo/db"
	"demo/db/model"
	"demo/router"
)

func main() {
	config.InitConfig()
	db.InitDB()
	model.InitModel()
	router.InitRouter()
	router.Router.Run(":" + config.Config.Server.Port)
}
