package simulation

import (
	"geneticAutomat/internal/console"
	"geneticAutomat/internal/model"
	"math/rand"
	"time"
)

const population int = 64

func Run() {
	//TODO: init logger

	//TODO: init console print
	var printer console.Console = console.Console{
		[]byte("H.E "),
	}
	//TODO: create world
	var world model.World = model.CreateWorld(25, 25, population)
	world.GenerateWalls()
	world.GenerateFood(2)

	//TODO: create goroutine of simulation
	for age := 0; age < 120; age++ {
		world.CountOfEntity = 0
		for i := 0; i < len(world.ArrayEntity); i++ {
			if world.ArrayEntity[i].Hp > 0 {
				world.ArrayEntity[i].RunDNA(&world)
				world.CountOfEntity++
			} else if world.ArrayEntity[i].Hp != -1 {
				world.ArrayEntity[i].Hp = -1
				world.UpdateEntityCell(world.ArrayEntity[i].Coordinates, nil)
			}
		}
		time.Sleep(2 * time.Millisecond)
		//TODO: print world Frame
		printer.Print(&world)
	}
}

func RunTrain(endTrainAge, endPopulation int) {
	var startPopulation = endPopulation * endPopulation
	//Это функция обучения ботов в заданных условиях мира
	//Следите за отношением endPopulation к startPopulation
	//todo: создать установку условий мира

	var printer console.Console = console.Console{
		[]byte("H.E "),
	}
	_ = printer

	var world model.World = model.CreateWorld(25, 25, startPopulation)
	world.GenerateWalls()

	//Создаём цикл обучения
	for counterWorld := 0; world.WorldAge <= endTrainAge; counterWorld++ { //работаем, пока время жизни мира не сравняется с требуемым временем

		world.WorldAge = 0 //сбрасываем возраст мира после прошлой попытки
		world.CountOfEntity = startPopulation
		world.GenerateFood(2)

		//работаем, пока в мире не останется endPopulation ботов
		for world.CountOfEntity > endPopulation {
			//проходимся по всем ботам
			world.CountOfEntity = 0
			for i := 0; i < startPopulation; i++ {
				//если хп больше нуля, то выполняем генетический код
				if world.ArrayEntity[i].IsLive {
					world.ArrayEntity[i].RunDNA(&world)
					world.CountOfEntity++
				}
			}
			world.WorldAge += 1
			//Отрисовать кадр
			time.Sleep(2 * time.Millisecond)
			printer.AlterPrint(&world, counterWorld)
		}
		//ботов стало меньше или равно endPopulation
		//Сортировка ботов по возрасту
		world.SortEntityByAge()
		//Замена генома в 8 группах
		for i := 0; i < endPopulation; i++ { //Лучшие endPopulation ботов
			for j := 1; j < endPopulation; j++ { //заменяют геном остальных
				world.ArrayEntity[i*endPopulation+j].SetDNA(world.ArrayEntity[i].DNA)
			}
		}
		//Мутирование генома у отдельных ботов
		for i := 0; i < endPopulation; i++ {
			world.ArrayEntity[rand.Intn(startPopulation-1)].Mutation(2)
		}
		//установка ботов
		for i := 0; i < len(world.ArrayEntity); i++ {
			world.ArrayEntity[i].Age = 0
			world.ArrayEntity[i].Hp = 100
			world.ArrayEntity[i].IsLive = true
			world.ArrayEntity[i].X = rand.Intn(world.Height-2) + 1
			world.ArrayEntity[i].Y = rand.Intn(world.Height-2) + 1

		}
	}
}
