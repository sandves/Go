package ztorage

import (
	"github.com/sandves/zaplab/chzap"
	"time"
)

type Zaps struct {
	Zaps map[string]int
	TotalNumberOfZaps int 
	TotalZapDuration time.Duration
	PrevZap chzap.ChZap
}

func NewZapStore() *Zaps {
	return &Zaps{make(map[string]int, 0, 0, nil}
}

func (zs Zaps) StoreZap(z chzap.ChZap) {

	zs.TotalNumberOfZaps++
	if zs.PrevZap != nil {
		dur := z.Duration(zs.PrevZap)
		zs.TotalZapDuration += dur
	}
	zs.PrevZap = z

	/*If the channel doesn't exist in the map,
	put the key (channelname) in the map and
	assign its value (number of viewers) to zero*/
	if _, ok := zs.Zaps[z.ToChan]; !ok {
		zs.Zaps[z.ToChan] = 0
	}
	if _, ok := zs.Zaps[z.FromChan]; !ok {
		zs.Zaps[z.FromChan] = 0
	}

	/*if a viewer zaps to a channel, increment the
	number of viewers by one.
	if a viewer leaves a channel, decrement the number
	of viewers by one*/
	for key, _ := range zs.Zaps {
		if z.ToChan == key {
			zs.Zaps[key]++
		}
		if z.FromChan == key {
			zs.Zaps[key]--
		}
	}
}

func (zs Zaps) AverageZapDuration() time.Duration {
	return (zs.TotalZapDuration)/(time.Duration(zs.TotalNumberOfZaps))
}

func (zs Zaps) TopTenChannels() []string {
	return Top10(zs.Zaps)
}
