package main

import (
	database "github.com/nugrohosam/gosampleapi/services/databases"
	httpConn "github.com/nugrohosam/gosampleapi/services/http"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(".env.yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(err)
	}

	if err := database.Conn(); err != nil {
		panic(err)
	}

	if err := httpConn.Serve(); err != nil {
		panic(err)
	}
}
