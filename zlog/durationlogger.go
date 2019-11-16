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
	MaxValue       time.Duration
	MinValue       time.Duration
	AllDurations   time.Duration
	Entries        int
}

//NewDurationLogger creates new DurationLog object
func NewDurationLogger() *DurationLog {
	chDur := make(map[string]*ChZap)
	durations := make([]time.Duration, 0)
	t, _ := time.ParseDuration("24h")
	return &DurationLog{
		chansDurations: chDur,
		Durations:      durations,
		MaxValue:       0,
		MinValue:       t,
		AllDurations:   0,
		Entries:        0,
	}
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
	fmt.Printf("Entires before: %d\n", zs.Entries)
	zs.Entries += len(zs.Durations)
	fmt.Printf("Entires after: %d\n", zs.Entries)
	fmt.Printf("Durations before: %d\n", zs.AllDurations)
	for _, dur := range zs.Durations {
		zs.AllDurations += dur
	}
	fmt.Printf("Durations after: %d\n", zs.AllDurations)
	if zs.Entries > 0 {
		return time.Duration((int64(zs.AllDurations) / int64(zs.Entries))) / 1000000000
	}
	return 0
}

//Max is...
func (zs *DurationLog) Max() time.Duration {
	sort.Slice(zs.Durations, func(i, j int) bool {
		return zs.Durations[i] > zs.Durations[j]
	})

	if len(zs.Durations) > 0 {
		if zs.MaxValue < zs.Durations[0]/1000000000 {
			zs.MaxValue = zs.Durations[0] / 1000000000
		}
	}

	return zs.MaxValue
}

//Min is...
func (zs *DurationLog) Min() time.Duration {
	sort.Slice(zs.Durations, func(i, j int) bool {
		return zs.Durations[i] < zs.Durations[j]
	})

	if len(zs.Durations) > 0 {
		if zs.MinValue > zs.Durations[0]/1000000000 {
			zs.MinValue = zs.Durations[0] / 1000000000
		}
	}
	if len(zs.Durations) <= 0 {
		return 0
	}
	return zs.MinValue
}

//ClearDurations is...
func (zs *DurationLog) ClearDurations() {
	fmt.Println(zs.Durations)
	zs.Durations = nil
	fmt.Println(zs.Durations)
}
