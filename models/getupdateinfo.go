package models

import "time"

type GetUpdateInfo struct {
	Status     int8      `db:"status"`
	Version    string    `db:"version"`
	Content    string    `db:"content"`
	Url        string    `db:"url"`
	CreateTime time.Time `db:"create_time"`
}
