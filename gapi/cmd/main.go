package main

import (
	"gapi-agp/infrastructure/config"
	"gapi-agp/server"
)

func main() {
	config.LoadConfig(config.CONFIG_PATH)

	//Iniciar repos de base de datos
	//Iniciar servicios (casos de uso)
	//Pasarselos al servidor
	//Iniciar servidor
	srv := server.NewServer()
	srv.StartServer()

}
