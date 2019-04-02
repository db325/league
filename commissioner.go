package main

import (
	"fmt"
	"time"
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

//Suspend suspends a lm for a set number of hours/days
func (commish *Commissioner) Suspend(lm LeagueMember, numOfHrs int, reason string) {
	now := time.Now()
	numOfHrs += now.Hour()
	then := time.Date(now.Year(), now.Month(), now.Day(), numOfHrs, 0, 0, 0, time.UTC)
	//now.String()
	Slip := &Slip{
		Type:    Suspend,
		Reason:  reason,
		SActive: true,
		Until:   now.Sub(then).Hours(),
	}
	lm.SetActive(false)
	lm.SendSlip(Slip)
	//check type of member

}

func checkType(lm LeagueMember) {
	switch lm.(type) {
	case *Athelete:
		ath := lm
		fmt.Println(ath)

	}
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
type Slip struct {
	Type    string
	Reason  string
	SActive bool
	Until   float64
}

func (slip *Slip) Show() string {
	return fmt.Sprintf(`
	Type:   %s
	Reason: %s
	Slip Active: %b
	Until: %s
	`, slip.Type, slip.Reason, slip.SActive, slip.Until)
}
