package zlog

import (
	"fmt"
	"sort"
	"time"
)

// ViewersLog data structure is a more efficient storage of zap events.
type ViewersLog struct {
	//TODO(student) finish struct
	Zaps         chan ChZap
	viewers      int
	channelNames map[string]bool
	Chans        map[string]ChannelViewers
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
//using channelnames as keys with empty struct as value to save memory
// then converting to a string slice
func (zs *ViewersLog) Channels() []string {
	defer TimeElapsed(time.Now(), "simple.Channels")
	str := make([]string, 0)
	zap := (*zs).Chans
	if len(zap) < 1 {
		return nil
	}
	for _, v := range zap {
		str = append(str, v.Channel)
	}
	return str
}

//ChannelsViewers works similarly to ChannelViewers in simplelogger.go.
func (zs *ViewersLog) ChannelsViewers() []*ChannelViewers {
	defer TimeElapsed(time.Now(), "simple.ChannelsViewers")
	channels := zs.Channels()
	result := make([]*ChannelViewers, 0)
	if channels == nil || len(channels) == 0 {
		return nil
	}
	result = make([]*ChannelViewers, len(channels))
	for _, str := range channels {
		viewers := zs.Viewers(str)
		channelViewer := ChannelViewers{Channel: str, Viewers: viewers}
		result = append(result, &channelViewer)
	}

	return result
}
func (zs *ViewersLog) Top10() []*ChannelViewers {
	ch := zs.ChannelsViewers()
	sort.SliceStable(ch, func(i, j int) bool {
		return ch[i].Viewers > ch[j].Viewers
	})
	if len(ch) >= 10 {
		s := (ch)[:10]
		return s
	}
	return nil
}
