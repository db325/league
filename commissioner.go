package main

import (
	"fmt"
)

//***********************    BEGIN COMMISSIONER    ***********************************************************
type Commissioner struct {
	Level        float32
	FirstName    string
	LastName     string
	Age          int
	Complaints   []*Complaint
	LeagueMember []*LeagueMember
}

func createCommish(fname, lname string) *Commissioner {

	commish := &Commissioner{
		FirstName: fname,
		LastName:  lname,
	}
	return commish
}

func (comish *Commissioner) Fine(lm *LeagueMember, amount float32, reason string) string {
	var lm1 LeagueMember
	lm1 = *lm
	lm1.Fine(amount)
	return fmt.Sprintf("You've been fined %d by the commisioner: %s", amount, reason)
}

//Returns a slice of stings. The first value [0]is the name of League member(Commissioner). The second value [1] is the age.
func (comish *Commissioner) GetName() string {
	name := comish.FirstName + " " + comish.LastName
	return name
}

func (commish *Commissioner) GetLevel() float32 {
	return commish.Level
}

//*
//*
