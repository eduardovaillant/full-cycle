package main

import (
	"fmt"	
	"github.com/eduardovaillant/full-cycle/application/route"
	"github.com/joho/godotenv"
)

func init()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	route := route.Route{
		ID: "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()
	fmt.Println(stringJson[0])
}