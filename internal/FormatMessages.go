package parser

import "github.com/andrewgari/roll-counter/internal/types"

type PartyMessages struct {
	AineRolls   []types.RollMessage
	KintosRolls []types.RollMessage
	TreeRolls   []types.RollMessage
	FunRolls    []types.RollMessage
	ZedRolls    []types.RollMessage
	RoRolls     []types.RollMessage
}

type PartyMessages2 struct {
	Messages []types.RollMessage
}

func (c PartyMessages2) GetDiceAverage() float64 {
	dieRolls := 0
	for _, message := range c.Messages {
		dieRolls += message.DieRoll
	}
	return float64(dieRolls) / float64(len(c.Messages))
}

func (c PartyMessages2) GetModAverage() float64 {
	dieRolls := 0
	for _, message := range c.Messages {
		dieRolls += message.ModRoll
	}
	return float64(dieRolls) / float64(len(c.Messages))
}

func (c PartyMessages2) GetSuccessAverage() int {
	numOfSuccess := 0
	for _, message := range c.Messages {
		if message.Result == types.SUCCESS || message.Result == types.CRITICAL {
			numOfSuccess++
		}
	}
	return numOfSuccess / len(c.Messages)
}

func (c PartyMessages2) GetFailureAverage() int {
	numOfSuccess := 0
	for _, message := range c.Messages {
		if message.Result == types.FAILURE || message.Result == types.FUMBLE {
			numOfSuccess++
		}
	}
	return numOfSuccess / len(c.Messages)
}

func FormatMessages(messages []types.RollMessage) PartyMessages {
	var partyMessages = PartyMessages{
		AineRolls:   make([]types.RollMessage, 0),
		KintosRolls: make([]types.RollMessage, 0),
		TreeRolls:   make([]types.RollMessage, 0),
		FunRolls:    make([]types.RollMessage, 0),
		ZedRolls:    make([]types.RollMessage, 0),
		RoRolls:     make([]types.RollMessage, 0),
	}
	for _, message := range messages {
		if types.IsKintos(message.PlayerName) {
			partyMessages.KintosRolls = append(partyMessages.KintosRolls, message)
		}

		if types.IsFun(message.PlayerName) {
			partyMessages.FunRolls = append(partyMessages.FunRolls, message)
		}

		if types.IsAine(message.PlayerName) {
			partyMessages.AineRolls = append(partyMessages.AineRolls, message)
		}

		if types.IsZed(message.PlayerName) {
			partyMessages.ZedRolls = append(partyMessages.ZedRolls, message)
		}

		if types.IsRo(message.PlayerName) {
			partyMessages.RoRolls = append(partyMessages.RoRolls, message)
		}

		if types.IsTree(message.PlayerName) {
			partyMessages.TreeRolls = append(partyMessages.TreeRolls, message)
		}
	}
	return partyMessages
}
