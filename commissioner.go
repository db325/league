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
func (commish *Commissioner) Suspend(lm LeagueMember, mins int, reason string) {

	tm := time.Now()
	if mins > 30 {
		mins = 30
	} else if mins < 30 && mins > 90 {
		mins = 60
	} else {
		mins = 90
	}
	after30 := tm.Add(time.Minute * 30)
	afterHr := tm.Add(time.Minute * 30 * 2)
	hrHalf := tm.Add(time.Minute * 90)

	switch mins {
	case 30:

		Slip := &Slip{
			Time2Chech: after30,
			Type:       Suspend,
			Reason:     reason,
			SActive:    true,
			Start:      tm.Format(time.Kitchen),
			End:        after30.Format(time.Kitchen),
		}
		Slip.SetTimeLeft()
		lm.SendSlip(Slip)
		lm.SetActive(false)
		return

	case 60:
		Slip := &Slip{
			Time2Chech: afterHr,
			Type:       Suspend,
			Reason:     reason,
			SActive:    true,
			Start:      tm.Format(time.Kitchen),
			End:        afterHr.Format(time.Kitchen),
		}
		Slip.SetTimeLeft()
		lm.SendSlip(Slip)
		lm.SetActive(false)
		return

	case 90:
		Slip := &Slip{
			Time2Chech: hrHalf,
			Type:       Suspend,
			Reason:     reason,
			SActive:    true,
			Start:      tm.Format(time.Kitchen),
			End:        hrHalf.Format(time.Kitchen),
		}
		Slip.SetTimeLeft()

		lm.SendSlip(Slip)
		lm.SetActive(false)
		return
	default:
		fmt.Println("you need to enter a multiple of 30 minutes. Not to exceed 1.5 hrs.")
	}
	lm.SetActive(false)

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
	Time2Chech time.Time

	Type    string
	Reason  string
	SActive bool
	Start   string
	End     string

	TimeLeft time.Duration
}

// func (slip *Slip) Show() *Slip {
// 	//checkTime(slip)
// 	// return fmt.Sprintf(`
// 	// Type:   %s
// 	// Reason: %s
// 	// Slip Active: %v
// 	// Start: %v   ComeBack:%v
// 	// Time Left:%v
// 	// time show func called: %v

// 	// `, slip.Type, slip.Reason, slip.SActive, slip.Start, slip.End, slip.TimeLeft, time.Now())
// 	return slip

// }

func (slip *Slip) SetTimeLeft() {
	now := time.Now()
	slip.TimeLeft = slip.Time2Chech.Sub(now)
}
