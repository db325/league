package main

import (
	"fmt"
	"strings"
)

//***********************    BEGIN COACH    ***********************************************************

//Defines COACH type to be used on/in leagues/teams.
type Coach struct {
	*Team

	Salary         float32
	Level          float32
	Atti           CAttributes
	TeamName       string
	AccountBalance float32
	Eligible       *Eligible
}

////////////////////////////////////////////////////////////////////
// Values for Coach object
type CAttributes struct {
	FirstName        string
	LastName         string
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

////////////////////////////////////////////////////////////////////

//*********************    Begin TradeReq  ***********************************************************************
type TradeRequest struct {
	Coach       *Coach
	ProposeThis *Athelete
	AndOrThis1  float32
	ForThis     *Athelete
	AndOrThis2  float32
	Approved    bool
}

func (tr *TradeRequest) ShowRequest() string {
	req := fmt.Sprintf(`
	 					Trade Request     
	 					Coach: %v proposes a trade 

						%v And Or  %v

	 					For This: %v And Or %v

	 
	 					Approved? :%v

`, strings.ToUpper(tr.Coach.Atti.FirstName+" "+tr.Coach.Atti.LastName), strings.ToUpper(tr.ProposeThis.Atti.Firstname+" "+tr.ProposeThis.Atti.Lastname), float32(tr.AndOrThis1), strings.ToUpper(tr.ForThis.Atti.Firstname+" "+tr.ForThis.Atti.Lastname), tr.AndOrThis2, false)
	return req
}

//*********************************************************************************
//Creates a Coach object with at least the first and last name. The team is optional. Usually upon creation,
//you will not set the team property. If left emty, it will default to n/a.
func createCoach(Fname, Lname string) *Coach {

	coach := &Coach{

		Atti: CAttributes{
			FirstName: Fname,
			LastName:  Lname,
		},
	}

	return coach
}

////////////////////////////////////////////////////////////////////

//Sign function adds Athelete to a roster. It also makes the player.Team variable equal to team.Name by default.
func (coach *Coach) Sign(name, pos string, team *Team, player *Athelete, jersey int) {
	player.Team = coach.Team
	//team.Players = append(team.Players, player)
	nm1 := fmt.Sprintf("%s\t%s", player.Atti.Firstname, player.Atti.Lastname)
	player.Atti.JerseyNum = jersey
	player.Team = coach.Team
	name = nm1
	player = (player)
	player.Team.Players = append(player.Team.Players, player)
	player.Atti.Team = player.Team.Name
	//player.SetActive(false)

}

////////////////////////////////////////////////////////////////////

func (coach *Coach) Cut(name, pos string, team *Team, player *Athelete) {
	nm1 := fmt.Sprintf("%s\t%s", player.Atti.Firstname, player.Atti.Lastname)
	name = nm1
	team.Roster[pos] = nil
	player.TeamName = "Free Agent"
	player.Team = nil
}

////////////////////////////////////////////////////////////////////

func (coach *Coach) GiveSpeach(speach string) string {
	return speach
}

////////////////////////////////////////////////////////////////////

func (coach *Coach) MakeTradeReq(requester *Coach, personPpl2Trade *Athelete, money1 float32, money2 float32, approve bool, approver *Coach, oneOrMore *Athelete) {
	TR := &TradeRequest{
		Coach:       requester,
		ProposeThis: personPpl2Trade,
		AndOrThis1:  money1,
		ForThis:     oneOrMore,
		AndOrThis2:  money2,
		Approved:    false,
	}
	approver.Atti.Requests = append(approver.Atti.Requests, TR)
}
func (coach *Coach) SetSalary(amount float32) {
	coach.Salary = amount
}

//Pay updates AccountBalance.
func (coach *Coach) Pay(amount float32) {
	PayAmount := coach.Salary / 12
	coach.AccountBalance += PayAmount
}
func (coach *Coach) Fine(amount float32) {
	coach.AccountBalance = coach.AccountBalance - amount
}
func (coach *Coach) GetName() string {
	name := coach.Atti.FirstName + " " + coach.Atti.LastName

	return name
}
func (coach *Coach) MediaPost(t, m string, v bool) {
	MP := &Message{
		Title:   t,
		Message: m,
		Visible: v,
	}
	if v == true {
		//IMPLEMENT LEAGUE BOARD
	} else if v == false {
		coach.Team.MessBoard = append(coach.Team.Posts, MP)

	}

}

//////////////////////////////////////////////////////////////////
func (coach *Coach) GetLevel() float32 {
	return coach.Level
}

func (coach *Coach) CreateComplaint(from, about LeagueMember, commish *Commissioner, issue string) {
	Complaint := &Complaint{
		From:  from.GetName(),
		About: about.GetName(),
		Issue: issue,
	}

	commish.Complaints = append(commish.Complaints, Complaint)
}
func (coach *Coach) SetActive(yn bool) {
	coach.Eligible.LMActive = yn
}

//SendSlip func appends a slip to the Coach's []*Slips
func (coach *Coach) SendSlip(slip *Slip) {
	coach.Eligible.Slips = append(coach.Eligible.Slips, slip)
}

func (coach *Coach) GetSlips() []*Slip {

	return coach.Eligible.Slips
}

////////////////////////////////////////////////////////////////////
