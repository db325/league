package main

import (
	"errors"
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
}

////////////////////////////////////////////////////////////////////

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

//createTeam Returns a pointer to a Team struct, initialized with a name.
func (owner *Owner) CreateTeam(name string) (*Team, error) {

	Eligible := &Eligible{
		Slips:      make([]*Slip, 0, 30),
		LMActive:   true,
		Reason:     "",
		ReturnDate: 0,
	}
	owner.Eligible = Eligible

	m := make(map[string]*Athelete)
	if name == "" {
		err := errors.New("Enter a name.")
		return nil, err
	} else {
		team := &Team{
			Name:   name,
			Roster: m,
		}
		owner.Team = team
		return team, nil
	}

}

func (owner *Owner) SignGM(gm *Manager) {
	gm.Team = owner.Team
	owner.Team.UpperManagement = append(owner.Team.UpperManagement, gm)
	gm.TeamName = owner.Team.Name
}

////////////////////////////////////////////////////////////////////

func (owner *Owner) Pay(amount float32) {
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
