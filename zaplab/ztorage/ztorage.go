package ztorage

import (
	"github.com/sandves/zaplab/chzap"
)

type Zaps map[string]int

func NewZapStore() *Zaps {
	zs := make(Zaps)
	return &zs
}

func (zs Zaps) StoreZap(z chzap.ChZap) {

	/*If the channel doesn't exist in the map,
	put the key (channelname) in the map and
	assign its value (number of viewers) to zero*/
	if _, ok := zs[z.ToChan]; !ok {
		zs[z.ToChan] = 0
	}
	if _, ok := zs[z.FromChan]; !ok {
		zs[z.FromChan] = 0
	}

	/*if a viewer zaps to a channel, increment the
	number of viewers by one.
	if a viewer leaves a channel, decrement the number*/
	//of viewers by one.
	for key, _ := range zs {
		if z.ToChan == key {
			zs[key]++
		}
		if z.FromChan == key {
			zs[key]--
		}
	}
}

func (zs Zaps) TopTenChannels() []string {
	return Top10(zs)
}
