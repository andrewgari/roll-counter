package types

type RollResult int

const (
	CRITICAL RollResult = iota
	SUCCESS
	FAILURE
	FUMBLE
)
