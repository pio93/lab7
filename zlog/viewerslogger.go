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

//Add checks if channel already exists. If not it adds it to map. It also increments or decrements number of viewers
//based on FromChan and ToChan
func (zs *ViewersLog) Add(z ChZap) {

	_, okTo := zs.chansViewers[z.ToChan]

	_, okFrom := zs.chansViewers[z.FromChan]

	if okTo == false {
		zs.chansViewers[z.ToChan] = &ChannelViewers{Channel: z.ToChan, Viewers: 1}

	}

	if okFrom == false {
		zs.chansViewers[z.FromChan] = &ChannelViewers{Channel: z.FromChan, Viewers: 0}

	}

	if okTo == true {
		zs.chansViewers[z.ToChan].Viewers++
	}

	if okFrom == true {
		zs.chansViewers[z.FromChan].Viewers--

		if zs.chansViewers[z.FromChan].Viewers <= 0 {
			zs.chansViewers[z.FromChan].Viewers = 0
		}

	}

}

//Entries returns length of a map
func (zs *ViewersLog) Entries() int {
	return len(zs.chansViewers)
}

func (zs *ViewersLog) String() string {
	return fmt.Sprintf("SS: %d", len(zs.chansViewers))
}

//Viewers returns number of viewers on specific channel
func (zs *ViewersLog) Viewers(chName string) int {
	defer TimeElapsed(time.Now(), "simple.Viewers")
	return zs.chansViewers[chName].Viewers
}

//Channels return a list with channels
func (zs *ViewersLog) Channels() []string {
	defer TimeElapsed(time.Now(), "simple.Channels")
	//TODO(student) write this method
	result := make([]string, 0)
	for key := range zs.chansViewers {
		result = append(result, key)
	}
	return result
}

//ChannelsViewers returns a list with pointers of ChannelViewers objects.
func (zs *ViewersLog) ChannelsViewers() []*ChannelViewers {
	defer TimeElapsed(time.Now(), "simple.ChannelsViewers")
	//TODO(student) write this method

	result := make([]*ChannelViewers, 0)
	for _, value := range zs.chansViewers {
		result = append(result, value)
	}

	return result
}
