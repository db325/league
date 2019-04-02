package main

import (
	"fmt"
	//"time"
)

//************************    BEGIN LEAGUE    ************************************************************************

//League struct is a generic struct for a League object.
type League struct {
	LeagueType   struct{}
	Commissioner Commissioner
	Owners       []*Owner
}

//***********************     END LEAGUE      ***************************************************************************
//*
//*  									Interfaces
//*************************       BEGIN LEAGUE MEMBER   *****************************************************************
//LeagueMember Interface describes common functionality between all league members.
type LeagueMember interface {
	GetName() string
	SetSalary(amount float32)
	MediaPost(message *Message, board *Board)
	SetActive(yn bool)
	GetLevel() float32
	Fine(amount float32)
	Pay(amount float32)
	SendSlip(slip *Slip)
	GetSlips() []string
}

type Complaint struct {
	From  string
	About string
	Issue string
}

///////////////////////////////////////////////////////////////////
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

//////////////////////////////////////////////////////////////////
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
)
