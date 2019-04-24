package main

import (
	"fmt"
	//"time"
)

type Manager struct {
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

//////////////////////////////////////////////////////////////////////////
func createManager(fname, lname string) *Manager {

	man := &Manager{
		FirstName: fname,
		LastName:  lname,
	}

	return man
}

//////////////////////////////////////////////////////////////////
//Function to sign Head Coach. **Must be a General Manager.
func (man *Manager) Sign(coach *Coach, player *Athelete) {

	if coach.GetType() == "head coach" && player == nil {

		coach.Eligible = &Eligible{
			Slips:      make([]*Slip, 0),
			LMActive:   true,
			Reason:     "",
			ReturnDate: 0,
		}
		coach.Atti.Requests = make([]*TradeRequest, 0)
		coach.Team = man.Team
		man.Team.Coaches = make([]*Coach, 0)
		man.Coaches = append(man.Coaches, coach)
		players := make([]*Athelete, 0)
		coach.Team.Players = players

	} else if player.GetType() == "player" && coach == nil {
		eligible := &Eligible{
			Slips:      make([]*Slip, 0),
			LMActive:   true,
			Reason:     "",
			ReturnDate: 0,
		}
		player.Eligible = eligible
	}

}

//fmt.Print(man.Team)

/////////////////////////////////////////////////////////////////
func (manager *Manager) Fire(coach *Coach, reason string) {
	coach.Team = nil

}

////////////////////////////////////////////////////////////////////

//******* Functions to implement LeagueMember Interface ************
func (manager *Manager) SetSalary(amount float32) {
	manager.Salary = amount
}

////////////////////////////////////////////////////////////////////
func (manager *Manager) Fine(amount float32) {
	manager.AccountBalance = manager.AccountBalance - amount
}

////////////////////////////////////////////////////////////////////
func (manager *Manager) GetLevel() int {
	return manager.Level
}
func (manager *Manager) MediaPost(t, m string, v bool) {
	Message := &Message{
		Title:   t,
		Message: m,
		Visible: v,
	}
	if v == true {
		//IMPLEMENT LEAGUE BOARD
	} else if v == false {
		manager.Team.MessBoard = append(manager.Team.MessBoard, Message)
	}
}
func (manager *Manager) GetName() string {

	name := manager.FirstName + " " + manager.LastName
	return name
}
func (manager *Manager) Pay() {
	PayAmount := manager.Salary / 12
	manager.AccountBalance += PayAmount
}

func (man *Manager) ToggleElig() {
	man.Eligible.LMActive = !man.Eligible.LMActive
}

func (manager *Manager) SendSlip(slip *Slip) {

	manager.Eligible.Slips = append(manager.Eligible.Slips, slip)
	return
}

func (man *Manager) GetType() string {
	return man.LMType
}

//slip[0] can be the only active slip
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

func (manager *Manager) GetSlips() []*Slip {
	for _, v := range manager.Eligible.Slips {
		v.SetTimeLeft()
	}

	return manager.Eligible.Slips
}
func (man *Manager) SetActive(yn bool) {}
