package model

import "time"

type Executors struct {
	Id        int
	Name      string
	Type      string
	RenewTime time.Time
}
