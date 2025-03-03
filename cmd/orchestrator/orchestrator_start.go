package main

import (
	"log"

	"github/stepanchigg/Final_Zad_2_Vozvrat/internal/orchestrator"
)

func main() {
	app := orchestrator.NewOrchestrator()
	log.Println("Оркестратор запускается на порту", app.Config.Addr)
	if err := app.RunServer(); err != nil {
		log.Fatal(err)
	}
}
