package main

import (
	"fmt"
)

func main() {

	commish := createCommish("hi", "there")
	fmt.Println(commish)
	coach := createCoach("ko", "io")
	fmt.Println(coach)
	p1 := createPlayer("Dajovan", "Bryant", "Rec")
	p1.AccountBalance = 5000.00
	train := CreateTrainer("mr", "armstrong", 32, "Strength/Agility", 500)
	train.Kind = SA
	fmt.Println(train)
	//p1.Train(train)
	fmt.Println(p1.AccountBalance, " ", p1.Atti.Speed, p1.Atti.Agility)
	p1.Train(train)
	p1.Train(train)
	fmt.Println(p1.AccountBalance, " ", p1.Atti.Speed, p1.Atti.Agility)
	p1.Train(train)
	p1.Train(train)
	p1.Train(train)
	p1.Train(train)
	fmt.Println(p1.AccountBalance, " ", p1.Atti.Speed, p1.Atti.Agility)

}
