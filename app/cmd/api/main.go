package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Default().Printf("Erro ao carregar o arquivo de configuracao %s", err)
	}
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8000"
	}
	var cfg = Config{
		addr: addr,
	}
	var application = &Application{
		config: cfg,
	}
	mux := application.mount()
	log.Fatal(application.run(mux))
}
