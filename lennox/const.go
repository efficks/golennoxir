package lennox

import (
	"errors"
)

type FanSpeed int

const (
	FAN_NONE FanSpeed = 0 + iota
	FAN_MINIMUM
	FAN_MEDIUM
	FAN_MAXIMUM
	FAN_AUTO
)

func ToFanSpeed(s string) (FanSpeed, error) {
	switch s {
	case "minimum":
		return FAN_MINIMUM, nil
	case "medium":
		return FAN_MEDIUM, nil
	case "maximum":
		return FAN_MAXIMUM, nil
	case "auto":
		return FAN_AUTO, nil
	default:
		return FAN_NONE, errors.New("Invalid fan speed string")
	}
}
