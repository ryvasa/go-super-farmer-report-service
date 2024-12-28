package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/logrus"
	wire_excel "github.com/ryvasa/go-super-farmer-report-service/pkg/wire"
)

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.Run() // listen and serve on 0.0.0.0:8080
// }

func main() {
	logrus.Log.Info("Starting Report service...")
	app, err := wire_excel.InitializeReportApp()
	if err != nil {
		log.Fatal(err)
		logrus.Log.Fatalf("failed to initialize app: %v", err)
	}

	go app.Handler.ReportHandler.ConsumerHandler()

	defer app.RabbitMQ.Close()
	app.Router.Use(gin.Recovery())
	app.Router.Use(gin.Logger())
	app.Router.Run(":" + app.Env.Report.Port)
}
