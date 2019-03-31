package main

import ()

/***********************    BEGIN ATHELETE    ***********************************************************/

//Defines Athelete type to be used on/in leagues/teams.
type Athelete struct {
	*Team
	Salary         float32
	Level          float32
	Atti           Attributes
	TeamName       string
	AccountBalance float32
}

////////////////////////////////////////////////////////////////////

//Creates an Athelete that must be initialized with First and Last name values. All other values are modified after player creation.
//The default value for team is Undrafted if team field is empty
func createPlayer(fname, lname, team string) *Athelete {
	team = ""
	if team == "" {
		player1 := &Athelete{
			TeamName: "Undrafted",
		}
		return player1
	}
	player1 := &Athelete{
		Team: nil,
		Atti: Attributes{
			Firstname: fname,
			Lastname:  lname,
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

type Message struct {
	Title   string
	Message string
	Visible bool
}

type Board struct {
	Name  string
	Posts []*Message
}

func (player *Athelete) MediaPost(message *Message, board *Board) {
	board.Posts = append(board.Posts, message)

}
func (player1 *Athelete) SetSalary(amount float32) {
	player1.Salary = amount
}

func (player *Athelete) GetLevel() float32 {
	return player.Level
}

////////////////////////////////////////////////////////////////////

func (player *Athelete) Fine(amount float32) {
	player.SetSalary((player.Salary - amount))
}
func (player *Athelete) GetName() string {
	name := player.Atti.Firstname + " " + player.Atti.Lastname
	return name
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////

//***********************    END ATHELETE   ******************************************************************
