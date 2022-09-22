package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wewebplus/wewebapi/db"
	router "github.com/wewebplus/wewebapi/routes"
)

var server = db.Server{}
var route = router.Routes{}

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	route.Initialize()
	server.Run("api.localhost:8080")

}
