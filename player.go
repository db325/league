package main

import (
	"fmt"
	//"strings"
	"time"
)

/***********************    BEGIN ATHELETE    ***********************************************************/

//Defines Athelete type to be used on/in leagues/teams.
type Athelete struct {
	*League
	*Team
	*Coach
	Salary         float32
	Level          int
	Atti           Attributes
	TeamName       string
	AccountBalance float32
	Eligible       *Eligible
	LMType         string
	Firstname      string
	Lastname       string
	Age            int
	DOB            string
	Height         float32
}

////////////////////////////////////////////////////////////////////

//Creates an Athelete that must be initialized with First and Last name values. All other values are modified after player creation.
//The default value for team is Free Agent until signed to a team.
func createPlayer(fname, lname, pos string) *Athelete {

	player1 := &Athelete{

		Firstname: fname,
		Lastname:  lname,
		LMType:    "player",
		TeamName:  "Free Agent",
		Eligible: &Eligible{
			Reason:     "",
			Slips:      make([]*Slip, 10, 30),
			LMActive:   true,
			ReturnDate: 0,
		},
		Atti: Attributes{
			Position: pos,
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

//Adjusts attributes based on Trainer type.
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

//Posts message to a message board.
func (player *Athelete) MediaPost(t, m string, v bool) {
	Message := &Message{
		From:    player.GetName(),
		Title:   t,
		Message: m,
		Visible: v,
	}
	if v == true {
		//IMPLEMENT LEAGUE BOARD
		player.League.MessBoard = append(player.League.MessBoard, Message)
	} else if v == false {
		player.Team.MessBoard = append(player.Team.MessBoard, Message)
	}

}

//Sets salary variable.
func (player1 *Athelete) SetSalary(amount float32) {
	player1.Salary = amount
}

//Returns level.
func (player *Athelete) GetLevel() int {
	return player.Level
}

//Adds payamount to AccountBalance.
func (player *Athelete) Pay() {
	PayAmount := player.Salary / 12
	player.AccountBalance += PayAmount

}

//Subtracts amount from AccountBalance.
func (player *Athelete) Fine(amount float32) {
	player.AccountBalance -= amount
}

//Returns name.
func (player *Athelete) GetName() string {
	name := player.Firstname + " " + player.Lastname
	return name
}

//Attribute struct belonging to all atheletes. Must call createPlayer.
type Attributes struct {
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

//Puts a slip in Slips.
func (player *Athelete) SendSlip(slip *Slip) {
	slip.SetTimeLeft()
	player.Eligible.Slips = append(player.Eligible.Slips, slip)

}

//Checks eligibilty.
func (player *Athelete) CheckSuspension() {

	if player.Eligible.LMActive == false {
		fmt.Println("just called")
		for i := 0; i < len(player.Eligible.Slips); i++ {

			if player.Eligible.Slips[i].SActive == true {
				fmt.Println("true")
				//check time
				if player.Eligible.Slips[i].Time2Chech.Sub(time.Now()) < 0.0 {
					player.ToggleElig()
					player.Eligible.Slips[i].SActive = false
					player.Eligible.Slips[i].TimeLeft = 0.0

				} else {
					fmt.Println("Banned Until ", player.Eligible.Slips[i].End)
				}
			}
		}
	} else {
		fmt.Println("Player Active")
	}
}

//Returns league member role
func (player *Athelete) GetType() string {
	return player.LMType
}

//Toggles eligibilty.
func (player *Athelete) ToggleElig() {
	player.Eligible.LMActive = !player.Eligible.LMActive

}

//***********************    END ATHELETE   ******************************************************************
