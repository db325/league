package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//***********************    BEGIN COACH    ***********************************************************

//Defines COACH type to be used on/in leagues/teams.
type Coach struct {
	*Team
	*League
	Salary         float32
	Level          int
	Atti           CAttributes
	TeamName       string
	AccountBalance float32
	Eligible       *Eligible
	LMType         string
	FirstName      string
	LastName       string
}

// Values for Coach object
type CAttributes struct {
	CoachingStyle    string
	PlayCallAbility  float32
	SituationalAware float32
	RespectByPlayers bool
	Motivational     bool
	Unconventional   bool
	LikedByPlayers   bool
	TypeOfCoach      string
	Requests         []*TradeRequest
}

//*********************    Begin TradeReq  **************************************************************
//TradeReequest type
type TradeRequest struct {
	From        string
	ProposeThis *Athelete
	ForThis     *Athelete
	Approved    bool
}

//ShowRequest func returns a formatted string of the request.
func ShowReq(reqs []*TradeRequest) error {
	if len(reqs) == 0 {
		err := errors.New("Sorry, nothing to show here")
		return err
	}
	for _, req := range reqs {
		fmt.Printf("***REQUEST***\n%v Proposes trade below\n**********\nTrade: %v\nFor: %v\n%v", req.From, req.ProposeThis.GetName(), req.ForThis.GetName())
	}
	return nil
}

//*********************************************************************************
//Creates a Coach object with at least the first and last name. The team is optional. Usually upon creation,
//you will not set the team property. If left emty, it will default to n/a.
func createCoach(Fname, Lname string) *Coach {

	Coach := &Coach{
		FirstName: Fname,
		LastName:  Lname,
	}
	return Coach
}

//Sign function adds Athelete to a roster. It also initializes the player.Eligible struct's values to their
//zero-values except LMActive. LmActive is set to true as well as connecting all team/player fields.
func (coach *Coach) Sign(player *Athelete) {
	player.League = coach.League

	player.Eligible = &Eligible{
		Slips:      make([]*Slip, 0),
		Reason:     "",
		ReturnDate: 0,
		LMActive:   true,
	}
	player.Team = coach.Team
	player.Atti.Team = coach.TeamName
	player.TeamName = player.Atti.Team
	coach.Team.Players = append(coach.Players, player)
	coach.Roster[player.Atti.Position] = player
	player.Coach = coach
}

//Removes player from team.
func (coach *Coach) Cut(player *Athelete) {

	//nm1 := fmt.Sprintf("%s\t%s", player.Atti.Firstname, player.Atti.Lastname)

	CutPlayer(coach.Team, player)

}

//***Implement accepted and rejected requests.****
//Makes request and sends to coach.
func (coach *Coach) MakeTradeReq(person *Athelete, oneOrMore *Athelete) {
	TR := &TradeRequest{
		From:        coach.GetName(),
		ProposeThis: person,
		ForThis:     oneOrMore,
		Approved:    false,
	}
	oneOrMore.Coach.Atti.Requests = append(oneOrMore.Coach.Atti.Requests, TR)
}

//League Member Implementation
func (coach *Coach) GetName() string {
	name := coach.FirstName + " " + coach.LastName

	return strings.ToUpper(name)
}

//Sets salary variable.
func (coach *Coach) SetSalary(amount float32) {
	coach.Salary = amount
}

//Makes a post to a message board.
func (coach *Coach) MediaPost(t, m string, v bool) {
	MP := &Message{
		From:    coach.GetName(),
		Title:   t,
		Message: m,
		Visible: v,
	}
	if v == true {
		//IMPLEMENT LEAGUE BOARD
		coach.League.MessBoard = append(coach.League.MessBoard, MP)
	} else if v == false {
		coach.Team.MessBoard = append(coach.Team.MessBoard, MP)

	}

}

//Returns level variable.
func (coach *Coach) GetLevel() int {
	return coach.Level
}

//Subtracts amount from AccountBalance.
func (coach *Coach) Fine(amount float32) {
	coach.AccountBalance = coach.AccountBalance - amount
}

//Adds payamount to AccountBalance.
func (coach *Coach) Pay() {
	PayAmount := coach.Salary / 12
	coach.AccountBalance += PayAmount
}

//SendSlip func appends a slip to the Coach's []*Slips
func (coach *Coach) SendSlip(slip *Slip) {
	coach.Eligible.Slips = append(coach.Eligible.Slips, slip)
}

//Checks eligibility.
func (coach *Coach) CheckSuspension() {
	if coach.Eligible.LMActive == false {
		fmt.Println("Coach Check")
		for i := 0; i < len(coach.Eligible.Slips); i++ {
			if coach.Eligible.Slips[i].SActive == true {
				fmt.Println("Active Slip")
				if coach.Eligible.Slips[i].Time2Chech.Sub(time.Now()) < 0.0 {
					fmt.Println("Times Up!")
					coach.ToggleElig()
					coach.Eligible.Slips[i].SActive = false
				} else {
					fmt.Println("Banned Until ", coach.Eligible.Slips[i].End)
				}
			}
		}
	} else {
		fmt.Println("Coach Active")
	}
}

//Returns league member role.
func (coach *Coach) GetType() string {
	return strings.ToUpper(coach.LMType)
}

//Toggles eligibilty.
func (coach *Coach) ToggleElig() {
	coach.Eligible.LMActive = !coach.Eligible.LMActive
}

//Returns League.
func (coach *Coach) LeagueInfo() *League {
	return coach.League
}
