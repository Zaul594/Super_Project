package controllers

import (
	initializers "github.com/Zaul594/github.com/Zaul594/Super_Project/initializes"
	"github.com/Zaul594/github.com/Zaul594/Super_Project/models"
	"github.com/gin-gonic/gin"
)

func CreateReservations(c *gin.Context) {
	//Get data of req body
	var reservations struct {
		Name string
		Type string
	}

	c.Bind(&reservations)

	//Create a reservation
	reservation := models.Reservation{
		Name: reservations.Name,
		Type: reservations.Type,
	}
	result := initializers.DB.Create(&reservation)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(208, gin.H{
		"message": reservation,
	})
}

func ReservationsIdnex(c *gin.Context) {
	//Get the posts
	var getreservations []models.Reservation
	initializers.DB.Find(&getreservations)

	//Respond with them
	c.JSON(208, gin.H{
		"reservations": getreservations,
	})
}

func GetReservation(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	//Get the posts
	var getreservation models.Reservation
	initializers.DB.First(&getreservation, id)

	//Respond with them
	c.JSON(208, gin.H{
		"reservation": getreservation,
	})
}

func PostUpdate(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	//Get the data off req body

	var reservations struct {
		Name string
		Type string
	}

	c.Bind(&reservations)

	//find the post bing updated
	var getreservation models.Reservation
	initializers.DB.First(&getreservation, id)

	//Update post
	initializers.DB.Model(&getreservation).Updates(models.Reservation{
		Name: reservations.Name,
		Type: reservations.Type,
	})

	//Respond with them
	c.JSON(208, gin.H{
		"reservation": getreservation,
	})
}

func DeleteReservation(c *gin.Context) {
	//Get id of url
	id := c.Param("id")

	//Delete the post
	initializers.DB.Delete(&models.Reservation{}, id)

	//respond
	c.Status(200)
}
