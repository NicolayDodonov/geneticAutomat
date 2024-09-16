package main

import (
	"geneticAutomat/internal/simulation"
	"geneticAutomat/internal/slogger"
	"log"
	"os"
)

const (
	EndTrainAge   = 1000000
	EndPopulation = 10
)

func main() {
	_ = os.MkdirAll(`log`, 0777)

	//TODO: добавить конфиг

	startLogger()

	simulation.RunTrain(EndTrainAge, EndPopulation)
}

func startLogger() {
	WorldAgeInfoWriter, err := os.Create(`log\WorldAgeInfo.txt`)
	if err != nil {
		log.Fatalf("Ошибка запуска логгирования: WorldAgeInfo\n%v", err)
	}
	WorldBestInfoWriter, err := os.Create(`log\WorldEndInfo.txt`)
	if err != nil {
		log.Fatalf("Ошибка запуска логгирования: WorldBestInfo.txt\n%v", err)
	}
	EntityInfoWriter, err := os.Create(`log\EntityInfo.txt`)
	if err != nil {
		log.Fatalf("Ошибка запуска логгирования: EntityInfo.txt\n%v", err)
	}
	Errors, err := os.Create(`log\Errors.txt`)
	if err != nil {
		log.Fatalf("Ошибка запуска логгирования: Error.txt\n%v", err)
	}

	slogger.LogWorldAge = slogger.SetupLogger("dev", WorldAgeInfoWriter)
	slogger.LogWorldBest = slogger.SetupLogger("dev", WorldBestInfoWriter)
	slogger.LogEntityInfo = slogger.SetupLogger("dev", EntityInfoWriter)
	slogger.LogErrors = slogger.SetupLogger("dev", Errors)
}
