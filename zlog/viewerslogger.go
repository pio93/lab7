package zlog

import (
	"fmt"
	"time"
)

// ViewersLog data structure is a more efficient storage of zap events.
type ViewersLog struct {
	//TODO(student) finish struct
	Zaps         chan ChZap
	viewers      int
	channelNames map[string]bool
}

// NewViewersZapLogger returns a viewers logger.
func NewViewersZapLogger() *ViewersLog {
	//TODO(student) finish constructor
	channel := make(chan ChZap)
	return &ViewersLog{Zaps: channel, viewers: 0}
}

//TODO(student) implement ZapLogger interface for your more efficient data structure.

//Add puts zap on Zaps channel and strings with names to channelNames.
func (zs *ViewersLog) Add(z ChZap) {
	zs.Zaps <- z
	zs.channelNames[z.ToChan] = true
	zs.channelNames[z.FromChan] = true
}

//Entries return length of channel.
func (zs *ViewersLog) Entries() int {
	return len(zs.Zaps)
}

func (zs *ViewersLog) String() string {
	return fmt.Sprintf("SS: %d", len(zs.Zaps))
}

//Viewers reads from zapChannel and and updates viewers counter.
func (zs *ViewersLog) Viewers(chName string) int {
	defer TimeElapsed(time.Now(), "simple.Viewers")

	v := <-zs.Zaps

	if v.ToChan == chName {
		zs.viewers++
	}
	if v.FromChan == chName {
		zs.viewers--
		if zs.viewers <= 0 {
			zs.viewers = 0
		}
	}
	return zs.viewers

}

//Channels return list of all channels without duplicates.
func (zs *ViewersLog) Channels() []string {
	defer TimeElapsed(time.Now(), "simple.Channels")

	result := make([]string, 0)

	for key := range zs.channelNames {
		result = append(result, key)
	}

	return result
}

//ChannelsViewers works similarly to ChannelViewers in simplelogger.go.
func (zs *ViewersLog) ChannelsViewers() []*ChannelViewers {
	defer TimeElapsed(time.Now(), "simple.ChannelsViewers")
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
