// handlers.go

package main

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showCalcPage(ctx *gin.Context) {
	render(ctx, gin.H{
		"title": "Calculator",
	}, "calculator.html")
}

func doCalcSum(ctx *gin.Context) {
	val1, err := strconv.ParseFloat(ctx.PostForm("val1"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	val2, err := strconv.ParseFloat(ctx.PostForm("val2"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := strconv.FormatFloat(val1+val2, 'f', 2, 64)
	ctx.JSON(http.StatusOK, gin.H{
		"res": res,
	})
}

func doCalcSub(ctx *gin.Context) {
	val1, err := strconv.ParseFloat(ctx.PostForm("val1"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	val2, err := strconv.ParseFloat(ctx.PostForm("val2"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := strconv.FormatFloat(val1-val2, 'f', 2, 64)
	ctx.JSON(http.StatusOK, gin.H{
		"res": res,
	})
}

func doCalcMult(ctx *gin.Context) {
	val1, err := strconv.ParseFloat(ctx.PostForm("val1"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	val2, err := strconv.ParseFloat(ctx.PostForm("val2"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := strconv.FormatFloat(val1*val2, 'f', 2, 64)
	ctx.JSON(http.StatusOK, gin.H{
		"res": res,
	})
}

func doCalcDiv(ctx *gin.Context) {
	val1, err := strconv.ParseFloat(ctx.PostForm("val1"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	val2, err := strconv.ParseFloat(ctx.PostForm("val2"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if val2 == 0 {
		ctx.JSON(http.StatusOK, gin.H{"res": "Cannot divide by 0"})
		return
	}
	res := strconv.FormatFloat(val1/val2, 'f', 2, 64)
	ctx.JSON(http.StatusOK, gin.H{"res": res})
}

func doCalcRes(ctx *gin.Context) {
	val1, err := strconv.ParseFloat(ctx.PostForm("val1"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	val2, err := strconv.ParseFloat(ctx.PostForm("val2"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if val2 == 0 {
		ctx.JSON(http.StatusOK, gin.H{"res": "Cannot divide by 0"})
		return
	}
	res := strconv.FormatFloat(math.Mod(val1, val2), 'f', 2, 64)
	ctx.JSON(http.StatusOK, gin.H{"res": res})

}
