package main

import (
	router "myapp/router"
)

func main() {
routes := router.Router();
routes.Logger.Fatal(routes.Start(":1323"))
}

