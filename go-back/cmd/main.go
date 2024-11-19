package main

import (
	"project-hackathon/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {

	app, err := bootstrap.App()
	if err != nil {
		logrus.Fatal(err)
	}

	env := app.Env
	db := app.Database
	logger := app.Logger

	gin := gin.Default()
	timeout := time.Duration(5) * time.Second
}
