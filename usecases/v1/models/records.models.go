package models

type GetRecords struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Min       int    `json:"min"`
	Max       int    `json:"max"`
}
