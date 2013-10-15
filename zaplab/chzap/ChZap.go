package chzap

import (
	"net"
	"time"
	"strings"
)

struct ChZap {
	Date Time
	IP IP
	FromCh string
	ToCh string
}

func(ch *ChZap) String() string {
	s := []string{
		(*ch).Date.String(),
		(*ch).IP.String(),
		(*ch).FromCH,
		(*ch).ToCh
	}
	return strings.join(s, ", ")
}

func (ch *ChZap) Duration(provided ChZap) time.Duration {
	duration := (*ch).Date.Sub(provided.Date)
	return duration
}
