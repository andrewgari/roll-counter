package parser

import (
	"errors"
	"github.com/andrewgari/roll-counter/internal/types"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const headerRegex = `\[(\d{1,2}\/\d{1,2}\/\d{2,4}), ((0[1-9]|1[0-2]):([0-5][0-9]):([0-5][0-9]) (AM|PM))\] (.*)\n`
const verboseRoll5e = `[1,2]d20(kh|kl)? \d{1,2} \d{1,2} (\d{1,2})\n(\d{1,2})`
const terseRoll5e = `[1,2]d20(kh|kl)?.*= (\d{1,2}).*= (\d{1,2})`

func ReadFile(system types.GameSystem, fileName string) []types.RollMessage {
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		panic("Oh Shit the files Broke")
	}
	fileString := string(fileBytes[:])

	var chatMessages []types.RollMessage
	messages := strings.Split(fileString, "---------------------------")

	for _, message := range messages {
		chatMessage, err := parseMessage(message)
		if err != nil {
			continue
		}
		chatMessage.System = system
		rollMessage, err := parseRollMessage(&chatMessage)
		if err != nil {
			continue
		}
		chatMessages = append(chatMessages, rollMessage)
	}
	return chatMessages
}

func parseMessage(message string) (types.ChatMessage, error) {
	regex := regexp.MustCompile(headerRegex)
	matches := regex.FindStringSubmatch(message)

	if matches == nil {
		return types.ChatMessage{}, errors.New("no matches on message")
	}

	var date = matches[1]
	var time = matches[2]
	var name = types.GetPlayerName(matches[7])

	var msg = types.ChatMessage{
		PlayerName: name,
		System:     types.DND_5e,
		Message:    message,
		Date:       date,
		Time:       time,
	}
	return msg, nil
}

func parseRollMessage(message *types.ChatMessage) (types.RollMessage, error) {
	if message == nil {
		panic("message is nil")
	}
	var rollMessage = types.RollMessage{}
	rollMessage.ChatMessage = message

	dieRoll, err := parseDieRoll(message.Message)
	if err != nil {
		return types.RollMessage{}, err
	}
	if dieRoll != math.Trunc(dieRoll) {
		return types.RollMessage{}, errors.New("die roll is not a whole number")
	}
	rollMessage.DieRoll = dieRoll

	modRoll, err := parseModRoll(message.Message)
	if err != nil {
		return types.RollMessage{}, err
	}
	rollMessage.ModRoll = modRoll

	var result = parseRollResult(dieRoll, message.Message)
	rollMessage.Result = result

	var rollType = parseRollType(message.Message)
	rollMessage.RollType = rollType

	return rollMessage, nil
}

func parseDieRoll(message string) (float64, error) {
	regex := regexp.MustCompile(terseRoll5e)
	matches := regex.FindStringSubmatch(message)
	if matches == nil {
		regex = regexp.MustCompile(verboseRoll5e)
		matches = regex.FindStringSubmatch(message)
	}

	if matches == nil {
		return -1, errors.New("not a roll message")
	}

	return strconv.ParseFloat(matches[2], 64)
}

func parseModRoll(message string) (float64, error) {
	regex := regexp.MustCompile(terseRoll5e)
	matches := regex.FindStringSubmatch(message)
	if matches == nil {
		regex = regexp.MustCompile(verboseRoll5e)
		matches = regex.FindStringSubmatch(message)
	}

	if matches == nil {
		return -1, errors.New("not a roll message")
	}

	return strconv.ParseFloat(matches[3], 64)
}

func parseRollResult(dieRoll float64, message string) types.RollResult {
	if dieRoll == 20 {
		return types.CRITICAL
	}
	if dieRoll == 1 {
		return types.FUMBLE
	}

	regex := regexp.MustCompile(`hits|misses`)
	matches := regex.FindStringSubmatch(message)

	if matches == nil {
		if dieRoll > 10 {
			return types.SUCCESS
		}
		if dieRoll <= 10 {
			return types.FAILURE
		}
		return types.UNKNOWN
	}

	switch matches[0] {
	case "hits":
		return types.SUCCESS
	case "misses":
		return types.FAILURE
	default:
		return types.UNKNOWN // then what is it?
	}
}

func parseRollType(message string) types.RollType {
	regex := regexp.MustCompile(`[1,2]d20(kh|kl)?`)
	matches := regex.FindStringSubmatch(message)

	switch matches[1] {
	case "kh":
		return types.ADVANTAGE
	case "kl":
		return types.DISADVANTAGE
	default:
		return types.NORMAL
	}
}
