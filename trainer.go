package main

import (
	"fmt"
)

type Trainer struct {
	FirstName string
	LastName  string
	Age       int
	Kind      string
	Price     float32
	Training  bool
	Rating    float32
	Eligible  *Eligible
	LMType    string
}

func CreateTrainer(fname, lname string, age int, kind string, price float32) *Trainer {
	trainer := &Trainer{
		FirstName: fname,
		LastName:  lname,
		Age:       age,
		Kind:      kind,
		Price:     price,
	}
	trainer.LMType = "trainer"
	trainer.Eligible.LMActive = true
	return trainer
}

func (trainer *Trainer) GetName() string {
	return fmt.Sprintln(trainer.FirstName + " " + trainer.LastName)
}

func (trainer *Trainer) GetPrice() float32 {
	return trainer.Price
}

func (trainer *Trainer) IsTraining() bool {
	return trainer.Training
}

func (trainer *Trainer) GetKind() string {
	return trainer.Kind
}

//Kinds of trainer.
const (
	SA = "Speed/Agilty"
	SC = "Strength/Conditioning"
	AE = "Accuracy/Endurance"
)

func (trainer *Trainer) SetActive(yn bool) {
	if yn == true {
		trainer.Eligible.LMActive = true

	} else if yn == false {
		trainer.Eligible.LMActive = false
	}

}

func (trainer *Trainer) GetType() string {
	return trainer.LMType
}
