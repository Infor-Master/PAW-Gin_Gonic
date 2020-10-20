// handlers.go

package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(ctx *gin.Context) {
	render(ctx, gin.H{
		"title": "Home Page",
	}, "index.html")
}

func showCalcPage(ctx *gin.Context) {
	render(ctx, gin.H{
		"title": "Calculator",
	}, "calculator.html")
}

func doCalculation(ctx *gin.Context) {
	strval1 := ctx.PostForm("val1")
	strval2 := ctx.PostForm("val2")
	val1, _ := strconv.ParseFloat(strval1, 64)
	val2, _ := strconv.ParseFloat(strval2, 64)
	op := ctx.PostForm("op")
	res := strval1 + " " + op + " " + strval2 + " = "
	switch op {
	case "+":
		res = res + strconv.FormatFloat(val1+val2, 'f', 2, 64)
	case "-":
		res = res + strconv.FormatFloat(val1-val2, 'f', 2, 64)
	case "*":
		res = res + strconv.FormatFloat(val1*val2, 'f', 2, 64)
	case "/":
		if val2 == 0 {
			res = "Cannot divide by 0"
		} else {
			res = res + strconv.FormatFloat(val1/val2, 'f', 2, 64)
		}
	}

	render(ctx, gin.H{
		"title":  "Calculator",
		"result": res,
	}, "calculator.html")

}
