package postgre

import (
	"context"
	"skeleton-svc/databases/postgre/models"
	"skeleton-svc/helpers"
	uModels "skeleton-svc/usecases/v1/models"
	"time"
)

func (d *postgreDatabase) GetRecords(ctx context.Context, params uModels.GetRecords) ([]models.Records, error) {
	helpers.PrintHeader()
	newStartDate, _ := time.Parse("2006-01-02", params.StartDate)
	newEndDate, _ := time.Parse("2006-01-02", params.EndDate)
	d.Logs.Println(newStartDate)
	var data []models.Records
	query := d.Db.Table("records")
	// Query for sum int array on postgresql
	query = query.Select("id, name, (SELECT SUM(A) FROM UNNEST(marks) AS A) AS total, created_at")
	query = query.Where("created_at::date <= ?::date and created_at::date >= ?::date", newStartDate, newEndDate)
	query = query.Order("created_at asc")
	query.Find(&data)

	return data, query.Error
}
