package main

import (
	"fmt"

	"github.com/henilthakor/bank_api/controller"
	"github.com/labstack/echo/v4"
)

func main() {

	fmt.Println("Welcome to the Bank API")

	e := echo.New()

	e.GET("/accountdetail/:id", controller.GetAccountDetail)

	e.POST("/createaccount", controller.CreateNewAccount)

	e.POST("/withdraw", controller.Withdraw)

	e.POST("/deposit", controller.Deposit)

	e.POST("/transfer", controller.Transfer)

	e.PUT("/updateaccountdetails", controller.UpdateAccountDetails)

	e.DELETE("/deleteaccount/:id", controller.DeleteAccount)

	err := e.Start(":8000")
	if err != nil {
		panic("Error occured in starting the server")
	}

}
