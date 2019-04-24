package main

import (
	"fmt"
	"strings"
	"time"
)

//***********************    BEGIN COACH    ***********************************************************

//Defines COACH type to be used on/in leagues/teams.
type Coach struct {
	*Team
	Salary         float32
	Level          int
	Atti           CAttributes
	TeamName       string
	AccountBalance float32
	Eligible       *Eligible
	LMType         string
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
func (coach *Coach) Sign(player *Athelete, coach2 *Coach) {
	if player == nil {
		coach2.Team = coach.Team
		eligible := &Eligible{
			Slips:      make([]*Slip, 0),
			LMActive:   true,
			Reason:     "",
			ReturnDate: 0,
		}
		coach.Eligible = eligible
		coach.Team.Coaches = append(coach.Team.Coaches, coach2)

	} else if coach2 == nil {

		player.Atti.Team = coach.Team.Name
		eligible := &Eligible{
			Slips:      make([]*Slip, 0),
			LMActive:   true,
			Reason:     "",
			ReturnDate: 0,
		}
		player.Eligible = eligible
		player.TeamName = coach.TeamName
		player.Team = coach.Team
		player.Atti.Team = player.TeamName
		coach.Team.Players = append(coach.Players, player)
		coach.Team.Roster[player.Atti.Position] = player

	}

}

////////////////////////////////////////////////////////////////////

func (coach *Coach) Cut(player *Athelete) {

	//nm1 := fmt.Sprintf("%s\t%s", player.Atti.Firstname, player.Atti.Lastname)
	player.Roster[player.Atti.Position] = nil
	player.TeamName = "Free Agent"
	player.Atti.Team = ""
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

////////////////////////////////////////////////////////////////////////
//League Member Implementation
func (coach *Coach) GetName() string {
	name := coach.Atti.FirstName + " " + coach.Atti.LastName

	return name
}
func (coach *Coach) SetSalary(amount float32) {
	coach.Salary = amount
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
		coach.Team.MessBoard = append(coach.Team.MessBoard, MP)

	}

}
func (coach *Coach) SetActive(yn bool) {
	coach.Eligible.LMActive = yn
}
func (coach *Coach) GetLevel() int {
	return coach.Level
}
func (coach *Coach) Fine(amount float32) {
	coach.AccountBalance = coach.AccountBalance - amount
}
func (coach *Coach) Pay() {
	PayAmount := coach.Salary / 12
	coach.AccountBalance += PayAmount
}

//SendSlip func appends a slip to the Coach's []*Slips
func (coach *Coach) SendSlip(slip *Slip) {
	coach.Eligible.Slips = append(coach.Eligible.Slips, slip)
}

func (coach *Coach) GetSlips() []*Slip {

	return coach.Eligible.Slips
}
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
func (coach *Coach) GetType() string {
	return strings.ToLower(coach.LMType)
}

/////////////////////////**********************/////////////////////////////////////////

func (coach *Coach) CreateComplaint(from, about LeagueMember, commish *Commissioner, issue string) {
	Complaint := &Complaint{
		From:  from.GetName(),
		About: about.GetName(),
		Issue: issue,
	}
	commish.Complaints = append(commish.Complaints, Complaint)
}
func (coach *Coach) ToggleElig() {
	coach.Eligible.LMActive = !coach.Eligible.LMActive
}
