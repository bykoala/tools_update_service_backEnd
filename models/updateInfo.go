package models

import "time"

type UpdateInfo struct {
	Version    string    `db:"version"`
	Content    string    `db:"content"`
	Forced     int8      `db:"forced"`
	Url        string    `db:"url"`
	Status     int8      `db:"status"`
	CreateTiem time.Time `db:"create_time"`
}
