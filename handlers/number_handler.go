package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rowjay007/number-classifier-api/services"
	"github.com/rowjay007/number-classifier-api/utils"
	"net/http"
	"strconv"
)

func ClassifyNumber(c *gin.Context) {
	numStr := c.DefaultQuery("number", "")
	num, err := strconv.Atoi(numStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"number": numStr,
			"error":  true,
		})
		return
	}

	properties := services.GetNumberProperties(num)

	response := gin.H{
		"number":     num,
		"is_prime":   services.IsPrime(num),
		"is_perfect": false,
		"properties": properties,
		"digit_sum":  services.DigitSum(num),
		"fun_fact":   utils.FetchFunFact(num),
	}

	c.JSON(http.StatusOK, response)
}
