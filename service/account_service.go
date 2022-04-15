package service

import (
	"errors"
	"fmt"

	"github.com/henilthakor/bank_api/domain"
)

type WithdrawData struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

type DepositData struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

type TransferData struct {
	FromID string  `json:"fromid"`
	ToID   string  `json:"toid"`
	Amount float64 `json:"amount"`
}

func GetAccountDetail(id string) (domain.Account, error) {
	account, err := domain.GetAccountDetail(id)
	return account, err

}

func CreateNewAccount(a *domain.Account) (string, error) {
	id, err := domain.AddAccount(a)
	return id.(string), err
}

func Withdraw(wd *WithdrawData) error {
	fmt.Println("IN WITHDRAW SERVICE")
	a, err := domain.GetAccountDetail(wd.ID)
	if err != nil {
		return err
	}
	if a.Balance >= wd.Amount {
		a.Balance = a.Balance - wd.Amount
		if err := domain.UpdateAccount(&a); err != nil {
			return err
		}
		return nil
	}
	return errors.New("Insufficient Balance")
}

func Deposit(dd *DepositData) error {
	a, err := domain.GetAccountDetail(dd.ID)
	if err != nil {
		return err
	}
	a.Balance = a.Balance + dd.Amount
	if err := domain.UpdateAccount(&a); err != nil {
		return err
	}
	return nil
}

func Transfer(td *TransferData) error {
	a1, err1 := domain.GetAccountDetail(td.FromID)
	if err1 != nil {
		return err1
	}

	a2, err2 := domain.GetAccountDetail(td.ToID)
	if err2 != nil {
		return err2
	}

	if a1.Balance >= td.Amount {
		a1.Balance = a1.Balance - td.Amount
		a2.Balance = a2.Balance + td.Amount
		if err := domain.UpdateAccount(&a1); err != nil {
			return err
		}
		if err := domain.UpdateAccount(&a2); err != nil {
			return err
		}
		return nil
	}
	return errors.New("Insufficient Balance")
}

func UpdateAccountDetails(a *domain.Account) error {
	return domain.UpdateAccount(a)
}

func DeleteAccount(id string) error {
	return domain.DeleteAccount(id)
}
