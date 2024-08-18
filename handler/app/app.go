package app

import (
	"handler/api"
	"handler/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	
		// We actualy initialize our DB store here then pass it to the ApiHandler
		db := database.NewDynamoDBClient()
		apiHandler := api.NewApiHandler(db)

		return App {
			ApiHandler: apiHandler,
		}
	
}