package types

import "regexp"

type PlayerName int

const (
	Kintos PlayerName = iota
	Aine
	Tree
	Fun
	Rowan
	Zed
	GOD
)

func (pn PlayerName) String() string {
	return []string{
		"Sir Kintos Vanhausen Krinleback the 3rd, Holy Knight of the Albion Kingdom under the Golden Sun",
		"Aine Vicis",
		"Chett Slitherbottom",
		"Fun!",
		"Rowan 'Ro' Goldsnatch",
		"Zed",
		"The GM",
	}[pn]
}

func GetPlayerName(name string) PlayerName {
	if IsKintos(name) {
		return Kintos
	}
	if IsFun(name) {
		return Fun
	}
	if IsAine(name) {
		return Aine
	}
	if IsTree(name) {
		return Tree
	}
	if IsZed(name) {
		return Zed
	}
	if IsRo(name) {
		return Rowan
	}
	return GOD
}

func IsKintos(name string) bool {
	regex := regexp.MustCompile(`(?i)\bkintos\b`)
	return regex.MatchString(name)
}

func IsFun(name string) bool {
	regex := regexp.MustCompile(`(?i)\bfun/??\b`)
	return regex.MatchString(name)
}

func IsAine(name string) bool {
	regex := regexp.MustCompile(`(?i)\baine\b`)
	return regex.MatchString(name)
}

func IsTree(name string) bool {
	regex := regexp.MustCompile(`(?i)\btree|chett\b`)
	return regex.MatchString(name)
}

func IsRo(name string) bool {
	regex := regexp.MustCompile(`(?i)\browan\b`)
	return regex.MatchString(name)
}

func IsZed(name string) bool {
	regex := regexp.MustCompile(`(?i)\bzed\b`)
	return regex.MatchString(name)
}
