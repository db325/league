package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//***********************    BEGIN OWNER    ******************************************************************
//Owner Struct. You must call createOwner to get an Owner type.
type Owner struct {
	League         *League
	Level          int
	FirstName      string
	LastName       string
	Team           *Team
	Age            int
	Salary         float32
	AccountBalance float32
	Eligible       *Eligible
	LMType         string
}

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

//Create a Team.
func (owner *Owner) CreateTeam(name string) (*Team, error) {
	//Make Eligible
	Eligible := &Eligible{
		Slips:      make([]*Slip, 0, 30),
		LMActive:   true,
		Reason:     "",
		ReturnDate: 0,
	}
	owner.Eligible = Eligible
	//Team Properties
	r := make(map[string]*Athelete)
	c := make([]*Coach, 0)
	p := make([]*Athelete, 0)
	m := make([]*Message, 0)
	u := make([]*Manager, 0)
	if name == "" {
		err := errors.New("Enter a name.")
		return nil, err
	} else {
		//Create Team
		team := &Team{
			Name:            name,
			Roster:          r,
			Coaches:         c,
			Players:         p,
			MessBoard:       m,
			UpperManagement: u,
		}
		//Set Team
		owner.Team = team
		owner.LMType = "owner"
		return team, nil
	}

}

//Func CutPlayer removes player from team and nullifies any team/player fields.
func CutPlayer(t *Team, p1 *Athelete) {
	for i := 0; i < len(t.Players); i++ {
		if t.Players[i] == p1 {
			t.Players = append(t.Players[:i], t.Players[i+1:]...)
			team := make([]*Athelete, 1)
			t.Players = append(t.Players, team...)
			p1.Roster[p1.Atti.Position] = nil
			p1.TeamName = "Free Agent"
			p1.Atti.Team = ""
			p1.Team = nil
			p1.Coach = nil
			fmt.Println("deleting player: ", p1.GetName(), " from roster")
		} else {
			fmt.Println("player isnt here")
		}
	}
}

//***********************   END TEAM   ***********************************************************************
//Create an Owner.
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

//Sign General Manager.
func (owner *Owner) Sign(gm *Manager) {
	if gm.GetType() == "general manager" {
		gm.League = owner.League
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

//Adds payamount to AccontBalance.
func (owner *Owner) Pay() {
	PayAmount := owner.Salary / 12
	owner.AccountBalance += PayAmount

}

//Makes post to MessageBoard.
func (owner *Owner) MediaPost(t, m string, v bool) {
	Message := &Message{
		From:    owner.GetName(),
		Title:   t,
		Message: m,
		Visible: v,
	}
	if v == true {
		//IMPLEMENT LEAGUE BOARD
		owner.League.MessBoard = append(owner.League.MessBoard, Message)

	} else if v == false {
		owner.Team.MessBoard = append(owner.Team.MessBoard, Message)
	}

}

//Sets Salary.
func (owner *Owner) SetSalary(amount float32) {
	owner.Salary = amount
}

//Returns a string of the first and last name.
func (owner *Owner) GetName() string {
	name := owner.FirstName + " " + owner.LastName
	return strings.ToUpper(name)
}

//Returns Level.
func (owner *Owner) GetLevel() int {
	return owner.Level
}

//Subtracts amount from AccountBalance
func (owner *Owner) Fine(amount float32) {
	if owner.AccountBalance < 0 {
		owner.AccountBalance = 0.0
	} else {
		owner.AccountBalance -= amount
	}
}

//Puts slip in Slips slice.
func (owner *Owner) SendSlip(slip *Slip) {
	owner.Eligible.Slips = append(owner.Eligible.Slips, slip)
}

//Returns a string for the league member type variable.
func (owner *Owner) GetType() string {
	return strings.ToUpper(owner.LMType)
}

//***********POSSIBLY CONVER TO A REGULAR FUNCTION****************
func (owner *Owner) Suspend(lm LeagueMember, mins float32, reason string) {

	if mins <= 30 {
		mins = 30
	} else if mins > 30 && mins <= 60 {
		mins = 60
	} else if mins > 60 {
		mins = 90
	}

	tm := time.Now()
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
			End:        tm.Format(time.Kitchen),
		}
		lm.SendSlip(Slip)
		lm.ToggleElig()
		return

	default:
		fmt.Println("Please say who you want to suspend, for how long, and the reason.")
		return
	}

}

//Toggles eligibility of owner to participate in league activities.
func (owner *Owner) ToggleElig() {
	owner.Eligible.LMActive = !owner.Eligible.LMActive
}

//Checks for active slips.Usually called before attempting any league related actions.
func (owner *Owner) CheckSuspension() {
	if owner.Eligible.LMActive == false {
		for i := 0; i < len(owner.Eligible.Slips); i++ {
			if owner.Eligible.Slips[i].TimeLeft < 0.0 {
				owner.ToggleElig()
				owner.Eligible.Slips[i].SActive = false
			} else {
				fmt.Println("Banned until ", owner.Eligible.Slips[i].End)
			}
		}
	} else {
		fmt.Println("Owner Active")
	}
}
