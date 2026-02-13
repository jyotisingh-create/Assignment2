package controllers

import (
	"banking-system/config"
	"banking-system/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CustomerController struct{}

func (cc *CustomerController) CreateCustomer(c *gin.Context) {
	var request struct {
		BranchID uint   `json:"branch_id" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Phone    string `json:"phone" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var branch models.Branch
	if err := config.DB.First(&branch, request.BranchID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Branch not found"})
		return
	}

	customer := &models.Customer{
		BranchID:  request.BranchID,
		Name:      request.Name,
		Email:     request.Email,
		Phone:     request.Phone,
		CreatedAt: time.Now().Unix(),
	}

	if err := config.DB.Create(customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func (cc *CustomerController) GetCustomer(c *gin.Context) {
	customerID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	var customer models.Customer
	if err := config.DB.Preload("CustomerAccounts.Account").Preload("Loans").First(&customer, uint(customerID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}
