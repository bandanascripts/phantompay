package controllers

import (
	"net/http"

	"github.com/bandanascripts/phantompay/pkg/client/auth"
	"github.com/bandanascripts/phantompay/pkg/client/token"
	"github.com/bandanascripts/phantompay/pkg/service/database"
	"github.com/bandanascripts/phantompay/pkg/service/middleware"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {

	var loginData auth.LoginRequest

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, userId, err := auth.Login(loginData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken, _, err := token.GenerateTokens(c.Request.Context(), userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message, "accesstoken": accessToken})

}

func RegisterHandler(c *gin.Context) {

	var registerData auth.RegisterRequest

	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := auth.Register(registerData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": message})
}

func DepositHandler(c *gin.Context) {

	var deposit database.DepositRequest

	if err := c.ShouldBindJSON(&deposit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := middleware.ExtractUserId(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := database.DepositMoney(userId, deposit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}

func CreateWalletHandler(c *gin.Context) {

	userId, err := middleware.ExtractUserId(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	message, walletId, err := database.InitWallet(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": message, "walletid": walletId})
}

func WithdrawMoneyHandler(c *gin.Context) {

	var withdraw database.WithdrawRequest

	if err := c.ShouldBindJSON(&withdraw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := middleware.ExtractUserId(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := database.WithdrawMoney(userId, withdraw)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}

func TransactionHandler(c *gin.Context) {

	var transaction database.TransactionRequest

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := middleware.ExtractUserId(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction.UserId = userId

	message, err := database.Transaction(transaction)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}

func TransactionHistoryHandler(c *gin.Context) {

	userId, err := middleware.ExtractUserId(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactions, err := database.TransactionHistory(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": transactions})
}
