package usecases

import (
	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/model"
	logger "github.com/georgexx009-wizeline/ondemand-go-bootcamp/utils"
	"reflect"
	"strconv"
)

func (u *UseCases) SavePokemonsInCSV(fileName string, pokemons []model.Pokemon) error {
	fields := getFieldsFromStruct(pokemons[0])
	data := [][]string{fields}

	for _, pokemon := range pokemons {
		row := []string{strconv.Itoa(pokemon.Id), pokemon.Name, strconv.Itoa(pokemon.Weight)}
		data = append(data, row)
	}
	err := u.CsvManager.Save(fileName, data)
	if err != nil {
		logger.Log(err.Error())
		return err
	}

	return nil
}

func getFieldsFromStruct(s model.Pokemon) []string {
	var fields []string
	e := reflect.ValueOf(&s).Elem()

	for i := 0; i < e.NumField(); i++ {
		fields = append(fields, e.Type().Field(i).Name)
	}

	return fields
}
