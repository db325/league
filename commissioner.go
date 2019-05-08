package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//***********************    BEGIN COMMISSIONER    ***********************************************************
type Commissioner struct {
	*League
	Level          int
	FirstName      string
	LastName       string
	Age            int
	Complaints     []*Complaint
	LeagueMember   []LeagueMember
	LMType         string
	Salary         float32
	AccountBalance float32
	Eligible       *Eligible
}

//Create a Commissioner.
func createCommish(fname, lname string) *Commissioner {

	commish := &Commissioner{

		FirstName: fname,
		LastName:  lname,
		//League
	}
	Elig := &Eligible{
		Reason:     "",
		LMActive:   true,
		Slips:      make([]*Slip, 0),
		ReturnDate: 0,
	}
	commish.Eligible = Elig
	return commish
}

//Func AddOwner adds owner to league. If there is a team with the same name as the owner, err will not be nil.
func (com *Commissioner) AddOwner(own *Owner) error {
	var err = errors.New("")
	for i := 0; i < len(com.League.Teams); i++ {
		if own.League.Teams[i].Name == own.Team.Name {
			err1 := errors.New("Team already exists. Please try another name.")
			err = err1
			return err
		}

	}
	own.League = com.League
	com.League.Owners = append(com.League.Owners, own)
	return nil

}

//Func GetName returns a string of the first and last name.
func (comish *Commissioner) GetName() string {
	name := comish.FirstName + " " + comish.LastName
	return name
}

//Func MediaPost posts a Message to a messageboard. Err is nil if the Message posts successfully.
func (commish *Commissioner) MediaPost(t, m string, v bool) (*Message, error) {
	Msg := &Message{
		Title:   t,
		Message: m,
		Visible: v,
	}
	if v == true {
		commish.League.MessBoard = append(commish.League.MessBoard, Msg)
		return Msg, nil
	} else {
		return nil, errors.New("Post was unsuccessful.")
	}
}

//Suspend suspends a lm for a set number of hours/days
func (commish *Commissioner) Suspend(lm LeagueMember, mins int, reason string) {
	//getting current time.
	tm := time.Now()
	//Creating time properties.
	if mins <= 30 {
		mins = 30
	} else if mins > 30 && mins < 90 {
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
		lm.ToggleElig()
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
		lm.ToggleElig()
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
		lm.ToggleElig()
		return
	default:
		fmt.Println("you need to enter a multiple of 30 minutes. Not to exceed 1.5 hrs.")
	}

}

func (commish *Commissioner) GetLevel() int {
	return commish.Level
}

//Func GetType returns a string value for the league member's role in the league.
func (com *Commissioner) GetType() string {
	return strings.ToLower(com.LMType)

}

//SetSalary sets salary variable in com struct.
func (commish *Commissioner) SetSalary(amount float32) {
	commish.Salary = amount

}

//Adds payamount into AccountBalance.
func (commish *Commissioner) Pay() {
	payAmount := commish.Salary / 12
	commish.AccountBalance += payAmount
}

//Subtracts amount from AccountBalance.
func (commish *Commissioner) Fine(amount float32) {
	commish.AccountBalance -= amount
}

// MAY GET RID OF>>>>> SetActive(yn bool)
// func (commish *Commissioner) Fined(amount float32) {
// 	commish.AccountBalance -= amount

// }

//Toggles eligible status.
func (commish *Commissioner) ToggleElig() {
	commish.Eligible.LMActive = !commish.Eligible.LMActive
}

//Puts a Slip in Slips
func (commish *Commissioner) SendSlip(slip *Slip) {
	commish.Eligible.Slips = append(commish.Eligible.Slips, slip)
}

//Check Suspension.
func (commish *Commissioner) CheckSuspension() {
	if commish.Eligible.LMActive == false {
		for i := 0; i < len(commish.Eligible.Slips); i++ {
			if commish.Eligible.Slips[i].SActive == true {
				if commish.Eligible.Slips[i].Time2Chech.Sub(time.Now()) < 0.0 {
					commish.ToggleElig()
					commish.Eligible.Slips[i].SActive = false
					commish.Eligible.Slips[i].TimeLeft = 0.0
				} else {
					//MAY NEED TO CHANGE TO SPRINTF IN THE FUTURE
					fmt.Println("Banned Until ", commish.Eligible.Slips[i].End)
				}
			}
		}
	} else {
		fmt.Println("Commish Active")
	}
}

//Slip struct
type Slip struct {
	Time2Chech time.Time

	Type    string
	Reason  string
	SActive bool
	Start   string
	End     string

	TimeLeft time.Duration
}

//Sets time left on active Ticket.
func (slip *Slip) SetTimeLeft() {
	now := time.Now()
	slip.TimeLeft = slip.Time2Chech.Sub(now)
}
