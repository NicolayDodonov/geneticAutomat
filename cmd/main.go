package main

import (
	"fmt"
	"geneticAutomat/internal/simulation"
	"geneticAutomat/internal/slogger"
	"log"
	"os"
)

func main() {
	err := os.MkdirAll(`log`, 0777)
	WorldAgeInfoWriter, err := os.Create(`log\WorldAgeInfo.txt`)
	if err != nil {
		log.Fatalf("Ошибка запуска логгирования: WorldAgeInfo\n%v", err)
	}
	WorldBestInfoWriter, err := os.Create(`log\WorldEndInfo.txt`)
	if err != nil {
		log.Fatalf("Ошибка запуска логгирования: WorldBestInfo.txt\n%v", err)
	}
	slogger.LogWorldAge = slogger.SetupLogger("dev", WorldAgeInfoWriter)
	slogger.LogWorldBest = slogger.SetupLogger("dev", WorldBestInfoWriter)

	simulation.RunTrain(1000000, 10)
	fmt.Scanln()
}
