package entity

import "time"

type Todo struct {
	ID          int64
	Title       string
	Description string
	Duedate     time.Time
	Completed 	int
	CreatedAt time.Time
}
