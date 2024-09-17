package usecases

import (
	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/model"
	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/pkg/fetcher"
	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/pkg/logger"
)

//var URL = "https://pokeapi.co/api/v2/pokemon/"

func getPokemon(fetcher fetcher.Fetcher, pokemonName string) (*model.Pokemon, error) {
	logger.Log("consuming poke api")

	//deserializer := func(pokemonBytes []byte) (*model.Pokemon, error) {
	//	var pokeResult model.Pokemon
	//	err := json.Unmarshal(pokemonBytes, &pokeResult)
	//	if err != nil {
	//		logger.Log(err.Error())
	//		return nil, err
	//	}
	//
	//	logger.Log(pokeResult)
	//	return &pokeResult, nil
	//}

	// pokemon, err := fetcher.Get(URL + pokemonName, deserializer)
	//var pokemonModel model.Pokemon
	//pokemon, err := fetcher.GetV2(URL+pokemonName, pokemonModel)
	pokemon, err := fetcher.GetPokemon(pokemonName)
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}
	return pokemon, nil
}

func (u *UseCases) ConsumePokeapi() ([]model.Pokemon, error) {
	pikachuInfo, err := getPokemon(u.Fetcher, "pikachu")
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}

	charmanderInfo, err := getPokemon(u.Fetcher, "charmander")
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}

	return []model.Pokemon{*pikachuInfo, *charmanderInfo}, nil
}
