package chzap

import (
	"fmt"
	"net"
	"strings"
	"time"
)

const timeLayout = "2006/01/02, 15:04:05"

type ChZap struct {
	Date   time.Time
	IP     net.IP
	FromCh string
	ToCh   string
}

func NewChZap(chzap string) *ChZap {
	chzapSlice := strings.Split(chzap, ", ")
	if len(chzapSlice) == 5 {
		dateOnly := chzapSlice[0]
		timeOnly := chzapSlice[1]
		dateTimeSlice := []string{dateOnly, timeOnly}
		dateTime := strings.Join(dateTimeSlice, ", ")
		date, err := time.Parse(timeLayout, dateTime)

		ip := net.ParseIP(chzapSlice[2])
		fromCh := chzapSlice[3]
		toCh := chzapSlice[4]

		if err != nil {
			fmt.Println(err)
		}

		return &ChZap{date, ip, fromCh, toCh}
	} else {
		return new(ChZap)
	}
}

func (ch *ChZap) String() string {
	s := []string{
		ch.Date.Format(timeLayout),
		ch.IP.String(),
		ch.FromCh,
		ch.ToCh,
	}
	return strings.Join(s, ", ")
}

func (ch *ChZap) Duration(provided ChZap) time.Duration {
	duration := ch.Date.Sub(provided.Date)
	return duration
}
