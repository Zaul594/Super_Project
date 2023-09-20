package main

import (
	"net/http"
	//"github.com/gofiber/fiber/v2"

	"github.com/Zaul594/github.com/Zaul594/Super_Project/controllers"
	initializers "github.com/Zaul594/github.com/Zaul594/Super_Project/initializes"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func init() {
	initializers.LoadEnvvariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcom to Z's Car Wash",
		})
	})

	r.GET("/price", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "a normal car wash is 20$",
		})
	})

	rGroup := r.Group("/rsv")
	rGroup.GET("", controllers.ReservationsIdnex)
	rGroup.GET("/:id", controllers.GetReservation)
	rGroup.PUT("/:id", controllers.PostUpdate)
	rGroup.POST("", controllers.CreateReservations)
	rGroup.DELETE("/:id", controllers.DeleteReservation)

	r.Run()
}
