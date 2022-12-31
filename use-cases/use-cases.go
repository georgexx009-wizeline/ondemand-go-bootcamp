package usecases

import (
	csvmanager "github.com/georgexx009-wizeline/ondemand-go-bootcamp/pkg/csv-manager"
	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/pkg/fetcher"
)

type UseCases struct {
	Fetcher    fetcher.Fetcher
	CsvManager csvmanager.CsvManager
}

func NewUseCases(f fetcher.Fetcher, c csvmanager.CsvManager) *UseCases {
	return &UseCases{Fetcher: f, CsvManager: c}
}
