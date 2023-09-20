package main

import (
	initializers "github.com/Zaul594/github.com/Zaul594/Super_Project/initializes"
	"github.com/Zaul594/github.com/Zaul594/Super_Project/models"
)

func init() {
	initializers.LoadEnvvariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Reservation{})
}
