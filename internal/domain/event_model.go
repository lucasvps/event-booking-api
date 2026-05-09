package domain

import (
	"time"
)

type Event struct {
	ID                          int64
	Name, Description, Location string    `binding:"required"`
	DateTime                    time.Time `binding:"required"`
	UserID                      int64
}
