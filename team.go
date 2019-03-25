//Package Team Models A team in a league
package main

import (
	"errors"
	"fmt"
	//"log"
	"strings"
)

func main() {
	c1 := createCoach("dae-dae", "B", "Recievers Coach")
	c2 := createCoach("d", "b", "head coach")
	p1 := createPlayer("clay", "matthews", "jumpers")
	p2 := createPlayer("db", "hu", "io")
	c2.MakeTradeReq(c2, p1, 55.88, 450000.00, false, c1, p2)

	fmt.Println(c1.Atti.Requests[0].ShowRequest())

}

//*
//*
//********************************    BEGIN MANAGER    **************************************************
type Manager struct {
	Level     float32
	FirstName string
	LastName  string
	Type      string
	Team      string
	CanHire   bool
	CanFire   bool
	Salary    float32
}

//////////////////////////////////////////////////////////////////////////
func createManager(fname, lname string) *Manager {
	man := &Manager{
		FirstName: fname,
		LastName:  lname,
	}
	return man
}

//////////////////////////////////////////////////////////////////
func (man *Manager) SignHeadCoach(coach *Coach, team *Team) error {
	genmanager := strings.ToUpper(man.Type)
	if coach.Atti.TypeOfCoach != "Head Coach" && man.Type != genmanager {
		flag := errors.New("FLAG!!! You Can Not Sign This Person!")
		return flag
	} else {
		team.Coaches = append(team.Coaches, coach)
	}
	return nil
}

/////////////////////////////////////////////////////////////////
func (manager *Manager) Trade(player *Athelete, team *Team) {

}

////////////////////////////////////////////////////////////////
//************************    END MANAGER    **********************************************************
//*
//*
//***********************    BEGIN COACH    ***********************************************************

