package controllers

import (
	"banking-system/config"
	"banking-system/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BankController struct{}

func (bc *BankController) CreateBank(c *gin.Context) {
	var request struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bank := &models.Bank{
		Name:      request.Name,
		CreatedAt: time.Now().Unix(),
	}

	if err := config.DB.Create(bank).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, bank)
}

func (bc *BankController) GetBank(c *gin.Context) {
	bankID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bank ID"})
		return
	}

	var bank models.Bank
	if err := config.DB.Preload("Branches").First(&bank, uint(bankID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bank not found"})
		return
	}

	c.JSON(http.StatusOK, bank)
}

type BranchController struct{}

func (bc *BranchController) CreateBranch(c *gin.Context) {
	var request struct {
		BankID  uint   `json:"bank_id" binding:"required"`
		Name    string `json:"name" binding:"required"`
		Address string `json:"address" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if bank exists
	var bank models.Bank
	if err := config.DB.First(&bank, request.BankID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bank not found"})
		return
	}

	branch := &models.Branch{
		BankID:    request.BankID,
		Name:      request.Name,
		Address:   request.Address,
		CreatedAt: time.Now().Unix(),
	}

	if err := config.DB.Create(branch).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, branch)
}

func (bc *BranchController) GetBranch(c *gin.Context) {
	branchID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid branch ID"})
		return
	}

	var branch models.Branch
	if err := config.DB.Preload("Customers").First(&branch, uint(branchID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Branch not found"})
		return
	}

	c.JSON(http.StatusOK, branch)
}
