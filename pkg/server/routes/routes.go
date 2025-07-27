package routes

import (
	"github.com/bandanascripts/phantompay/pkg/server/controllers"
	"github.com/bandanascripts/phantompay/pkg/service/middleware"
	"github.com/gin-gonic/gin"
)

func RegisteredRoutes(c *gin.Engine) {

	c.POST("/phantompay/login", controllers.LoginHandler)
	c.POST("/phantompay/register", controllers.RegisterHandler)
	c.POST("/phantompay/deposit", middleware.TokenMiddleware(), controllers.DepositHandler)
	c.POST("/phantompay/withdraw", middleware.TokenMiddleware(), controllers.WithdrawMoneyHandler)
	c.POST("/phantompay/transaction", middleware.TokenMiddleware(), controllers.TransactionHandler)

	c.GET("phantompay/wallet", middleware.TokenMiddleware(), controllers.CreateWalletHandler)
	c.GET("/phantompay/transactionhistory", middleware.TokenMiddleware(), controllers.TransactionHistoryHandler)
}
