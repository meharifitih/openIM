package base_info

import (
	"time"

	"gorm.io/datatypes"
)

type MiniApp struct {
	Id              uint           `json:"id" gorm:"primaryKey"`
	Dir             string         `json:"dir,omitempty"`
	Lang            string         `json:"lang,omitempty"`
	AppName         string         `json:"app_name,omitempty" gorm:"unique;not null"`
	ShortName       string         `json:"short_name,omitempty" gorm:"unique"`
	Description     string         `json:"description,omitempty"`
	Icons           datatypes.JSON `json:"icons,omitempty"`
	VerName         string         `json:"ver_name,omitempty"`
	VerCode         int            `json:"ver_code,omitempty"`
	MiniPlatformVer string         `json:"mini_ver,omitempty"`
	Pages           datatypes.JSON `json:"pages,omitempty"`
	Widget          datatypes.JSON `json:"widget,omitempty"`
	ReqPermission   datatypes.JSON `json:"req_permission,omitempty"`
	CreatedAt       time.Time      `json:"created_at,omitempty"`
	UpdatedAt       time.Time      `json:"updated_at,omitempty"`
	DeletedAt       *time.Time     `json:"deleted_at,omitempty" gorm:"index"`
}

type MiniAppDto struct {
	Id              uint     `json:"id" gorm:"primaryKey"`
	Dir             string   `json:"dir,omitempty"`
	Lang            string   `json:"lang,omitempty"`
	AppName         string   `json:"app_name,omitempty" binding:"required"`
	ShortName       string   `json:"short_name,omitempty"`
	Description     string   `json:"description,omitempty"`
	Icons           []string `json:"icons,omitempty" binding:"required"`
	VerName         string   `json:"ver_name,omitempty" binding:"required"`
	VerCode         int      `json:"ver_code,omitempty" binding:"required"`
	MiniPlatformVer string   `json:"mini_ver,omitempty" binding:"required"`
	Pages           []string `json:"pages,omitempty" binding:"required"`
	Widget          []string `json:"widget,omitempty" binding:"required"`
	ReqPermission   []string `json:"req_permission,omitempty"`
}

type UpdateMiniAPp struct {
	VerName         string `json:"ver_name,omitempty" binding:"required"`
	VerCode         int    `json:"ver_code,omitempty" binding:"required"`
	MiniPlatformVer string `json:"mini_ver,omitempty" binding:"required"`
}
