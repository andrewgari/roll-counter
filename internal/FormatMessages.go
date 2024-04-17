package parser

import "github.com/andrewgari/roll-counter/internal/types"

type PartyMessages struct {
	AineRolls   CharacterMessages
	KintosRolls CharacterMessages
	TreeRolls   CharacterMessages
	FunRolls    CharacterMessages
	ZedRolls    CharacterMessages
	RoRolls     CharacterMessages
	GodRolls    CharacterMessages
}

type CharacterMessages struct {
	Messages []types.RollMessage
}

func (p PartyMessages) addMessage(message types.RollMessage) {
	switch message.PlayerName {
	case types.Kintos:
		p.KintosRolls.addMessage(message)
		break
	case types.Aine:
		p.AineRolls.addMessage(message)
		break
	case types.Fun:
		p.FunRolls.addMessage(message)
		break
	case types.Zed:
		p.ZedRolls.addMessage(message)
		break
	case types.Tree:
		p.TreeRolls.addMessage(message)
		break
	case types.Rowan:
		p.RoRolls.addMessage(message)
		break
	default:
		p.GodRolls.addMessage(message)
		break
	}
}

func (c CharacterMessages) addMessage(message types.RollMessage) {
	c.Messages = append(c.Messages, message)
}

func (c CharacterMessages) GetDiceAverage() float64 {
	dieRolls := 0
	for _, message := range c.Messages {
		dieRolls += message.DieRoll
	}
	return float64(dieRolls) / float64(len(c.Messages))
}

func (c CharacterMessages) GetModAverage() float64 {
	dieRolls := 0
	for _, message := range c.Messages {
		dieRolls += message.ModRoll
	}
	return float64(dieRolls) / float64(len(c.Messages))
}

func (c CharacterMessages) GetSuccessAverage() int {
	numOfSuccess := 0
	for _, message := range c.Messages {
		if message.Result == types.SUCCESS || message.Result == types.CRITICAL {
			numOfSuccess++
		}
	}
	return numOfSuccess / len(c.Messages)
}

func (c CharacterMessages) GetFailureAverage() int {
	numOfSuccess := 0
	for _, message := range c.Messages {
		if message.Result == types.FAILURE || message.Result == types.FUMBLE {
			numOfSuccess++
		}
	}
	return numOfSuccess / len(c.Messages)
}

func FormatMessages(messages []types.RollMessage) PartyMessages {
	var partyMessages PartyMessages
	for _, message := range messages {
		switch message.PlayerName {
		case types.Kintos:
			partyMessages.KintosRolls.addMessage(message)
			break
		case types.Aine:
			partyMessages.AineRolls.addMessage(message)
			break
		case types.Tree:
			partyMessages.TreeRolls.addMessage(message)
			break
		case types.Fun:
			partyMessages.FunRolls.addMessage(message)
			break
		case types.Rowan:
			partyMessages.RoRolls.addMessage(message)
			break
		case types.Zed:
			partyMessages.ZedRolls.addMessage(message)
			break
		default:
			partyMessages.GodRolls.addMessage(message)
			break
		}
	}
	return partyMessages
}
