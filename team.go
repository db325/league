//Package Team Models A team in a league
package main

import (
	"fmt"
	//"net/http"
	//"log"
	//"strings"
)

func main() {
	commish := createCommish("Dajovan", "Bryant")
	owner := createOwner("mike", "will", "Head Hunters")
	gm := createManager("Killer", "Mike")
	gm.Type = "general manager"
	owner.SignGM(gm)
	//fmt.Println(gm)
	hc := createCoach("D", "Treez")
	gm.SignCoach(hc)
	//hc.Team = gm.Team
	//fmt.Println("Head Coach Team:--------->>> ", hc.Team.Name, "  *********gm team: ******** ", gm.Team.Coaches[0])
	//info := owner.GetInfo()
	CreateComplaint(hc, gm, commish, "he sucks")
	complain := commish.Complaints[0]
	fmt.Println(complain.ShowComplaint())
	jungo := CreateTrainer("Dajovan", "Bryant", 31, "Strength", 250.00)
	fmt.Println(jungo)

}
