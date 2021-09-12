package main

import (
	genshinMW "GenshinPlanner/Backend/middleware"

	"github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("api/plan", genshinMW.GetAllPlan)
	e.POST("api/plan", genshinMW.CreateOnePlan)
	e.PUT("api/plan", genshinMW.UpdateOnePlan)
	e.DELETE("api/plan/:id", genshinMW.DeleteOnePlan)

	e.Logger.Fatal(e.Start(":8080"))
}
