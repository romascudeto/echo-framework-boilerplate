package main

import (
	"echo-framework/article"
	"echo-framework/author"
	Config "echo-framework/config"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var err error

func main() {
	godotenv.Load()
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}
	e := echo.New()
	article.Routes(e)
	author.Routes(e)
	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	log.Println(string(data))
	e.Logger.Fatal(e.Start(":1323"))
}
