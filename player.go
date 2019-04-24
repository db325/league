package main

import (
	"fmt"
	"time"
)

/***********************    BEGIN ATHELETE    ***********************************************************/

//Defines Athelete type to be used on/in leagues/teams.
type Athelete struct {
	*Team
	Salary         float32
	Level          int
	Atti           Attributes
	TeamName       string
	AccountBalance float32
	Eligible       *Eligible
	LMType         string
}

////////////////////////////////////////////////////////////////////

//Creates an Athelete that must be initialized with First and Last name values. All other values are modified after player creation.
//The default value for team is Free Agent until signed to a team.
func createPlayer(fname, lname, pos string) *Athelete {

	player1 := &Athelete{
		LMType:   "Player",
		TeamName: "Free Agent",
		Eligible: &Eligible{
			Reason:     "",
			Slips:      make([]*Slip, 10, 30),
			LMActive:   true,
			ReturnDate: 0,
		},
		Team: nil,
		Atti: Attributes{
			Firstname: fname,
			Lastname:  lname,
			Position:  pos,
		},
	}

	return player1
}

////////////////////////////////////////////////////////////////////
//Define structs for these methods
//Implement play
// func (player *Athelete) Play() {

// }

////////////////////////////////////////////////////////////////////

//implement train
func (player *Athelete) Train(trainer *Trainer) {
	player.AccountBalance -= trainer.Price
	switch trainer.Kind {
	case SA:
		player.Atti.Agility += 1
		player.Atti.Speed += 1
	case SC:
		player.Atti.Strength += 1
		player.Atti.Stamina += 1
	case AE:
		player.Atti.Accuracy += 1
		player.Atti.Willpower += 1

	}
}

////////////////////////////////////////////////////////////////////

func (player *Athelete) MediaPost(t, m string, v bool) {
	Message := &Message{
		Title:   t,
		Message: m,
		Visible: v,
	}
	if v == true {
		//IMPLEMENT LEAGUE BOARD
	} else if v == false {
		player.Team.MessBoard = append(player.Team.MessBoard, Message)
	}

}
func (player1 *Athelete) SetSalary(amount float32) {
	player1.Salary = amount
}

func (player *Athelete) GetLevel() int {
	return player.Level
}
func (player *Athelete) Pay() {
	PayAmount := player.Salary / 12
	player.AccountBalance += PayAmount

}

////////////////////////////////////////////////////////////////////

func (player *Athelete) Fine(amount float32) {
	player.SetSalary((player.Salary - amount))
}
func (player *Athelete) GetName() string {
	name := player.Atti.Firstname + " " + player.Atti.Lastname
	return name
}

//Attribute struct belonging to all atheletes. Must call createPlayer.
type Attributes struct {
	Firstname string
	Lastname  string
	Age       int
	DOB       string
	Height    float32
	Weight    float32
	Hometown  string
	Sport     string
	Superstar bool
	Starter   bool
	Strength  int
	Speed     int
	Stamina   int
	Accuracy  int
	Agility   int
	Willpower int
	JerseyNum int
	Position  string
	Team      string
}

func (player *Athelete) SetActive(yn bool) {
	if yn == true {
		player.Eligible.LMActive = true

	} else if yn == false {
		player.Eligible.LMActive = false
	}

}

func (player *Athelete) SendSlip(slip *Slip) {
	slip.SetTimeLeft()
	player.Eligible.Slips = append(player.Eligible.Slips, slip)
	return
}

// func (player *Athelete) GetSlips() []*Slip {
// 	for _, v := range player.Eligible.Slips {
// 		//
// 		if v.Time2Chech == time.Now() {
// 			//player.ToggleElig()
// 			fmt.Println("from get slip")
// 		}
// 		v.SetTimeLeft()
// 	}

// 	player.CheckSuspension()
// 	return player.Eligible.Slips

// }
func (player *Athelete) CheckSuspension() {

	if player.Eligible.LMActive == false {
		fmt.Println("just called")
		for i := 0; i < len(player.Eligible.Slips); i++ {

			if player.Eligible.Slips[i].SActive == true {
				//player.Eligible.Slips[i].SetTimeLeft()
				fmt.Println("true")
				//check time
				if player.Eligible.Slips[i].Time2Chech.Sub(time.Now()) < 0.0 {
					//fmt.Println("times up!")
					player.ToggleElig()
					player.Eligible.Slips[i].SActive = false
					player.Eligible.Slips[i].TimeLeft = 0.0
					//fmt.Println("Active")
				} else {
					fmt.Println("Banned Until ", player.Eligible.Slips[i].End)
				}
			}
		}
	} else {
		fmt.Println("Player Active")
	}
}

func (player *Athelete) GetType() string {
	return player.LMType
}
func (player *Athelete) ToggleElig() {
	player.Eligible.LMActive = !player.Eligible.LMActive

}

////////////////////////////////////////////////////////////////////

//***********************    END ATHELETE   ******************************************************************
