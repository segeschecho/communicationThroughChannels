package main

var baseStarMass = 8e28
func massUnits(mass int) float64 {
	return float64(mass)*baseStarMass
}


func main() {
	// Suma y multiplicacion en dos gorutinas
	//sumAndMulti()

	// El proceso principal le envia tareas a UN worker y queda esperando hasta que termine
	//mainWaitForGoroutinesFinishes()

	// Se crean 2 go rutinas y se espera a que terminen usando waitGroup
	//waitGroupSimple()

	// Se crean muchas gorutinas que suman 1 a una unica variable hasta 100, pero el resultado es 1000
	//raceCondition()

	//raceCondition2()

	//
	//goRoutinesWaitMainToStart()

	//
	//queueToProcess()

	//
	serverHttp()
}

