package types

import "strings"

func IsKintos(name string) bool {
	return strings.Contains(name, "Sir Kintos")
}

func IsAine(name string) bool {
	return strings.Contains(name, "Aine")
}

func IsTree(name string) bool {
	return strings.Contains(name, "Tree") || strings.Contains(name, "Chett")
}

func IsFun(name string) bool {
	return strings.Contains(strings.ToLower(name), "fun")
}

func IsZed(name string) bool {
	return strings.Contains(name, "Zed")
}

func IsRo(name string) bool {
	return strings.Contains(name, "Ro")
}
