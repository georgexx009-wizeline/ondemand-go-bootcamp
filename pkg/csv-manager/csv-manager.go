package csv_manager

import (
	"encoding/csv"
	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/pkg/logger"
	"os"
)

type CsvManager struct {
}

func (c *CsvManager) Save(fileName string, data [][]string) error {
	csvFile, err := os.Create(fileName + ".csv")

	defer func() {
		err = csvFile.Close()
		if err != nil {
			logger.Log(err.Error())
		}
	}()
	if err != nil {
		logger.Log(err.Error())
		return err
	}

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	err = csvWriter.WriteAll(data)
	if err != nil {
		logger.Log(err.Error())
		return err
	}

	return nil
}
