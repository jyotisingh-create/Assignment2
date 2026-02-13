package controllers

import (
	"banking-system/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	service *services.AccountService
}

func NewAccountController() *AccountController {
	return &AccountController{
		service: &services.AccountService{},
	}
}

func (ac *AccountController) OpenSavingsAccount(c *gin.Context) {
	var request struct {
		CustomerID uint `json:"customer_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := ac.service.OpenSavingsAccount(request.CustomerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, account)
}

func (ac *AccountController) GetAccount(c *gin.Context) {
	accountID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	account, err := ac.service.GetAccount(uint(accountID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (ac *AccountController) Deposit(c *gin.Context) {
	accountID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	var request struct {
		Amount float64 `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := ac.service.Deposit(uint(accountID), request.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (ac *AccountController) Withdraw(c *gin.Context) {
	accountID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	var request struct {
		Amount float64 `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := ac.service.Withdraw(uint(accountID), request.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (ac *AccountController) GetTransactions(c *gin.Context) {
	accountID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	transactions, err := ac.service.GetTransactions(uint(accountID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (ac *AccountController) GetBalance(c *gin.Context) {
	accountID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	balance, err := ac.service.GetBalance(uint(accountID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

func (ac *AccountController) AddAccountHolder(c *gin.Context) {
	accountID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	var request struct {
		CustomerID uint   `json:"customer_id" binding:"required"`
		HolderRole string `json:"holder_role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	holder, err := ac.service.AddAccountHolder(uint(accountID), request.CustomerID, request.HolderRole)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, holder)
}

func (ac *AccountController) RemoveAccountHolder(c *gin.Context) {
	accountID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	customerID, err := strconv.ParseUint(c.Param("customer_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	err = ac.service.RemoveAccountHolder(uint(accountID), uint(customerID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (ac *AccountController) GetAccountHolders(c *gin.Context) {
	accountID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	holders, err := ac.service.GetAccountHolders(uint(accountID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, holders)
}
