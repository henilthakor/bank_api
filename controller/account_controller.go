package controller

import (
	"encoding/json"
	"net/http"

	"github.com/henilthakor/bank_api/domain"
	"github.com/henilthakor/bank_api/service"
	"github.com/labstack/echo/v4"
)

func GetAccountDetail(ctx echo.Context) error {
	id := ctx.Param("id")

	a, err := service.GetAccountDetail(id)
	if err != nil {
		errmap := map[string]string{"error": err.Error()}
		errjson, _ := json.Marshal(errmap)
		return ctx.JSON(http.StatusInternalServerError, errjson)
	}

	return ctx.JSON(http.StatusOK, a)
}

func CreateNewAccount(ctx echo.Context) error {
	var a = new(domain.Account)

	if err := ctx.Bind(a); err != nil {
		errmap := map[string]string{"error": "Bad Request"}
		errjson, _ := json.Marshal(errmap)
		return ctx.JSON(http.StatusBadRequest, errjson)
	}

	id, err := service.CreateNewAccount(a)
	if err != nil {
		errmap := map[string]string{"error": err.Error()}
		errjson, _ := json.Marshal(errmap)
		return ctx.JSON(http.StatusInternalServerError, errjson)
	}

	response := map[string]string{"_id": id}
	responseJson, _ := json.Marshal(response)
	return ctx.JSON(http.StatusOK, responseJson)
}

func Withdraw(ctx echo.Context) error {
	var withdrawData = new(service.WithdrawData)

	if err := ctx.Bind(withdrawData); err != nil {
		errmap := map[string]string{"error": "Bad Request"}
		errjson, _ := json.Marshal(errmap)
		return ctx.JSON(http.StatusBadRequest, errjson)
	}

	if err := service.Withdraw(withdrawData); err != nil {
		errmap := map[string]string{"error": err.Error()}
		errjson, _ := json.Marshal(errmap)
		return ctx.JSON(http.StatusInternalServerError, errjson)
	}

	return ctx.JSON(http.StatusOK, withdrawData)

}

func Deposit(ctx echo.Context) error {
	var depositData = new(service.DepositData)

	if err := ctx.Bind(depositData); err != nil {
		errmap := map[string]string{"error": "Bad Request"}
		errjson, _ := json.Marshal(errmap)
		return ctx.JSON(http.StatusBadRequest, errjson)
	}

	if err := service.Deposit(depositData); err != nil {
		errmap := map[string]string{"error": err.Error()}
		errjson, _ := json.Marshal(errmap)
		return ctx.JSON(http.StatusInternalServerError, errjson)
	}

	return ctx.JSON(http.StatusOK, depositData)
}

func Transfer(ctx echo.Context) error {
	var transferData = new(service.TransferData)

	if err := ctx.Bind(transferData); err != nil {
		errmap := map[string]string{"error": "Bad Request"}
		errjson, _ := json.Marshal(errmap)
		return ctx.JSON(http.StatusBadRequest, errjson)
	}

	if err := service.Transfer(transferData); err != nil {
		errmap := map[string]string{"error": err.Error()}
		errjson, _ := json.Marshal(errmap)
		return ctx.JSON(http.StatusInternalServerError, errjson)
	}

	return ctx.JSON(http.StatusOK, transferData)
}

func UpdateAccountDetails(ctx echo.Context) error {
	var a = new(domain.Account)

	if err := ctx.Bind(a); err != nil {
		errmap := map[string]string{"error": "Bad Request"}
		errjson, _ := json.Marshal(errmap)
		return ctx.JSON(http.StatusBadRequest, errjson)
	}

	err := service.UpdateAccountDetails(a)
	if err != nil {
		errmap := map[string]string{"error": err.Error()}
		errjson, _ := json.Marshal(errmap)
		return ctx.JSON(http.StatusInternalServerError, errjson)
	}

	return ctx.JSON(http.StatusOK, a)
}

func DeleteAccount(ctx echo.Context) error {
	id := ctx.Param("id")

	err := service.DeleteAccount(id)
	if err != nil {
		errmap := map[string]string{"error": err.Error()}
		errjson, _ := json.Marshal(errmap)
		return ctx.JSON(http.StatusInternalServerError, errjson)
	}

	return ctx.JSON(http.StatusOK, nil)
}
