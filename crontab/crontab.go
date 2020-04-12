package crontab

import (
	"errors"
	"math/rand"
	"time"
)

type Crontab struct {
	Name       string
	IntervalMS int
	Handler    func()
}

func NewCrontab(name string, intervalMS int, handler func()) (*Crontab, error) {
	if name == "" {
		name = "default"
	}
	if intervalMS <= 0 {
		intervalMS = 2000
	}
	if handler == nil {
		return nil, errors.New("handler should not be nil")
	}

	return &Crontab{
		Name:       name,
		IntervalMS: intervalMS,
		Handler:    handler,
	}, nil
}

func (c Crontab) Start() {
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(10)) * 100 * time.Millisecond)
	ch := time.Tick(5 * time.Second)
	for {
		c.Handler()
		<-ch
	}
}
