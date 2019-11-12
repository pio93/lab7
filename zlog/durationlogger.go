package zlog

import (
	"fmt"
	"time"
)

//DurationLog is a struct that stores map with IP as key and with pointers to channelDurations as values
// and slice of strings that represent statistics
type DurationLog struct {
	chansDurations map[string]*ChannelDurations
	Stats          []string
	Durations      []time.Duration
}

//NewDurationLogger creates new DurationLog object
func NewDurationLogger() *DurationLog {
	chDur := make(map[string]*ChannelDurations)
	stats := make([]string, 0)
	return &DurationLog{chansDurations: chDur, Stats: stats}
}

//Add adds new channelDuration if it is not found in map. If zap's ip is already a key in the map, add function puts current zap as
//previous zap in channelDurations and calcualte duration using chZap Duration() function. It also assigns a stat string to stats slice.
func (zs *DurationLog) Add(z ChZap) {
	_, ok := zs.chansDurations[z.IP]

	if ok == false {
		zeroSec, _ := time.ParseDuration("0s")
		zs.chansDurations[z.IP] = &ChannelDurations{Zap: z, Duration: zeroSec}
	}

	if ok == true {
		if z.Time != zs.chansDurations[z.IP].Zap.Time {
			dur := zs.chansDurations[z.IP].Zap.Duration(z)
			zs.chansDurations[z.IP].Zap = z
			zs.chansDurations[z.IP].Duration = dur
			stat := fmt.Sprintf("Client %s watched %s for %v", zs.chansDurations[z.IP].Zap.IP, zs.chansDurations[z.IP].Zap.FromChan, zs.chansDurations[z.IP].Duration)
			zs.Stats = append(zs.Stats, stat)
		}
	}
}

//GetStats returns the list of statistics
func (zs *DurationLog) GetStats() []string {
	return zs.Stats
}

//Length returns the length of the map
func (zs *DurationLog) Length() int {
	return len(zs.chansDurations)
}

//ClearStats is clears the list of statistics
func (zs *DurationLog) ClearStats() {
	zs.Stats = nil
}
func (zs *DurationLog) AverageDuration() time.Duration {
	var totDuration, n int
	for _, dur := range (*zs).Durations {
		totDuration += int(dur)
	}
	if n = (len((*zs).Durations)); n > 0 {
		return time.Duration((totDuration / n))
	}
	return 0
}
