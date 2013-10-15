package chzap

import (
	"net"
	"time"
	"strings"
	"fmt"
)

const timeLayout = "2006/01/02, 15:04:05"

type ChZap struct {
	Date time.Time
	IP net.IP
	FromCh string
	ToCh string
}

func NewChZap(chzap string) *ChZap {
	chzapArr := strings.Split(chzap, ", ")
	dateOnly := chzapArr[0]
	timeOnly := chzapArr[1]
	dateTimeSlice := []string{dateOnly, timeOnly}
	dateTime := strings.Join(dateTimeSlice, ", ")
	date, err := time.Parse(timeLayout, dateTime)

	ip := net.ParseIP(chzapArr[2])
	fromCh := chzapArr[3]
	toCh := chzapArr[4]

	if err != nil {
		fmt.Println(err)
	}

	return &ChZap{date, ip, fromCh, toCh}
}

func(ch *ChZap) String() string {
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
