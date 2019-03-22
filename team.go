//Package Team Models A team in a league
package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func main() {
	p1 := createPlayer("mr", "bryant", "hithere")
	fmt.Println(p1.Team)
	ch := createCoach("ty", "law", "players coach")
	tm, err := createTeam("noles")
	if err != nil {
		log.Println(err)
	}
	li := *p1
	ch.Sign(p1.Atti.Firstname+" "+p1.Atti.Lastname, "QB", tm, li, 99)
	fmt.Println(*tm.Roster["QB"], tm.Players[0])
}

type Person struct {
	ID int
}

//*********************** BEGIN TEAM **********************

//Team object. You must call createTeam and initialize it with a name
type Team struct {
	Name            string
	UpperManagement []*Manager
	Players         []*Athelete
	Coaches         []*Coach
	Roster          map[string]*Athelete
}

//Returns a pointer to a Team struct, initialized with a name
func createTeam(name string) (*Team, error) {
	m := make(map[string]*Athelete)
	if name == "" {
		err := errors.New("Enter a name.")
		return nil, err
	} else {
		team := &Team{
			Name:   name,
			Roster: m,
		}
		return team, nil
	}

}

//*********************** END TEAM **********************

type Manager struct {
	Level     float32
	FirstName string
	LastName  string
	Type      string
	Team      string
	CanHire   bool
	CanFire   bool
	Person
}

func createManager(fname, lname string) *Manager {
	man := &Manager{
		FirstName: fname,
		LastName:  lname,
	}
	return man
}

func (man *Manager) SignHeadCoach(coach *Coach, team *Team) error {
	upper := strings.ToUpper("head coach")
	if coach.Atti.TypeOfCoach != upper {
		flag := errors.New("FLAG!!! You Can Not Sign This Person!")
		return flag
	} else {
		team.Coaches = append(team.Coaches, coach)
	}
	return nil
}

//TEST MANAGER

//***********************Begin Coach**********************

//Defines COACH type to be used on/in leagues/teams.
type Coach struct {
	Level float32
	Atti  CAttributes
	Team  string
}

// Values for Coach object
type CAttributes struct {
	ID int
	Person
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
}

//Creates a Coach object with at least the first and last name. The team is optional. Usually upon creation,
//you will not set the team property. If left emty, it will default to n/a.
func createCoach(Fname, Lname, team string) *Coach {

	if team == "" {

		team = "n/a"
	} else {
		team = team
	}
	coach := &Coach{
		Atti: CAttributes{
			FirstName: Fname,
			LastName:  Lname,
		},
		Team: team,
	}

	return coach
}

//Sign function adds Athelete to a roster. It also makes the player.Team variable equal to team.Name
func (coach *Coach) Sign(name, pos string, team *Team, player Athelete, jersey int) {
	team.Players = append(team.Players, &player)
	nm1 := fmt.Sprintf("%s\t%s", player.Atti.Firstname, player.Atti.Lastname)
	player.Atti.JerseyNum = jersey
	player.Team = team.Name
	name = nm1
	player = (player)
	team.Roster[pos] = &player

}
func (coach *Coach) Cut(name, pos string, team *Team, player *Athelete) {
	nm1 := fmt.Sprintf("%s\t%s", player.Atti.Firstname, player.Atti.Lastname)
	name = nm1
	team.Roster[pos] = nil
	player.Team = "Free Agent"
}

//***********************End Coach**********************

//***********************Begin Athelete**********************

//Defines Athelete type to be used on/in leagues/teams.
type Athelete struct {
	Level float32
	Person
	Atti Attributes
	Team string
}

//Creates an Athelete that must be initialized with First and Last name values. All other values are modified after player creation.
//The default value for team is Undrafted if team field is empty
func createPlayer(fname, lname, team string) *Athelete {
	if team == "" {
		player1 := &Athelete{
			Team: "Undrafted",
		}
		return player1
	}
	player1 := &Athelete{
		Team: team,
		Atti: Attributes{
			Firstname: fname,
			Lastname:  lname,
		},
	}

	return player1
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
	Accuracy  int
	Agility   int
	Willpower int
	JerseyNum int
	Position  string
	Team      string
}

//***********************End Athelete **********************
