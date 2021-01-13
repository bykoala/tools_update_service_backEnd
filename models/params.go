package models

type ParamUpdate struct {
	Version string `json:"version" binding:"required"`
	Content string `json:"content" binding:"required"`
	Forced  int8   `json:"forced" binding:"-"`
	Url     string `json:"url" binding:"required"`
	Status  int8   `json:"status" binding:"min=0,max=10"`
}

type ParamGetUpdateinfo struct {
	Version string `json:"version" binding:"required"`
}
