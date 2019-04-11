package main

import (
	//"fmt"
	"time"
)

func main() {
	//now := time.Now()
	com := createCommish("don", "draper")
	own := createOwner("the", "fans", "hot-shots")
	man := createManager("harry", "thompson")
	man.Type = "general manager"
	own.SignGM(man)
	c1 := createCoach("mr", "johns")
	man.SignCoach(c1)
	p1 := createPlayer("d", "b", "k")
	c1.Sign("rec", "lo", c1.Team, p1, 80)

	c2 := createCoach("mr", "wills")
	c2.MakeTradeReq(c2, p1, 80000.00, 500000.00, false, c1, nil)
	//fmt.Println(com, "\n", own, "\n", man, "\n", c1.Atti.Requests[0], "\n")
	//com.Suspend(c2, 1, "i dont like u")
	//jk := p1.GetSlips()
	// srt := p1.GetSlips()
	//now := time.Now()
	//fmt.Println(now)
	com.GetLevel()
	//com.Suspend(own, 4, "cause i can bitch!!")
	com.Suspend(p1, 302, "jkjkj")
	//slips := c1.GetSlips()
	if c1.Eligible.LMActive == false {

		c1.Eligible.Reason = c1.Eligible.Slips[0].Reason
	}
	//fmt.Println(c1.Eligible.Reason)
	//p1arr := p1.GetSlips()
	p1.CheckSuspension()
	time.Sleep(time.Minute * 5)
	p1.CheckSuspension()

}
