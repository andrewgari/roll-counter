package parser

import (
	"errors"
	"fmt"
	"github.com/andrewgari/roll-counter/internal/types"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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

		isGameNight, err := isMondayOrFriday(chatMessage.Date)
		if !isGameNight || err != nil {
			if err != nil {
				fmt.Println("Something went wrong", err.Error())
			}
			continue
		}

		isGameTime, err := isBetween8PMand1AM(chatMessage.Time)
		if !isGameTime || err != nil {
			//continue
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
	var timestamp = matches[2]
	var name = types.GetPlayerName(matches[7])

	var msg = types.ChatMessage{
		PlayerName: name,
		System:     types.DND_5e,
		Message:    message,
		Date:       date,
		Time:       timestamp,
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
		if dieRoll >= 17 {
			return types.SUCCESS
		}
		if dieRoll <= 5 {
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

func isMondayOrFriday(date string) (bool, error) {
	// Split the date string into month, day, and year
	dateParts := strings.Split(date, "/")
	if len(dateParts) != 3 {
		return false, fmt.Errorf("invalid date format")
	}

	// Add zero padding to the month and day if necessary
	month, err := strconv.Atoi(dateParts[0])
	if err != nil {
		return false, err
	}
	day, err := strconv.Atoi(dateParts[1])
	if err != nil {
		return false, err
	}
	year := dateParts[2]

	// Reconstruct the date string with zero padding
	date = fmt.Sprintf("%02d/%02d/%s", month, day, year)

	t, err := time.Parse("01/02/2006", date)
	if err != nil {
		return false, err
	}

	dayOfWeek := t.Weekday()
	return dayOfWeek == time.Monday || dayOfWeek == time.Friday, nil
}

func isBetween8PMand1AM(timeStr string) (bool, error) {
	// Parse the time string
	t, err := time.Parse("3:04:05 PM", timeStr)
	if err != nil {
		return false, err
	}

	// Load the EST time zone
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		return false, err
	}

	// Convert the time to EST
	t = t.In(location)

	// Get the hour
	hour := t.Hour()

	// Check if the hour is between 8 PM and 1 AM
	return hour >= 20 || hour < 1, nil
}
