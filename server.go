package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"echo-framework/article"
	"echo-framework/author"
	Config "echo-framework/config"

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
	ioutil.WriteFile("routes.json", data, 0644)
	e.Logger.Fatal(e.Start(":1323"))
}
