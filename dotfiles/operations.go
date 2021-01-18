package dotfiles

import (
    "strings"
)

type OperationType int

const (
	Link OperationType = iota
	Unlink
	Unknown
)

func (d OperationType) String() string {
	return [...]string{"Link", "Unlink"}[d]
}

func NewOperationType(rawType string) OperationType {
	rawType = strings.ToLower(rawType)

	switch rawType {
	case "link":
		return Link
	case "unlink":
		return Unlink
	default:
		return Unknown
	}
}
