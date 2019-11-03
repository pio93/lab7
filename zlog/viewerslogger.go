package zlog
import "time"

// ViewersLog data structure is a more efficient storage of zap events.
type ViewersLog struct {
	//TODO(student) finish struct
	zaps map[string]ChZap
	Channels map[string]ChannelViewers
	Durations[] time.Duration
}

// NewViewersZapLogger returns a viewers logger.
func NewViewersZapLogger() *ViewersLog {
	//TODO(student) finish constructor
	return &ViewersLog{Zaps: make(map[string]ChZap), Chans: make(map[string]ChannelViewers),Durations:make([]time.Duration,0)}
}

//TODO(student) implement ZapLogger interface for your more efficient data structure.
type ZapLogger interface{
	Add(ChZap)
	Entries() int
	Viewers(channelName string) int
	Channels() []string
	ChannelsViewers() []*ChannelViewers
}
func (zs *ViewersLog) Entries() int {
	return len((*zs).Zaps)
}
func (zs *ViewersLog) String() string {
	return fmt.Sprintf("SS: %d", len((*zs).Zaps))
}

func (zs *ViewersLog) Viewers(chName string) int {
	if v, ok := ((*zs).Channels)[chName]; ok {
		return v.Viewers
	} else {
		return 0
	}
}

func (zs *ViewersLog) Channels() []string {
	defer util.TimeElapsed(time.Now(), "simple.Channels")
	str := make([]string, 0)
	//copy map values to a slice
	zap := (*zs).Chans
	for _, v := range zap {
		str = append(str, v.Channel)
	}
	return str
}

// conversion from default map to slice
// this is done in order to sort the data later
func (zs *ViewersLog) ChannelsViewers() ChannelViewersList {
	defer util.TimeElapsed(time.Now(), "simple.ChannelsViewers")
	zap := (*zs).Chans
	chslice := make(ChannelViewersList, 0, len(zap))
	for _, v := range zap {
		chslice = append(chslice, v)
	}
	return chslice
}
