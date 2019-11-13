package zlog

import (
	"fmt"
	"sort"
	"time"
)

//DurationLog is ey and with pointers to channelDurations as values
// and slice of strings that represent statistics
type DurationLog struct {
	chansDurations map[string]*ChZap
	Durations      []time.Duration
}

//NewDurationLogger creates new DurationLog object
func NewDurationLogger() *DurationLog {
	chDur := make(map[string]*ChZap)
	durations := make([]time.Duration, 0)
	return &DurationLog{chansDurations: chDur, Durations: durations}
}

//Add adds new channelDuration if it is not found in map. If zap's ip is already a key in the map, add function puts current zap as
//previous zap in channelDurations and calcualte duration using chZap Duration() function. It also assigns a stat string to stats slice.
func (zs *DurationLog) Add(z ChZap) {
	_, ok := zs.chansDurations[z.IP]

	if ok == false {
		zs.chansDurations[z.IP] = &z
	}

	if ok == true {
		if z.Time != zs.chansDurations[z.IP].Time {
			dur := zs.chansDurations[z.IP].Duration(z)
			zs.chansDurations[z.IP] = &z
			//durs := fmt.Sprintf("Client %s watched %s for %v", zs.chansDurations[z.IP].Zap.IP, zs.chansDurations[z.IP].Zap.FromChan, zs.chansDurations[z.IP].Duration)
			zs.Durations = append(zs.Durations, dur)
		}
	}
}

//GetDurations returns the list of statistics
func (zs *DurationLog) GetDurations() []time.Duration {
	return zs.Durations
}

//Length returns the length of the map
func (zs *DurationLog) Length() int {
	return len(zs.chansDurations)
}

//AverageDuration is..
func (zs *DurationLog) AverageDuration() time.Duration {
	var totDuration int64
	n := len(zs.Durations)
	for _, dur := range zs.Durations {
		totDuration += int64(dur)
	}
	if n > 0 {
		return time.Duration((totDuration / int64(n))) / 1000000000
	}
	return 0
}

//Max is...
func (zs *DurationLog) Max() time.Duration {
	sort.SliceStable(zs.Durations, func(i, j int) bool {
		return zs.Durations[i] > zs.Durations[j]
	})

	if len(zs.Durations) > 0 {
		return zs.Durations[0] / 1000000000
	}

	return 0
}

//Min is...
func (zs *DurationLog) Min() time.Duration {
	sort.SliceStable(zs.Durations, func(i, j int) bool {
		return zs.Durations[i] < zs.Durations[j]
	})

	if len(zs.Durations) > 0 {
		return zs.Durations[0] / 1000000000
	}

	return 0

}

//ClearDurations is...
func (zs *DurationLog) ClearDurations() {
	fmt.Println(zs.Durations)
	zs.Durations = nil
	fmt.Println(zs.Durations)
}
