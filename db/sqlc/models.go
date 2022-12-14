// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"
)

type Task struct {
	ID        int64     `json:"id"`
	User      int64     `json:"user"`
	Title     string    `json:"title"`
	Priority  int32     `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
