package main

import (
	"fmt"
	//"time"
)

type Manager struct {
	*League
	*Team
	AccountBalance float32
	Level          int
	FirstName      string
	LastName       string
	LMType         string
	TeamName       string
	CanHire        bool
	CanFire        bool
	Salary         float32
	Eligible       *Eligible
}

//Makes a Manager .
func createManager(fname, lname string) *Manager {

	man := &Manager{
		LMType:    "general manager",
		FirstName: fname,
		LastName:  lname,
	}

	return man
}

//Function to sign Head Coach. **Must be a General Manager.
func (man *Manager) Sign(coach *Coach) {
	coach.League = man.League
	coach.Eligible = &Eligible{
		Slips:      make([]*Slip, 0),
		LMActive:   true,
		Reason:     "",
		ReturnDate: 0,
	}

	coach.Atti.Requests = make([]*TradeRequest, 0)
	coach.Team = man.Team
	man.Team.Coaches = append(man.Team.Coaches, coach)
	coach.TeamName = man.TeamName
	coach.Team.Players = make([]*Athelete, 0)
	coach.Roster = make(map[string]*Athelete, 0)
	coach.LMType = "head coach"

}

//Removes coach from team as well as all team/player properties.
func (manager *Manager) Fire(coach *Coach, reason string) {
	fmt.Println("fire called")
	for i := 0; i < len(manager.Team.Coaches); i++ {
		if manager.Team.Coaches[i] == coach {
			manager.Team.Coaches[i].Team = nil
			manager.Team.Coaches[i].TeamName = ""
			manager.Team.Coaches = append(manager.Coaches[:i], manager.Coaches[i+1:]...)
		} else {
			fmt.Println("Coach doesn't exist.")
		}

	}

}

//******* Functions to implement LeagueMember Interface ************
//Sets salary variable.
func (manager *Manager) SetSalary(amount float32) {
	manager.Salary += amount
}

//Subtracts amount from AccountBalance.
func (manager *Manager) Fine(amount float32) {
	manager.AccountBalance -= amount
}

//Returns Level.
func (manager *Manager) GetLevel() int {
	return manager.Level
}

//Posts to message board. If true, the entire league can see, if false only the team sees.
func (manager *Manager) MediaPost(t, m string, v bool) {
	Message := &Message{
		From:    manager.GetName(),
		Title:   t,
		Message: m,
		Visible: v,
	}
	if v == true {
		manager.League.MessBoard = append(manager.League.MessBoard, Message)
	} else if v == false {
		manager.Team.MessBoard = append(manager.Team.MessBoard, Message)
	}
}

//Returns a string of manager's first and last name.
func (manager *Manager) GetName() string {

	name := manager.FirstName + " " + manager.LastName
	return name
}

//Adds pay amount to account balance.
func (manager *Manager) Pay() {
	PayAmount := manager.Salary / 12
	manager.AccountBalance += PayAmount
}

//Toggles manager's eligibility.
func (man *Manager) ToggleElig() {
	man.Eligible.LMActive = !man.Eligible.LMActive
}

//Puts a slip in Slips.
func (manager *Manager) SendSlip(slip *Slip) {

	manager.Eligible.Slips = append(manager.Eligible.Slips, slip)
	return
}

//Returns league member's role.
func (man *Manager) GetType() string {
	return man.LMType
}

//slip[0] can be the only active slip ***********REDO THIS********
func (man *Manager) CheckSuspension() {
	if man.Eligible.LMActive == false {
		for i := 0; i < len(man.Eligible.Slips); i++ {
			if man.Eligible.Slips[i].TimeLeft < 0.0 {
				man.ToggleElig()
				man.Eligible.Slips[i].SActive = false
			} else {
				fmt.Println("Banned until ", man.Eligible.Slips[i].End)
			}
		}
	} else {
		fmt.Println("Manager Active")
	}

}
