package zlog

import (
	"fmt"
	"sync"
	"time"
)

// ConcurrentLog data structure is a completely accurate storage of zap events.
type ConcurrentLog struct {
	//TODO(student) finish struct
	chansViewers map[string]*ChannelViewers
	mutex        sync.Mutex
}

// NewConcurrentZapLogger returns a concurrent logger.
func NewConcurrentZapLogger() *ConcurrentLog {
	//TODO(student) finish constructor
	chViewers := make(map[string]*ChannelViewers)
	return &ConcurrentLog{chansViewers: chViewers}

}

//TODO(student) implement ZapLogger interface for your more efficient data structure.

//Add checks if channel already exists. If not it adds it to map. It also increments or decrements number of viewers
//based on FromChan and ToChan
func (zs *ConcurrentLog) Add(z ChZap) {
	defer zs.mutex.Unlock()
	zs.mutex.Lock()
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
	}
}

//Entries returns length of a map
func (zs *ConcurrentLog) Entries() int {
	defer zs.mutex.Unlock()
	zs.mutex.Lock()
	return len(zs.chansViewers)
}

func (zs *ConcurrentLog) String() string {
	defer zs.mutex.Lock()
	zs.mutex.Lock()
	return fmt.Sprintf("SS: %d", len(zs.chansViewers))
}

//Viewers returns number of viewers on specific channel
func (zs *ConcurrentLog) Viewers(chName string) int {
	defer func() {
		TimeElapsed(time.Now(), "simple.Viewers")
		zs.mutex.Unlock()
	}()
	zs.mutex.Lock()
	return zs.chansViewers[chName].Viewers
}

//Channels return a list with channels
func (zs *ConcurrentLog) Channels() []string {
	defer func() {
		TimeElapsed(time.Now(), "simple.Channels")
		zs.mutex.Unlock()
	}()
	zs.mutex.Lock()
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
		result = append(result, value)
	}

	return result
}
