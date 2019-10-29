// +build !solution

// Leave an empty line above this comment.

package zlog

import (
	"time"
)

const (
	timeFormat = "2006/01/02, 15:04:05"
	dateFormat = "2006/01/02"
	timeOnly   = "15:04:05"
	timeLen    = len(timeFormat)
)

// StatusChange holds information about a status change event.
type StatusChange struct {
	Time       time.Time
	IP         string
	Volume     int
	MuteStatus bool
	HDMIStatus bool
	//TODO(student) finish struct
}

// ChZap holds information about a channel change event, aka zap event.
type ChZap struct {
	Time     time.Time
	IP       string
	FromChan string
	ToChan   string
	//TODO(student) finish struct
}

// NewSTBEvent returns zap event or a status change event.
// If the input string does not match the expected format, an error is returned.
func NewSTBEvent(event string) (*ChZap, *StatusChange, error) {
	//TODO(student) write this method
	return nil, nil, nil
}

func (zap ChZap) String() string {
	//TODO(student) write this method
	return ""
}

func (schg StatusChange) String() string {
	//TODO(student) write this method
	return ""
}

// Duration between receiving (this) zap event and the provided event.
func (zap ChZap) Duration(provided ChZap) time.Duration {
	//TODO(student) write this method
	return time.Duration(0)
}

// Date returns the date of the zap event.
func (zap ChZap) Date() string {
	//TODO(student) write this method
	return ""
}
