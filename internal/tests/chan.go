package tests

import (
	"time"
	time1 "time"
)

//go:generate go run ../../fcgen/fcgen.go -type=Chan,ChanUnExport,ChanEmbed

type Chan struct {
	Chan0 chan int
	Chan1 chan<- int
	Chan2 <-chan int

	Chan3 chan *int
	Chan4 chan<- *any
	Chan5 <-chan *any

	Chan6 chan time.Time
	Chan7 chan *time1.Time

	Chan8          **chan **int
	Chan9          **<-chan **int
	Chan10         **chan<- **int
	Chan11, Chan12 **<-chan **[5][]func(a, b **any) (c, d **string)
}

//nolint:unused //need to test unexported fields
type ChanUnExport struct {
	_        chan int
	unExport chan int

	Chan0 chan *msg
	Chan1 chan msg
}

//nolint:unused //need to test unexported fields
type ChanEmbed struct {
	chanMsg
	ChanMsg

	Chan0 chanMsg
	Chan1 ChanMsg
}

type chanMsg chan msg
type ChanMsg chan msg

type msg struct{}
type Msg struct{}