//Defines COACH type to be used on/in leagues/teams.
type Coach struct {
	Salary float32
	Level  float32
	Atti   CAttributes
	Team   string
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

////////////////////////////////////////////////////////////////////

//Sign function adds Athelete to a roster. It also makes the player.Team variable equal to team.Name by default.
func (coach *Coach) Sign(name, pos string, team *Team, player Athelete, jersey int) {
	team.Players = append(team.Players, &player)
	nm1 := fmt.Sprintf("%s\t%s", player.Atti.Firstname, player.Atti.Lastname)
	player.Atti.JerseyNum = jersey
	player.Team = team.Name
	name = nm1
	player = (player)
	team.Roster[pos] = &player

}

////////////////////////////////////////////////////////////////////

func (coach *Coach) Cut(name, pos string, team *Team, player *Athelete) {
	nm1 := fmt.Sprintf("%s\t%s", player.Atti.Firstname, player.Atti.Lastname)
	name = nm1
	team.Roster[pos] = nil
	player.Team = "Free Agent"
}

////////////////////////////////////////////////////////////////////

func (coach *Coach) GiveSpeach() {

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

////////////////////////////////////////////////////////////////////

//***********************    END COACH    ***************************************************************
//*
//*
//***********************    BEGIN ATHELETE    ***********************************************************

//Defines Athelete type to be used on/in leagues/teams.
type Athelete struct {
	Salary float32
	Level  float32
	Atti   Attributes
	Team   string
}

////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////

//Implement play
func (player *Athelete) Play() {

}

////////////////////////////////////////////////////////////////////

//implement train
func (player *Athelete) Train() {

}

////////////////////////////////////////////////////////////////////

func (player *Athelete) Tweet() {

}

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
	Accuracy  int
	Agility   int
	Willpower int
	JerseyNum int
	Position  string
	Team      string
}

////////////////////////////////////////////////////////////////////

//***********************    END ATHELETE   ****************************************************************
//*
//*
//***********************    BEGIN OWNER    ******************************************************************

type Owner struct {
	Level     float32
	FirstName string
	LastName  string
	Team      *Team
	Age       int
	Salary    float32
}

////////////////////////////////////////////////////////////////////

func createOwner(fname, lname, tname string) *Owner {

	owner := &Owner{

		FirstName: fname,
		LastName:  lname,
	}
	kl, _ := owner.createTeam(tname)
	owner.Team = kl
	return owner
}

////////////////////////////////////////////////////////////////////

//Returns a pointer to a Team struct, initialized with a name.
func (owner *Owner) createTeam(name string) (*Team, error) {
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

////////////////////////////////////////////////////////////////////

func (owner *Owner) PayPeople(lm LeagueMember, amount float32) {

}

////////////////////////////////////////////////////////////////////

func (owner *Owner) Tweet() {

}

////////////////////////////////////////////////////////////////////

//****************************    End  Owner    ******************************************************************
//*
//*
//***********************   BEGIN TEAM   *********************************************************************

//Team object. You must call createTeam and initialize it with a name. It is called in the createOwner function by default.
type Team struct {
	Name            string
	UpperManagement []*Manager
	Players         []*Athelete
	Coaches         []*Coach
	Roster          map[string]*Athelete
}

//***********************   END TEAM   ***********************************************************************
//*
//*
//***********************    BEGIN COMMISSIONER    ***************************************************************
type Commissioner struct {
	Level     float32
	FirstName string
	LastName  string
	Age       int
}

////////////////////////////////////////////////////////////////////

func createCommish(fname, lname string) *Commissioner {

	commish := &Commissioner{
		FirstName: fname,
		LastName:  lname,
	}
	return commish
}

////////////////////////////////////////////////////////////////////

func (comish *Commissioner) Fine(lm LeagueMember, amount float32, reason string) {
	lvl := lm.GetLevel()
	switch lvl {
	case 5:
		lm.Fine(54.44)
		fmt.Println(lm)

	}
}

////////////////////////////////////////////////////////////////////
//*
//*
//************************    BEGIN LEAGUE    ************************************************************************

//League struct is a generic struct for a League object.
type League struct {
	LeagueType   struct{}
	Commissioner Commissioner
	Owners       []*Owner
}

//***********************     END LEAGUE      ***************************************************************************
//*
//*
//*************************       BEGIN LEAGUE MEMBER   *****************************************************************
//LeagueMember Interface describes common functionality between all league members.
type LeagueMember interface {
	SetSalary(amount float32)
	Tweet()
	GetLevel() float32
	Fine(amount float32)
}

////////////////////////////////////////////////////////////////////

func (player *Athelete) GetLevel() float32 {
	return player.Level
}

////////////////////////////////////////////////////////////////////

func (player *Athelete) SetSalary(amount float32) {
	player.Salary = amount

}

////////////////////////////////////////////////////////////////////

func (player *Athelete) Fine(amount float32) {
	player.SetSalary((player.Salary - amount))
}

////////////////////////////////////////////////////////////////////

func (manager *Manager) GetLevel() float32 {
	return manager.Level
}

////////////////////////////////////////////////////////////////////

func (manager *Manager) SetSalary(amount float32) {
	manager.Salary = amount

}

////////////////////////////////////////////////////////////////////

func (manager *Manager) Fine(amount float32) {
	manager.Salary += manager.Salary - amount
}

////////////////////////////////////////////////////////////////////

func (coach *Coach) GetLevel() float32 {
	return coach.Level
}

////////////////////////////////////////////////////////////////////

func (coach *Coach) SetSalary(amount float32) {
	coach.Salary = amount

}

////////////////////////////////////////////////////////////////////

func (coach *Coach) Fine(amount float32) {
	coach.Salary += coach.Salary - amount
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////

func (commish *Commissioner) GetLevel() float32 {
	return commish.Level
}

////////////////////////////////////////////////////////////////////

func (owner *Owner) GetLevel() float32 {
	return owner.Level
}

////////////////////////////////////////////////////////////////////

func (owner *Owner) SetSalary(amount float32) {
	owner.Salary = amount

}

////////////////////////////////////////////////////////////////////

func (owner *Owner) Fine(amount float32) {
	owner.Salary += owner.Salary - amount
}

////////////////////////////////////////////////////////////////////

//*********************    End League Member    ******************************************************************
//**
//**

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

	 
	 				Approved :%v




`, strings.ToUpper(tr.Coach.Atti.FirstName+" "+tr.Coach.Atti.LastName), strings.ToUpper(tr.ProposeThis.Atti.Firstname+" "+tr.ProposeThis.Atti.Lastname), float32(tr.AndOrThis1), strings.ToUpper(tr.ForThis.Atti.Firstname+" "+tr.ForThis.Atti.Lastname), tr.AndOrThis2, false)
	return req
}

////////////////////////////////////////////////////////////////////
