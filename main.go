package main

import (
	"fmt"
	vehicle "recieverFuncSample/handler"
	human "recieverFuncSample/mapper"
)

type Employee struct {
	Name string
	age  int
	city string
}

func main() {
	fmt.Println("start main func started")
	var bmw vehicle.Vehicle
	bmw = vehicle.Car("World Top Brand")

	var labour human.Human
	labour = human.Man("Software Developer")

	fmt.Println(bmw.Speed())
	fmt.Println(labour.Performance())
	fmt.Println("end of main func call")
}
