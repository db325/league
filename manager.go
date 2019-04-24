package main

import (
	"fmt"
)

type Manager struct {
	*Team
	AccountBalance float32
	Level          float32
	FirstName      string
	LastName       string
	Type           string
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
func (man *Manager) SignCoach(coach *Coach) {

	coach.Team = man.Team
	coach.Team.Coaches = append(coach.Team.Coaches, coach)
	// Elig:=&Eligible{

	// }
	coach.Eligible = &Eligible{
		Slips:      make([]*Slip, 0),
		LMActive:   true,
		Reason:     "",
		ReturnDate: 0,
	}

	fmt.Print(man.Team)

}

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
func (manager *Manager) GetLevel() float32 {
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
func (manager *Manager) Pay(amount float32) {
	PayAmount := manager.Salary / 12
	manager.AccountBalance += PayAmount
}

///////////////////////// *****End Interface Implementation********** ///////////////////////////////////////////
func (manager *Manager) SendSlip(slip *Slip) {

}
