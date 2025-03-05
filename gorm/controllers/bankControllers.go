package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leebrouse/gorm/models"
)

type IBankControllers interface {
	Increase(c *gin.Context)
	// Decease(c *gin.Context)
}

type BankControllers struct {
	BaseControllers
}

func NewBankControllers() IBankControllers {
	return &BankControllers{}
}

func (bank *BankControllers) Increase(c *gin.Context) {
	//begin
	tx := models.DB.Begin()
	//defer recover pain if have err,rollback
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			bank.Fail(c)
		}
	}()
	//todo

	u1 := models.Bank{ID: 1}
	tx.Find(&u1)
	u1.Balance = u1.Balance - 100
	if err := tx.Save(&u1).Error; err != nil {
		tx.Rollback()
	}

	u2 := models.Bank{ID: 2}
	tx.Find(&u2)
	u2.Balance = u2.Balance + 100
	if err := tx.Save(&u2).Error; err != nil {
		tx.Rollback()
		bank.Fail(c)
	}

	//defer commit
	tx.Commit()
	
	bank.Success(c)
}
