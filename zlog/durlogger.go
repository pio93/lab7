package zlog

import (
	"time"
)

// DurLogger is an interface for storing durations of zap events
type DurLogger interface {
	Add(ChZap)
	GetDurations() []time.Duration
	Length() int
	AverageDuration() time.Duration
	Max() time.Duration
	Min() time.Duration
	ClearDurations()
}
