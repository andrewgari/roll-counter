package types

type ChatMessage struct {
	PlayerName string
	System     GameSystem
	Message    string
	Date       string
	Time       string
}

type RollMessage struct {
	*ChatMessage
	DieRoll  int
	ModRoll  int
	RollType RollType
	Result   RollResult
}
