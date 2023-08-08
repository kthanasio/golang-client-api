package main

import (
	"github.com/kthanasio/golang-client-api/configs"
	"github.com/kthanasio/golang-client-api/internal/pkg"
)

var (
	apiEndpoint string
)

func main() {
	cfg, err := configs.LoadConfig("./")
	if err != nil {
		panic(err)
	}
	apiEndpoint = cfg.APIEndpoint

	resultado, err := pkg.GetCurrency(apiEndpoint)
	if err != nil {
		panic(err)
	}

	err = pkg.SaveFile(resultado)
	if err != nil {
		panic(err)
	}
}
