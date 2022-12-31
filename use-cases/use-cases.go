package usecases

import "github.com/georgexx009-wizeline/ondemand-go-bootcamp/utils/fetcher"

type UseCases struct {
	Fetcher fetcher.Fetcher
}

func NewUseCases(f fetcher.Fetcher) *UseCases {
	return &UseCases{Fetcher: f}
}
