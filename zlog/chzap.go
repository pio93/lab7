// +build !solution

// Leave an empty line above this comment.

package zlog

import (
	"errors"
	"fmt"
	"strings"
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
	Time   time.Time
	IP     string
	Status string
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

	if len(event) <= timeLen {
		err := fmt.Errorf("NewSTBEvent: too short event string: %s", event)
		return nil, nil, err
	}

	vals := strings.Split(event, ",")

	time, err := time.Parse(timeFormat, event[0:timeLen])

	if err != nil {
		err := errors.New("NewSTBEvent: failed to parse timestamp")
		return nil, nil, err
	}

	if len(vals) < 4 {
		err := fmt.Errorf("NewSTBEvent: event with too few fields: %s", event)
		return nil, nil, err
	}

	status := strings.Split(vals[3], ":")

	if len(status) == 1 {

		if len(vals) < 5 {
			err := fmt.Errorf("NewSTBEvent: event with too few fields: %s", event)
			return nil, nil, err
		}

		chZap := ChZap{
			Time:     time,
			IP:       strings.TrimSpace(vals[2]),
			FromChan: strings.TrimSpace(vals[3]),
			ToChan:   strings.TrimSpace(vals[4]),
		}
		return &chZap, nil, nil

	}

	statChange := StatusChange{
		Time:   time,
		IP:     strings.TrimSpace(vals[2]),
		Status: (strings.TrimSpace(vals[3])),
	}
	return nil, &statChange, nil

}

func (zap ChZap) String() string {
	//TODO(student) write this method
	return fmt.Sprintf("%v, %s, %s, %s", zap.Time, zap.IP, zap.FromChan, zap.ToChan)
}

func (schg StatusChange) String() string {
	//TODO(student) write this method
	return fmt.Sprintf("%v, %s, %s", schg.Time, schg.IP, schg.Status)
}

// Duration between receiving (this) zap event and the provided event.
func (zap ChZap) Duration(provided ChZap) time.Duration {
	//TODO(student) write this method
	return provided.Time.Sub(zap.Time)
}

// Date returns the date of the zap event.
func (zap ChZap) Date() string {
	//TODO(student) write this method
	vals := strings.Split(zap.Time.String(), ",")
	return strings.TrimSpace(vals[0])
}
