package base_info

import (
	open_im_sdk "Open_IM/pkg/proto/sdk_ws"
	"mime/multipart"
	"time"
)

type GetUsersInfoReq struct {
	OperationID string   `json:"operationID" binding:"required"`
	UserIDList  []string `json:"userIDList" binding:"required"`
}
type GetUsersInfoResp struct {
	CommResp
	UserInfoList []*open_im_sdk.PublicUserInfo `json:"-"`
	Data         []map[string]interface{}      `json:"data" swaggerignore:"true"`
}

type UpdateSelfUserInfoReq struct {
	ApiUserInfo
	OperationID string `json:"operationID" binding:"required"`
}
type SetGlobalRecvMessageOptReq struct {
	OperationID      string `json:"operationID" binding:"required"`
	GlobalRecvMsgOpt *int32 `json:"globalRecvMsgOpt" binding:"omitempty,oneof=0 1 2"`
}
type SetGlobalRecvMessageOptResp struct {
	CommResp
}
type UpdateUserInfoResp struct {
	CommResp
}

type GetSelfUserInfoReq struct {
	OperationID string `json:"operationID" binding:"required"`
	UserID      string `json:"userID" binding:"required"`
}
type GetSelfUserInfoResp struct {
	CommResp
	UserInfo *open_im_sdk.UserInfo  `json:"-"`
	Data     map[string]interface{} `json:"data" swaggerignore:"true"`
}

type GetFriendIDListFromCacheReq struct {
	OperationID string `json:"operationID" binding:"required"`
}

type GetFriendIDListFromCacheResp struct {
	CommResp
	UserIDList []string `json:"userIDList" binding:"required"`
}

type GetBlackIDListFromCacheReq struct {
	OperationID string `json:"operationID" binding:"required"`
}

type GetBlackIDListFromCacheResp struct {
	CommResp
	UserIDList []string `json:"userIDList" binding:"required"`
}

type UserStory struct {
	Id        uint                  `json:"id" form:"id" gorm:"primaryKey"`
	Name      string                `json:"name,omitempty" form:"name"`
	UserId    string                `json:"user_id,omitempty" form:"user_id" binding:"required"`
	File      *multipart.FileHeader `json:"file,omitempty" form:"file" binding:"required" gorm:"-"`
	Story     string                `json:"story,omitempty" form:"story"`
	CreatedAt time.Time             `json:"created_at,omitempty"`
	UpdatedAt time.Time             `json:"updated_at,omitempty"`
	DeletedAt *time.Time            `json:"deleted_at,omitempty" gorm:"index"`
}
