package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//***********************    BEGIN OWNER    ******************************************************************

type Owner struct {
	Level          float32
	FirstName      string
	LastName       string
	Team           *Team
	Age            int
	Salary         float32
	AccountBalance float32
	Eligible       *Eligible
	LMType         string
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////

//***********************   BEGIN TEAM   *********************************************************************

//Team object. You must call createTeam and initialize it with a name. It is called in the createOwner function by default.
type Team struct {
	Name            string
	UpperManagement []*Manager
	Players         []*Athelete
	Coaches         []*Coach
	Roster          map[string]*Athelete
	MessBoard       []*Message
}

//***********************   END TEAM   ***********************************************************************
//*
//*
func createOwner(fname, lname, tname string) (*Owner, error) {
	if fname == "" || lname == "" {
		err := errors.New("You must enter a first and last name.")
		return nil, err
	} else {
		owner := &Owner{

			FirstName: fname,
			LastName:  lname,
		}
		kl, _ := owner.CreateTeam(tname)
		owner.Team = kl
		return owner, nil
	}
}

//createTeam Returns a pointer to a Team struct, initialized with a name.
func (owner *Owner) CreateTeam(name string) (*Team, error) {

	Eligible := &Eligible{
		Slips:      make([]*Slip, 0, 30),
		LMActive:   true,
		Reason:     "",
		ReturnDate: 0,
	}
	owner.Eligible = Eligible

	r := make(map[string]*Athelete)
	c := make([]*Coach, 0)
	p := make([]*Athelete, 0)
	m := make([]*Message, 0)
	u := make([]*Manager, 0)
	if name == "" {
		err := errors.New("Enter a name.")
		return nil, err
	} else {
		team := &Team{
			Name:            name,
			Roster:          r,
			Coaches:         c,
			Players:         p,
			MessBoard:       m,
			UpperManagement: u,
		}

		owner.Team = team
		return team, nil
	}

}

func (owner *Owner) Sign(gm *Manager) {
	if gm.GetType() == "general manager" {
		gm.TeamName = owner.Team.Name
		gm.Team = owner.Team
		owner.Team.UpperManagement = append(owner.Team.UpperManagement, gm)
		sl := make([]*Slip, 0)
		eligible := &Eligible{
			Slips:      sl,
			LMActive:   true,
			Reason:     "",
			ReturnDate: 0,
		}
		gm.Eligible = eligible
	} else {
		fmt.Println("You can only sign a GM.")
	}

}

////////////////////////////////////////////////////////////////////

func (owner *Owner) Pay() {
	PayAmount := owner.Salary / 12
	owner.AccountBalance += PayAmount

}

func (owner *Owner) MediaPost(t, m string, v bool) {
	Message := &Message{
		Title:   t,
		Message: m,
		Visible: v,
	}
	if v == true {
		//IMPLEMENT LEAGUE BOARD
	} else if v == false {
		owner.Team.MessBoard = append(owner.Team.MessBoard, Message)
	}
}

func (owner *Owner) SetSalary(amount float32) {
	owner.Salary = amount
}

//GetInfo Returns the first and last name as one string.
func (owner *Owner) GetName() string {
	name := owner.FirstName + " " + owner.LastName
	return name
}

func (owner *Owner) GetLevel() float32 {
	return owner.Level
}
func (owner *Owner) Fine(amount float32) {
	owner.Salary += owner.Salary - amount
}

func (owner *Owner) SetActive(yn bool) {
	if yn == true {
		owner.Eligible.LMActive = true

	} else if yn == false {
		owner.Eligible.LMActive = false
	}

}

func (owner *Owner) SendSlip(slip *Slip) {
	owner.Eligible.Slips = append(owner.Eligible.Slips, slip)
}

func (owner *Owner) GetSlips() []*Slip {
	slips := []*Slip{}
	for _, val := range owner.Eligible.Slips {
		slips = append(slips, val)
	}
	return slips
}

////////////////////////////////////////////////////////////////////
func (owner *Owner) GetType() string {
	return strings.ToLower(owner.LMType)
}
func (owner *Owner) Suspend(lm LeagueMember, mins float32, reason string) {
	fivemin := time.Minute * 5
	Slip := &Slip{
		Time2Chech: time.Now().Add(fivemin),
		Type:       Suspend,
		Reason:     reason,
		Start:      time.Now().Format(time.Kitchen),
		SActive:    true,
		End:        time.Now().Add(fivemin).Format(time.Kitchen),
	}
	Slip.SetTimeLeft()
	lm.SendSlip(Slip)
	lm.ToggleElig()
}

// 	if mins > 30 {
// 		mins = 30
// 	} else if mins > 30 && mins < 61 {
// 		mins = 60
// 	} else if mins > 60 {
// 		mins = 90
// 	}

// 	tm := time.Now()
// 	after30 := tm.Add(time.Minute * 30)
// 	afterHr := tm.Add(time.Minute * 30 * 2)
// 	hrHalf := tm.Add(time.Minute * 90)

// 	switch mins {
// 	case 30:
// 		Slip := &Slip{
// 			Time2Chech: after30,
// 			Type:       Suspend,
// 			Reason:     reason,
// 			SActive:    true,
// 			Start:      tm.Format(time.Kitchen),
// 			End:        after30.Format(time.Kitchen),
// 		}
// 		lm.SendSlip(Slip)
// 		lm.ToggleElig()
// 		return

// 	case 60:
// 		Slip := &Slip{
// 			Time2Chech: afterHr,
// 			Type:       Suspend,
// 			Reason:     reason,
// 			SActive:    true,
// 			Start:      tm.Format(time.Kitchen),
// 			End:        afterHr.Format(time.Kitchen),
// 		}
// 		lm.SendSlip(Slip)
// 		lm.ToggleElig()
// 		return

// 	case 90:
// 		Slip := &Slip{
// 			Time2Chech: hrHalf,
// 			Type:       Suspend,
// 			Reason:     reason,
// 			SActive:    true,
// 			Start:      tm.Format(time.Kitchen),
// 			End:        tm.Format(time.Kitchen),
// 		}
// 		lm.SendSlip(Slip)
// 		lm.ToggleElig()
// 		return

// 	default:
// 		fmt.Println("Please say who you want to suspend, for how long, and the reason.")
// 		return
// 	}

// }

func (owner *Owner) ToggleElig() {
	owner.Eligible.Slips[0].SActive = !owner.Eligible.Slips[0].SActive
	owner.Eligible.LMActive = !owner.Eligible.LMActive
}
