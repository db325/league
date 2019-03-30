package main

type Trainer struct {
	FirstName string
	LastName  string
	Age       int
	Kind      string
	Price     float32
	Training  bool
	Rating    float32
}

func CreateTrainer(fname, lname string, age int, kind string, price float32) *Trainer {
	trainer := &Trainer{
		FirstName: fname,
		LastName:  lname,
		Age:       age,
		Kind:      kind,
		Price:     price,
	}
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
