package models

import "time"

type GetUpdateInfo struct {
	Version    string    `db:"version"`
	Content    string    `db:"content"`
	Url        string    `db:"url"`
	Size       float32   `db:"size"`
	Status     int8      `db:"status"`
	CreateTime time.Time `db:"create_time"`
}

type UpdateInfo struct {
	Version        string    `db:"version"`
	Content        string    `db:"content"`
	Url            string    `db:"url"`
	Md5            string    `db:"md5"`
	Status         int8      `db:"status"`
	Forced         int8      `db:"forced"`
	Classification int8      `db:"classification"`
	Size           float32   `db:"size"`
	CreateTime     time.Time `db:"create_time"`
}
