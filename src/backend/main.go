package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nambuitechx/go-monitor/backend/configs"
)

func main()  {
	envConfig := configs.NewEnvConfig()
	r := NewRouter(envConfig)

	log.Printf("Server is running on port %s", envConfig.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", envConfig.Port), r)
}
