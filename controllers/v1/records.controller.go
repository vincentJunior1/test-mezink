package controllers

import (
	"skeleton-svc/helpers"
	hModels "skeleton-svc/helpers/models"
	"skeleton-svc/usecases/v1/models"

	"github.com/gin-gonic/gin"
)

func (c *v1Controller) GetRecords(ctx *gin.Context) {
	helpers.PrintHeader()
	var res hModels.Meta
	params := models.GetRecords{}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		c.Logs.Error("Error bind json ", err)
		res.Code = 1
		res.Msg = "Error bind json"
		ctx.JSON(400, res)
		return
	}

	res = c.Usecase.GetRecords(ctx, params)
	// Logging response from get records usecases
	c.Logs.Info("Response Get Records", res)
	ctx.JSON(200, res)
}
