package middleware

import (
	"GenshinPlanner/Backend/functions/characterLevel"
	"GenshinPlanner/Backend/models"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllPlan(c echo.Context) error {
	c.Response().Header().Set("Context-Type", "application/x-www-form-urlencoded")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	payLoad := GetAllCharLvPlan()
	return c.JSON(http.StatusOK, payLoad)
}

func CreateOnePlan(c echo.Context) error {
	c.Response().Header().Set("Context-Type", "application/x-www-form-urlencoded")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "POST")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var plan models.CharacterPlan
	_ = json.NewDecoder(c.Request().Body).Decode(&plan)

	plan.RequiredEXP = characterLevel.GetRequiredEXP(plan.CurrentLevel, plan.TargetLevel)
	InsertOneCharLvPlan(plan)

	return c.JSON(http.StatusOK, plan)
}

func UpdateOnePlan(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/x-www-form-urlencoded")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "PUT")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var plan models.CharacterPlan
	_ = json.NewDecoder(c.Request().Body).Decode(&plan)

	plan.RequiredEXP = characterLevel.GetRequiredEXP(plan.CurrentLevel, plan.TargetLevel)
	UpdateOneCharLvPlan(plan)

	return c.JSON(http.StatusOK, plan)
}

func DeleteOnePlan(c echo.Context) error {
	c.Response().Header().Set("Context-Type", "application/x-www-form-urlencoded")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "DELETE")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")

	plan := c.Param("id")
	DeleteOneCharLvPlan(plan)
	return c.JSON(http.StatusOK, plan)
}
