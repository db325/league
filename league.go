package main

import (
	"fmt"
	//"time"
	"errors"
)

//************************    BEGIN LEAGUE    ************************************************************************

//League struct is a generic struct for a League object.
type League struct {
	Type         string
	Commissioner *Commissioner
	Owners       []*Owner
	MessBoard    []*Message
	Game         *Game
	Teams        []*Team
}

func ShowPosts(mess []*Message) error {
	if len(mess) == 0 {
		err := errors.New("Sorry, nothing to show.")
		return err
	}
	for _, v := range mess {
		fmt.Printf("%v says: %v \n", v.From, v.Message)
	}
	return nil
}

func CreateComplaint(from, about LeagueMember, issue string) {
	Complaint := &Complaint{
		From:  from.GetName(),
		About: about.GetName(),
		Issue: issue,
	}
	from.LeagueInfo().Commissioner.Complaints = append(from.LeagueInfo().Commissioner.Complaints, Complaint)

}
func createLeague(kind string) *League {
	//MAKE ROSTER TYPES BASED ON LEAGE TYPE
	LG := &League{
		Type:         kind,
		Commissioner: nil,
		Owners:       make([]*Owner, 0),
		MessBoard:    make([]*Message, 0),
		Game:         nil,
	}

	return LG
}
func (lg *League) AssignCommish(com *Commissioner) {
	lg.Commissioner = com
	com.League = lg
}

//Positions
const (
	WR = "Wide Receiver"
	QB = "Quarter Back"
	OL = "Offensive Lineman"
	RB = "Running Back"
	TE = "Tight End"
	DL = "Defensive Lineman"
	LB = "Line Backer"
	CB = "Corner Back"
	S  = "Safety"
)

//***********************     END LEAGUE      ***************************************************************************
//*
//*  									Interfaces
//*************************       BEGIN LEAGUE MEMBER   *****************************************************************
//LeagueMember Interface describes common functionality between all league members.
type LeagueMember interface {
	LeagueInfo() *League
	GetName() string
	SetSalary(amount float32)
	GetLevel() int
	Fine(amount float32)
	Pay()
	SendSlip(slip *Slip)
	CheckSuspension()
	GetType() string
	ToggleElig()
}

type Complaint struct {
	From  string
	About string
	Issue string
}

func CreateComplaint(from, about LeagueMember, commish *Commissioner, issue string) *Complaint {
	frm := from.GetName()
	abt := about.GetName()
	comp := &Complaint{
		From:  frm,
		About: abt,
		Issue: issue,
	}
	commish.Complaints = append(commish.Complaints, comp)
	return comp
}

func (complain *Complaint) ShowComplaint() string {
	return fmt.Sprintf(`
	
						From: %s
						About: %s
						Issue: %s
	
	`, complain.From, complain.About, complain.Issue)
}

type Eligible struct {
	Slips      []*Slip
	LMActive   bool
	Reason     string
	ReturnDate int64
}

const (
	DR      = "Doctor"
	Fine    = "Fine"
	Suspend = "Suspension"
	FB      = "Football"
	SOC     = "Soccer"
	BBALL   = "Basketball"
	BSBALL  = "Baseball"
)

//Message struct represents a social media post item
type Message struct {
	From    string
	Title   string
	Message string
	Visible bool
}

//Board struct is where Message structs are appended
type Board struct {
	Name  string
	Posts []*Message
}

/////////////////////////////////////////////
type LeagueType struct {
	League *League
	Type   string
}

type Game struct {
	Rules *Rules
}

type Rules struct {
	Penalty  int
	Points   int
	TeamSize int
}
