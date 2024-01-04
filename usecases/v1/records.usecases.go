package usecases

import (
	"context"
	pModels "skeleton-svc/databases/postgre/models"
	"skeleton-svc/helpers"
	hModels "skeleton-svc/helpers/models"
	"skeleton-svc/usecases/v1/models"
)

// GetRecords
func (uc *usecase) GetRecords(ctx context.Context, params models.GetRecords) hModels.Meta {
	helpers.PrintHeader()
	response := hModels.Meta{}
	var records []pModels.Records

	data, err := uc.DB.GetPostgre().GetRecords(ctx, params)
	if err != nil {
		// Logging for error
		uc.Logs.Error("Error get data ", err)
		response.Code = 1
		response.Msg = "Error Get Data"
		return response
	}

	for _, val := range data {
		if val.Total < params.Min {
			continue
		}

		if val.Total > params.Max {
			continue
		}
		records = append(records, val)
	}

	response.Code = 0
	response.Msg = "Success"
	response.Records = records
	return response
}
