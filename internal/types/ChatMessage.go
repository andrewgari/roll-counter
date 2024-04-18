package types

type ChatMessage struct {
	PlayerName PlayerName
	System     GameSystem
	Message    string
	Date       string
	Time       string
}

func (cm ChatMessage) GetPlayerName() {

}

type RollMessage struct {
	*ChatMessage
	DieRoll  float64
	ModRoll  float64
	RollType RollType
	Result   RollResult
}
