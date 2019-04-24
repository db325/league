package main

import (
	"fmt"
	"time"
)

func main() {
	//now := time.Now()
	//com := createCommish("don", "draper")
	own, err := createOwner("the", "fans", "hot-shots")
	if err != nil {
		fmt.Print(err)
	}
	gm := createManager("mr", "nobody")
	gm.LMType = "general manager"

	own.Sign(gm)
	//fmt.Println(own.Team.UpperManagement[0].GetName())
	coach := createCoach("dae dae", "Bryant")
	coach.LMType = "head coach"
	gm.Sign(coach, nil)
	//fmt.Println(coach.Team.UpperManagement, "  ", gm.Team.Coaches[0])
	p1 := createPlayer("john", "wick", "QB")
	fmt.Println(p1)
	coach.Sign(p1, nil)

	//fmt.Println(coach.Team.Roster["QB"])
	// p1.CheckSuspension()
	// time.Sleep(time.Minute * 5)
	// p1.CheckSuspension()

	//fmt.Println(p1.Team)

	own.Suspend(p1, 5, "cuz")
	own.Suspend(coach, 7, "he's an ass")
	fmt.Println(p1.Eligible.Slips[0])
	fmt.Println(coach.Eligible.Slips[0])
	//	fmt.Println(gm.Eligible.Slips[0])
	p1.CheckSuspension()
	coach.CheckSuspension()
	//p1.CheckSuspension()

	time.Sleep(time.Minute * 3)
	gm.CheckSuspension()
	coach.CheckSuspension()

	//	fmt.Println(coach.Eligible.Slips[0])
	time.Sleep(time.Minute * 3)
	p1.CheckSuspension()
	coach.CheckSuspension()

	p1.CheckSuspension()
	coach.CheckSuspension()

	fmt.Println(p1.Eligible.Slips[0])

	//fmt.Println(p1.Eligible.Slips[0])

}
