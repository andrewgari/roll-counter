package parser

import "github.com/andrewgari/roll-counter/internal/types"

type PartyMessages struct {
	AineRolls   []types.RollMessage
	KintosRolls []types.RollMessage
	TreeRolls   []types.RollMessage
	FunRolls    []types.RollMessage
	ZedRolls    []types.RollMessage
	RoRolls     []types.RollMessage
	GodRolls    []types.RollMessage
}

func CreatePartyMessages() PartyMessages {
	var p PartyMessages
	p.KintosRolls = make([]types.RollMessage, 0)
	p.AineRolls = make([]types.RollMessage, 0)
	p.FunRolls = make([]types.RollMessage, 0)
	p.ZedRolls = make([]types.RollMessage, 0)
	p.TreeRolls = make([]types.RollMessage, 0)
	p.RoRolls = make([]types.RollMessage, 0)
	p.GodRolls = make([]types.RollMessage, 0)

	return p
}

func (p *PartyMessages) AddMessage(message types.RollMessage) {
	switch message.ChatMessage.PlayerName {
	case types.Kintos:
		p.KintosRolls = append(p.KintosRolls, message)
		break
	case types.Aine:
		p.AineRolls = append(p.AineRolls, message)
		break
	case types.Fun:
		p.FunRolls = append(p.FunRolls, message)
		break
	case types.Zed:
		p.ZedRolls = append(p.ZedRolls, message)
		break
	case types.Tree:
		p.TreeRolls = append(p.TreeRolls, message)
		break
	case types.Rowan:
		p.RoRolls = append(p.RoRolls, message)
		break
	default:
		p.GodRolls = append(p.GodRolls, message)
		break
	}
}

func (p *PartyMessages) getPlayerMessages(name types.PlayerName) []types.RollMessage {
	switch name {
	case types.Kintos:
		return p.KintosRolls
	case types.Aine:
		return p.AineRolls
	case types.Fun:
		return p.FunRolls
	case types.Zed:
		return p.ZedRolls
	case types.Tree:
		return p.TreeRolls
	case types.Rowan:
		return p.RoRolls
	default:
		return p.GodRolls
	}
}

func (p *PartyMessages) GetMessageCount(name types.PlayerName) int {
	return len(p.getPlayerMessages(name))
}

func (p *PartyMessages) GetDiceAverage(name types.PlayerName) float64 {
	dieRolls := 0.0
	var messages = p.getPlayerMessages(name)
	for _, message := range messages {
		dieRolls += message.DieRoll
	}
	return dieRolls / float64(len(messages))
}

func (p *PartyMessages) GetModAverage(name types.PlayerName) float64 {
	dieRolls := 0.0
	var messages = p.getPlayerMessages(name)
	for _, message := range messages {
		dieRolls += message.ModRoll
	}
	return dieRolls / float64(len(messages))
}

func getNumOfSuccess(messages []types.RollMessage) float64 {
	numOfSuccess := 0.0
	for _, message := range messages {
		if message.Result == types.SUCCESS || message.Result == types.CRITICAL {
			numOfSuccess++
		}
	}
	return numOfSuccess
}

func getNumOfFailure(messages []types.RollMessage) float64 {
	numOfFailure := 0.0
	for _, message := range messages {
		if message.Result == types.FAILURE || message.Result == types.FUMBLE {
			numOfFailure++
		}
	}
	return numOfFailure
}

func (p *PartyMessages) GetSuccessAverage(name types.PlayerName) float64 {
	numOfSuccess := getNumOfSuccess(p.getPlayerMessages(name))
	numOfFailure := getNumOfFailure(p.getPlayerMessages(name))

	if numOfFailure == 0 {
		return 0
	}

	return (numOfSuccess / numOfFailure) * 1
}

func (p *PartyMessages) GetFailureAverage(name types.PlayerName) float64 {
	numOfSuccess := getNumOfSuccess(p.getPlayerMessages(name))
	numOfFailure := getNumOfFailure(p.getPlayerMessages(name))
	if numOfSuccess == 0 {
		return 0
	}
	return (numOfFailure / numOfSuccess) * 1
}
