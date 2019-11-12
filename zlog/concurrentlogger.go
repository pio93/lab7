package zlog

import (
	"fmt"
	"sync"
	"time"
)

// ConcurrentLog data structure is a completely accurate storage of zap events.
type ConcurrentLog struct {
	//TODO(student) finish struct
	chansViewers map[string]*channelLock
	Durations    []time.Duration
}

// NewConcurrentZapLogger returns a concurrent logger.
func NewConcurrentZapLogger() *ConcurrentLog {
	//TODO(student) finish constructor
	chLock := make(map[string]*channelLock)
	return &ConcurrentLog{chansViewers: chLock}

}

//TODO(student) implement ZapLogger interface for your more efficient data structure.

//Add checks if channel already exists. If not it adds it to map. It also increments or decrements number of viewers
//based on FromChan and ToChan
func (zs *ConcurrentLog) Add(z ChZap) {
	_, okTo := zs.chansViewers[z.ToChan]

	_, okFrom := zs.chansViewers[z.FromChan]

	if okTo == false {
		chanMutex := new(sync.Mutex)
		zs.chansViewers[z.ToChan] = &channelLock{chanViewers: &ChannelViewers{Channel: z.ToChan, Viewers: 1}, mutex: chanMutex}
	}

	if okFrom == false {
		chanMutex := new(sync.Mutex)
		zs.chansViewers[z.FromChan] = &channelLock{chanViewers: &ChannelViewers{Channel: z.FromChan, Viewers: 0}, mutex: chanMutex}
	}

	if okTo == true {
		zs.chansViewers[z.ToChan].mutex.Lock()
		zs.chansViewers[z.ToChan].chanViewers.Viewers++
		zs.chansViewers[z.ToChan].mutex.Unlock()
	}

	if okFrom == true {
		zs.chansViewers[z.FromChan].mutex.Lock()
		zs.chansViewers[z.FromChan].chanViewers.Viewers--
		if zs.chansViewers[z.FromChan].chanViewers.Viewers <= 0 {
			zs.chansViewers[z.FromChan].chanViewers.Viewers = 0
		}
		zs.chansViewers[z.FromChan].mutex.Unlock()
	}
}

//Entries returns length of a map
func (zs *ConcurrentLog) Entries() int {
	return len(zs.chansViewers)
}

func (zs *ConcurrentLog) String() string {
	return fmt.Sprintf("SS: %d", len(zs.chansViewers))
}

//Viewers returns number of viewers on specific channel
func (zs *ConcurrentLog) Viewers(chName string) int {
	defer func() {
		TimeElapsed(time.Now(), "simple.Viewers")
		zs.chansViewers[chName].mutex.Unlock()
	}()
	zs.chansViewers[chName].mutex.Lock()
	return zs.chansViewers[chName].chanViewers.Viewers
}

//Channels return a list with channels
func (zs *ConcurrentLog) Channels() []string {
	defer TimeElapsed(time.Now(), "simple.Channels")
	//TODO(student) write this method
	result := make([]string, 0)
	for key := range zs.chansViewers {
		result = append(result, key)
	}
	return result
}

//ChannelsViewers returns a list with pointers of ChannelViewers objects.
func (zs *ConcurrentLog) ChannelsViewers() []*ChannelViewers {
	//TODO(student) write this method
	defer TimeElapsed(time.Now(), "simple.ChannelsViewers")

	result := make([]*ChannelViewers, 0)
	for _, value := range zs.chansViewers {
		result = append(result, value.chanViewers)
	}

	return result
}
func (zs *ConcurrentLog) AverageDuration() time.Duration {
	var totDuration, n int
	for _, dur := range (*zs).Durations {
		totDuration += int(dur)
	}
	if n = (len((*zs).Durations)); n > 0 {
		return time.Duration((totDuration / n))
	}
	return 0
}
