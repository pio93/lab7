package zlog

import (
	"fmt"
)

// ZapLogger is the interface that must be implemented for different zap loggers.
type ZapLogger interface {
	Add(ChZap)
	Entries() int
	Viewers(channelName string) int
	Channels() []string
	ChannelsViewers() []*ChannelViewers
}

// ChannelViewers is a channel-viewers pair.
type ChannelViewers struct {
	Channel string
	Viewers int
}

// String returns a string representation for a channel-viewers pair.
func (cv ChannelViewers) String() string {
	return fmt.Sprintf("%s: %d", cv.Channel, cv.Viewers)
}
