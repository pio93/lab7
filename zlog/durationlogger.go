package zlog

import (
	"fmt"
	"time"
)

//DurationLog is...
type DurationLog struct {
	chansDurations map[string]*ChannelDurations
	Stats          []string
}

//NewDurationLogger is..
func NewDurationLogger() *DurationLog {
	chDur := make(map[string]*ChannelDurations)
	stats := make([]string, 0)
	return &DurationLog{chansDurations: chDur, Stats: stats}
}

//Add is...
func (zs *DurationLog) Add(z ChZap) {
	_, ok := zs.chansDurations[z.IP]

	if ok == false {
		zeroSec, _ := time.ParseDuration("0s")
		zs.chansDurations[z.IP] = &ChannelDurations{Zap: z, Duration: zeroSec}
	}

	if ok == true {
		dur := zs.chansDurations[z.IP].Zap.Duration(z)
		zs.chansDurations[z.IP].Zap = z
		zs.chansDurations[z.IP].Duration = dur
		stat := fmt.Sprintf("Client %s watched %s for %v", zs.chansDurations[z.IP].Zap.IP, zs.chansDurations[z.IP].Zap.FromChan, zs.chansDurations[z.IP].Duration)
		zs.Stats = append(zs.Stats, stat)
	}
}

//GetStats is...
func (zs *DurationLog) GetStats() []string {
	return zs.Stats
}

//Length is..
func (zs *DurationLog) Length() int {
	return len(zs.chansDurations)
}

//ClearStats is ...
func (zs *DurationLog) ClearStats() {
	zs.Stats = nil
	fmt.Println("Cleared stats")
	fmt.Println(len(zs.Stats))
}
