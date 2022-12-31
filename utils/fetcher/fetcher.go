package fetcher

import (
	"encoding/json"
	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/model"
	"io"
	"net/http"

	logger "github.com/georgexx009-wizeline/ondemand-go-bootcamp/utils"
)

type Fetcher struct {
}

// when is good to use a simple function and when to use a method from struct

// GetPokemon - function to get a Pokemon info
func (f *Fetcher) GetPokemon(pokemonName string) (*model.Pokemon, error) {
	var URL = "https://pokeapi.co/api/v2/pokemon/"
	res, err := http.Get(URL + pokemonName)
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}

	byteArr, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}

	var pokemonModel model.Pokemon
	if err := json.Unmarshal(byteArr, &pokemonModel); err != nil {
		logger.Log(err.Error())
		return nil, err
	}

	return &pokemonModel, nil
}

//
//func (f *Fetcher) Get[T any](url string, deserializer func([]byte, T) (*T, error)) (*T, error) {
//	res, err := http.Get(url)
//	if err != nil {
//		logger.Log(err.Error())
//		return nil, err
//	}
//
//	byteArr, err := io.ReadAll(res.Body)
//	if err != nil {
//		logger.Log(err.Error())
//		return nil, err
//	}
//
//	body, err := deserializer(byteArr)
//	if err != nil {
//		logger.Log(err.Error())
//		return nil, err
//	}
//	return body, nil
//}
//
//func (f *Fetcher) GetV2[T any](url string, responseBodyType T) (*T, error) {
//	res, err := http.Get(url)
//	if err != nil {
//		logger.Log(err.Error())
//		return nil, err
//	}
//
//	byteArr, err := io.ReadAll(res.Body)
//	if err != nil {
//		logger.Log(err.Error())
//		return nil, err
//	}
//
//	if err := json.Unmarshal(byteArr, &responseBodyType); err != nil {
//		logger.Log(err.Error())
//		return nil, err
//	}
//
//	return &responseBodyType, nil
//}
