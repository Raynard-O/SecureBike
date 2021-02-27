package clients

import "time"

type Client struct {
	ID   uint64
	Name string
	Secure bool
	Duration time.Time
}
