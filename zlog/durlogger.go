package zlog

import "time"

// DurLogger is...
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
