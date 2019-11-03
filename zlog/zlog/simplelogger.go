// +build !solution

// Leave an empty line above this comment.

package zlog

import (
	"fmt"
	"time"
)

// Zaps is a slice of ChZap events.
type Zaps []ChZap

var viewers int

// NewSimpleZapLogger returns a simple zap logger that operates over a slice data structure.
func NewSimpleZapLogger() *Zaps {
	z := make(Zaps, 0)
	return &z
}

// Add adds a zap event to the simple logger's data structure.
func (zs *Zaps) Add(z ChZap) {
	*zs = append(*zs, z)
}

// Entries returns the number of entries in the log.
func (zs *Zaps) Entries() int {
	return len(*zs)
}

func (zs *Zaps) String() string {
	return fmt.Sprintf("SS: %d", len(*zs))
}

// Viewers returns the current number of viewers for a channel.
func (zs *Zaps) Viewers(chName string) int {
	defer TimeElapsed(time.Now(), "simple.Viewers")
	fmt.Println(zs.String())
	for _, v := range *zs {
		if v.ToChan == chName {
			viewers++
		}
		if v.FromChan == chName {
			viewers--
			if viewers <= 0 {
				viewers = 0
			}
		}
	}

	return viewers
}

// Channels returns a slice of the channels found in the zaps (both to and from).
func (zs *Zaps) Channels() []string {
	defer TimeElapsed(time.Now(), "simple.Channels")
	//TODO(student) write this method
	channelNames := make(map[string]bool)
	for _, v := range *zs {
		channelNames[v.ToChan] = true
		channelNames[v.FromChan] = true
	}

	result := make([]string, 0)

	for key := range channelNames {
		result = append(result, key)
	}
	return result
}

// ChannelsViewers returns a slice of ChannelViewers, which is defined in zaplogger.go.
// This is the number of viewers for each channel.
func (zs *Zaps) ChannelsViewers() []*ChannelViewers {
	defer TimeElapsed(time.Now(), "simple.ChannelsViewers")
	//TODO(student) write this method
	channels := zs.Channels()
	result := make([]*ChannelViewers, 0)

	for _, v := range channels {
		viewers := zs.Viewers(v)
		channelViewer := ChannelViewers{
			Channel: v,
			Viewers: viewers,
		}
		result = append(result, &channelViewer)
	}

	return result
}
