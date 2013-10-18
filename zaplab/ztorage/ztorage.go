package ztorage

import (
	"github.com/sandves/zaplab/chzap"
)

//type Zaps []chzap.ChZap
type Zaps map[string]int

func NewZapStore() *Zaps {
	//zs := make(Zaps, 0)
	zs := make(Zaps)
	return &zs
}

/*func (zs *Zaps) StoreZap(z chzap.ChZap) {
	//*zs = append(*zs, z)
}*/

/*func (zs *Zaps) ComputeViewers(chName string) int {
	viewers := 0
	for _, v := range *zs {
		if v.ToChan == chName {
			viewers++
		}
		if v.FromChan == chName {
			viewers--
		}
	}
	return viewers
}*/

func (zs Zaps) StoreZap(z chzap.ChZap) {
	_, ok := zs[z.ToChan]
	if !ok {
		zs[z.ToChan] = 0
	}
	_, ok = zs[z.FromChan]
	if !ok {
		zs[z.FromChan] = 0
	}

	for key, _ := range zs {
		if z.ToChan  == key {
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
