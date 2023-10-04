package service

import (
	"errors"
	"log"

	"github.com/omerberkcan/banking-transfer/dto"
	"github.com/omerberkcan/banking-transfer/internal/repository"
	"github.com/omerberkcan/banking-transfer/model"
	"github.com/shopspring/decimal"
)

var (
	errInsufficientBalance = errors.New("insufficient balance in your wallet for this transfer")
	errWrongIDNumber       = errors.New("wrong ID Number")
	errSomethingWentWrong  = errors.New("something went wrong")
	errRecordNotFound      = errors.New("user not found")
	errInvalidTransfer     = errors.New("you cannot send money yourself")
)

type (
	transferService struct {
		store repository.Stores
	}

	TransferService interface {
		MoneyTransfer(id uint, t dto.TransferDTO) error
	}
)

func (ts transferService) MoneyTransfer(id uint, t dto.TransferDTO) error {
	originUsr, err := ts.store.Users().FindByID(id)
	if err != nil {
		return errRecordNotFound
	}

	if originUsr.Balance.Cmp(t.Amount) == -1 {
		return errInsufficientBalance
	}

	destUsr, err := ts.store.Users().FindByIDNo(t.IDNo)
	if err != nil {
		return errWrongIDNumber
	}

	if originUsr.IdNo == destUsr.IdNo {
		return errInvalidTransfer
	}

	db := ts.store.TxBegin()

	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	onePercent := decimal.NewFromFloat(0.01)
	tax := t.Amount.Mul(onePercent)

	originLastBalance := originUsr.Balance.Sub(t.Amount).Sub(tax)
	err = ts.store.Users().WithTrx(db).UpdateBalance(id, originLastBalance)
	if err != nil {
		log.Print("update balance error: ", err)
		db.Rollback()
		return errSomethingWentWrong
	}

	destLastBalance := destUsr.Balance.Add(t.Amount)
	err = ts.store.Users().WithTrx(db).UpdateBalance(destUsr.ID, destLastBalance)
	if err != nil {
		log.Print("update balance error: ", err)
		db.Rollback()
		return errSomethingWentWrong
	}

	transfer := &model.Transfer{
		UserOriginID:      originUsr.ID,
		UserDestinationID: destUsr.ID,
		Amount:            t.Amount,
		Description:       t.Description,
		Tax:               tax,
	}

	err = ts.store.Transfer().WithTrx(db).Create(transfer)
	if err != nil {
		db.Rollback()
		log.Print("create transfer error: ", err)
		return errSomethingWentWrong
	}

	if err := db.Commit().Error; err != nil {
		log.Print("trx commit error: ", err)
		return errSomethingWentWrong
	}
	return nil
}
