// routes.go

package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)

	// Handle the calculator route
	calc := router.Group("/calc")
	{
		calc.GET("/", showCalcPage)
		calc.POST("/sum", doCalcSum)
		calc.POST("/sub", doCalcSub)
		calc.POST("/mult", doCalcMult)
		calc.POST("/div", doCalcDiv)
		calc.POST("/res", doCalcRes)
	}
}
