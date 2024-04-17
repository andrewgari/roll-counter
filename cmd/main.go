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

	dnd5eRolls := parser.ReadFile(types.DND_5e, "assets/dnd5e.log")
	pathfinderRolls := parser.ReadFile(types.PF2e, "assets/pathfinder.log")
	var chatMessages = append(dnd5eRolls, pathfinderRolls...)

	var total = 0
	var mapOfRolls = make(map[string]int)
	for _, message := range chatMessages {
		var count = mapOfRolls[message.ChatMessage.PlayerName]
		total++
		mapOfRolls[message.ChatMessage.PlayerName] = count + 1
	}

	//dieRolls := 0
	//modRolls := 0

	//partyMessages := parser.FormatMessages(chatMessages)
	partyMessages := parser.FormatMessages(chatMessages)
	fmt.Println(partyMessages)
	//for _, value := range partyMessages {
	//	for _, message := range value.Messages {
	//		fmt.Printf("%s had an average roll of %s", message.PlayerName, value.GetDiceAverage())
	//	}
	//}

	//for _, message := range partyMessages.AineRolls {
	//	dieRolls += message.DieRoll
	//	modRolls += message.ModRoll
	//}
	//
	//avgDieRoll := float64(dieRolls) / float64(len(partyMessages.AineRolls))
	//aveModRoll := float64(modRolls) / float64(len(partyMessages.AineRolls))F

	//fmt.Printf("Aine has an average die roll of %f and an average weighted roll of %f", avgDieRoll, aveModRoll)
	fmt.Println("Done")

}
