package main

import (
	database "myapp/database"
	router "myapp/router"
)


func main() {
database.Connect()
defer database.Close()
routes := router.Router();
routes.Logger.Fatal(routes.Start(":1323"))
}

