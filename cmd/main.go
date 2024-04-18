package main

import (
	"fmt"
	parser "github.com/andrewgari/roll-counter/internal"
	"github.com/andrewgari/roll-counter/internal/types"
)

func main() {
	readText()
}

func readText() {

	dnd5eRolls := parser.ReadFile(types.DND_5e, "assets/pathfinder.log")
	//pathfinderRolls := parser.ReadFile(types.PF2e, "assets/pathfinder.log")
	//var chatMessages = append(dnd5eRolls, pathfinderRolls...)
	var chatMessages = dnd5eRolls

	var total = 0
	var playerRolls = parser.CreatePartyMessages()
	for _, message := range chatMessages {
		total++
		playerRolls.AddMessage(message)
	}

	printRolls(types.Aine, playerRolls)
	printRolls(types.Kintos, playerRolls)
	printRolls(types.Tree, playerRolls)
	printRolls(types.Fun, playerRolls)
	printRolls(types.Zed, playerRolls)
	printRolls(types.Rowan, playerRolls)
	printRolls(types.GOD, playerRolls)

	fmt.Println("Done")

}

func printRolls(name types.PlayerName, playerRolls parser.PartyMessages) {
	totalRolls := playerRolls.GetMessageCount(name)
	rollAvg := playerRolls.GetDiceAverage(name)
	modAvg := playerRolls.GetModAverage(name)
	successAvg := playerRolls.GetSuccessAverage(name)
	failureAvg := playerRolls.GetFailureAverage(name)

	fmt.Printf("%s rolled a total of %d rolls, with an average die roll of %f, with an average total roll of %f. Their average success rate was %f, and their failure rate was %f\n\n", name.String(), totalRolls, rollAvg, modAvg, successAvg, failureAvg)
}
