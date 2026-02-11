package controllers

import (
	"banking-system/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoanController struct {
	service *services.LoanService
}

func NewLoanController() *LoanController {
	return &LoanController{
		service: &services.LoanService{},
	}
}

func (lc *LoanController) TakeLoan(c *gin.Context) {
	var request struct {
		CustomerID      uint    `json:"customer_id" binding:"required"`
		PrincipalAmount float64 `json:"principal_amount" binding:"required"`
		LoanType        string  `json:"loan_type" binding:"required"`
		DurationYears   float64 `json:"duration_years" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loan, err := lc.service.TakeLoan(request.CustomerID, request.PrincipalAmount, request.LoanType, request.DurationYears)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, loan)
}

func (lc *LoanController) GetLoan(c *gin.Context) {
	loanID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	loan, err := lc.service.GetLoan(uint(loanID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loan)
}

func (lc *LoanController) RepayLoan(c *gin.Context) {
	loanID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	var request struct {
		Amount float64 `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loan, err := lc.service.RepayLoan(uint(loanID), request.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loan)
}

func (lc *LoanController) GetLoanDetails(c *gin.Context) {
	loanID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	details, err := lc.service.GetLoanDetails(uint(loanID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, details)
}

func (lc *LoanController) GetInterest(c *gin.Context) {
	loanID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	interest, err := lc.service.CalculateInterest(uint(loanID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"yearly_interest": interest})
}
