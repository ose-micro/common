package common

import "time"

type State int32

type Status interface {
	State() State
	Previous() *State
	OccurOn() time.Time
}
