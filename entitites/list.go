package entitites

import "time"

type List struct {
	Id        uint
	Task      string
	Deadline  time.Time 
	Completed bool
}