// handlers.go

package main

import (
	"github.com/gin-gonic/gin"
)

func showIndexPage(ctx *gin.Context) {
	render(ctx, gin.H{
		"title": "Home Page",
	}, "index.html")
}
