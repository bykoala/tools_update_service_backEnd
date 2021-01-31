package models

import "mime/multipart"

type ParamUpdate struct {
	Version string  `json:"version" binding:"required,min=6"`
	Content string  `json:"content" binding:"required"`
	Url     string  `json:"url" binding:"required" "startswith=http"`
	MD5     string  `json:"md5" binding:"required"`
	Size    float32 `json:"size" binding:"required"`
	Forced  int8    `json:"forced" binding:"min=0,max=1"`
	Status  int8    `json:"status" binding:"min=0,max=10"`
}

type ParamGetUpdateinfo struct {
	Version string `json:"version" binding:"required"`
}

type ParamFeedbackOrBug struct {
	Title          string `json:"title" binding:"required"`
	Content        string `json:"content" binding:"required"`
	Contact_way    string `json:"contact_way" binding:"required"`
	To_contact     string `json:"to_contact" binding:"required"`
	Classification int8   `json:"classification" binding:"required",min=0,max=1`
}

type ParamUpload struct {
	Version        string                `form:"version" binding:"required,min=6"`
	Classification string                `form:"classification" binding:"required"`
	Url            string                `form:"url" binding:"-"`
	Md5            string                `form:"md5" binding:"-"`
	Size           int64                 `form:"size" binding:"-"`
	FileName       *multipart.FileHeader `form:"fileName" binding:"required"`
}
