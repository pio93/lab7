package zlog

import "time"

// DurLogger is an interface for storing durations of zap events
type DurLogger interface {
	Add(ChZap)
	GetStats() []string
	Length() int
	ClearStats()
}

//ChannelDurations stores previous channel zap and duration between previous and curren zap for the same IP
type ChannelDurations struct {
	Zap      ChZap
	Duration time.Duration
}
