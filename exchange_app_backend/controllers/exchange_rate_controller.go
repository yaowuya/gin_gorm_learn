package controllers

import (
	"errors"
	"exchange_app/global"
	"exchange_app/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func CreateExchangeRate(ctx *gin.Context) {
	var exchangeRate models.ExchangeRate
	if err := ctx.ShouldBindJSON(&exchangeRate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exchangeRate.Date = time.Now()
	if err := global.Db.AutoMigrate(&exchangeRate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.Create(&exchangeRate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, exchangeRate)
}

func GetExchangeRates(ctx *gin.Context) {
	fromCurrency := ctx.Query("fromCurrency")
	toCurrency := ctx.Query("toCurrency")
	rateStr := ctx.Query("rate")

	query := global.Db.Model(&models.ExchangeRate{})
	if fromCurrency != "" {
		query = query.Where("from_currency = ?", fromCurrency)
	}
	if toCurrency != "" {
		query = query.Where("to_currency = ?", toCurrency)
	}
	if rateStr != "" {
		rete, err := strconv.ParseFloat(rateStr, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rate parameter"})
			return
		}
		query = query.Where("rate = ?", rete)
	}

	var exchangeRates []models.ExchangeRate
	if err := query.Find(&exchangeRates).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "No exchange rates found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, exchangeRates)
}
