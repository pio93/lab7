package zlog

import (
	"fmt"
	"time"
)

//ViewersLog data structure is a more efficient storage of zap events.
type ViewersLog struct {
	//TODO(student) finish struct
	chansViewers map[string]*ChannelViewers
}

// NewViewersZapLogger returns a viewers logger.
func NewViewersZapLogger() *ViewersLog {
	//TODO(student) finish constructor
	chViewers := make(map[string]*ChannelViewers)
	return &ViewersLog{chansViewers: chViewers}
}

//TODO(student) implement ZapLogger interface for your more efficient data structure.

//Add is...
func (zs *ViewersLog) Add(z ChZap) {

	_, okTo := zs.chansViewers[z.ToChan]

	_, okFrom := zs.chansViewers[z.FromChan]

	if okTo == false {
		zs.chansViewers[z.ToChan] = &ChannelViewers{Channel: z.ToChan, Viewers: 0}
	}

	if okFrom == false {
		zs.chansViewers[z.ToChan] = &ChannelViewers{Channel: z.ToChan, Viewers: 0}
	}

	if okTo == true {
		zs.chansViewers[z.ToChan].Viewers++
	}

	if okFrom == true {
		zs.chansViewers[z.FromChan].Viewers--
	}
}

//Entries is...
func (zs *ViewersLog) Entries() int {
	return len(zs.chansViewers)
}

func (zs *ViewersLog) String() string {
	return fmt.Sprintf("SS: %d", len(zs.chansViewers))
}

//Viewers is..
func (zs *ViewersLog) Viewers(chName string) int {
	defer TimeElapsed(time.Now(), "simple.Viewers")
	return zs.chansViewers[chName].Viewers
}

//Channels is...
func (zs *ViewersLog) Channels() []string {
	defer TimeElapsed(time.Now(), "simple.Channels")
	//TODO(student) write this method
	result := make([]string, 0)
	for key := range zs.chansViewers {
		result = append(result, key)
	}
	return result
}

//ChannelsViewers is..
func (zs *ViewersLog) ChannelsViewers() []*ChannelViewers {
	defer TimeElapsed(time.Now(), "simple.ChannelsViewers")
	//TODO(student) write this method

	result := make([]*ChannelViewers, 0)
	for _, value := range zs.chansViewers {
		result = append(result, value)
	}

	return result
}
