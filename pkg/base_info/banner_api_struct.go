package base_info

import (
	"mime/multipart"
	"time"
)

// type Banner struct {
// 	Id         uint       `json:"id" gorm:"primaryKey"`
// 	BannerName string     `json:"banner_name" gorm:"unique"`
// 	LinkUrl    string     `json:"link_url"`
// 	Internal   bool       `json:"internal"`
// 	ImgUrl     string     `json:"img_url"`
// 	CreatedAt  time.Time  `json:"created_at,omitempty"`
// 	UpdatedAt  time.Time  `json:"updated_at,omitempty"`
// 	DeletedAt  *time.Time `json:"deleted_at,omitempty" gorm:"index"`
// }

type Banner struct {
	Id        uint                  `json:"id" form:"id" gorm:"primaryKey"`
	Name      string                `json:"name" binding:"required" form:"name" gorm:"unique"`
	LinkUrl   string                `json:"link_url" binding:"required" form:"link_url"`
	Internal  bool                  `json:"internal" form:"internal"`
	File      *multipart.FileHeader `json:"file,omitempty" form:"file" gorm:"-"`
	ImgUrl    string                `json:"img_url,omitempty" form:"img_url"`
	CreatedAt time.Time             `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt time.Time             `json:"updated_at,omitempty" form:"updated_at"`
	DeletedAt *time.Time            `json:"deleted_at,omitempty" form:"deleted_at" gorm:"index"`
}

type QueryParams struct {
	Sort         string `json:"sort" form:"sort"`
	Filter       string `json:"filter" form:"filter"`
	Page         string `json:"page" form:"page"`
	PerPage      string `json:"per_page" form:"per_page"`
	LinkOperator string `json:"linkOperator" form:"linkOperator"`
}

type FilterParams struct {
	Page         int64  `json:"page"`
	PerPage      int64  `json:"per_page"`
	LinkOperator string `json:"linkOperator"`
	Total        int64  `json:"total"`
	NoSort       bool   `json:"-"`
	NoLimit      bool   `json:"-"`
}
