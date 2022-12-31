package main

import (
	"fmt"
	csvmanager "github.com/georgexx009-wizeline/ondemand-go-bootcamp/utils/csv-manager"
	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/utils/fetcher"

	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/controller"
	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/router"
	usecases "github.com/georgexx009-wizeline/ondemand-go-bootcamp/use-cases"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	useCases := usecases.NewUseCases(fetcher.Fetcher{}, csvmanager.CsvManager{})
	c := controller.NewController(useCases)
	c.Router = r
	router.NewRouter(r, c)

	fmt.Println("running app") // TODO - replace with logger
	err := r.Run()             // TODO - port from env vars
	if err != nil {
		fmt.Println("error running app") // TODO - replace with logger
	}
}
