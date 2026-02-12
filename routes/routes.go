package routes

import (
	"banking-system/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	bankCtrl := &controllers.BankController{}
	router.POST("/banks", bankCtrl.CreateBank)
	router.GET("/banks/:id", bankCtrl.GetBank)

	branchCtrl := &controllers.BranchController{}
	router.POST("/branches", branchCtrl.CreateBranch)
	router.GET("/branches/:id", branchCtrl.GetBranch)

	customerCtrl := &controllers.CustomerController{}
	router.POST("/customers", customerCtrl.CreateCustomer)
	router.GET("/customers/:id", customerCtrl.GetCustomer)

	accountCtrl := controllers.NewAccountController()
	router.POST("/accounts/savings", accountCtrl.OpenSavingsAccount)
	router.GET("/accounts/:id", accountCtrl.GetAccount)
	router.POST("/accounts/:id/deposit", accountCtrl.Deposit)
	router.POST("/accounts/:id/withdraw", accountCtrl.Withdraw)
	router.GET("/accounts/:id/transactions", accountCtrl.GetTransactions)
	router.GET("/accounts/:id/balance", accountCtrl.GetBalance)
	router.POST("/accounts/:id/holders", accountCtrl.AddAccountHolder)
	router.DELETE("/accounts/:id/holders/:customer_id", accountCtrl.RemoveAccountHolder)
	router.GET("/accounts/:id/holders", accountCtrl.GetAccountHolders)

	loanCtrl := controllers.NewLoanController()
	router.POST("/loans", loanCtrl.TakeLoan)
	router.GET("/loans/:id", loanCtrl.GetLoan)
	router.POST("/loans/:id/repay", loanCtrl.RepayLoan)
	router.GET("/loans/:id/details", loanCtrl.GetLoanDetails)
	router.GET("/loans/:id/interest", loanCtrl.GetInterest)
}
