package banner

import (
	"Open_IM/internal/rpc"
	api "Open_IM/pkg/base_info"
	"context"
)

// func DB() *gorm.DB {
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
// 		config.Config.Mysql.DBUserName, config.Config.Mysql.DBPassword, config.Config.Mysql.DBAddress[0], "mysql")
// 	var db *gorm.DB
// 	db, err := gorm.Open(mysql.Open(dsn), nil)
// 	if err != nil {
// 		fmt.Println("Open failed ", err.Error(), dsn)
// 	}

// 	return db
// }

func SaveBanner(ctx context.Context, banner *api.Banner) (*api.Banner, error) {
	conn := rpc.DB()
	err := conn.Model(&api.Banner{}).Create(&banner).Error
	if err != nil {
		return nil, err
	}
	return banner, nil
}

func GetBannerByName(ctx context.Context, banner *api.Banner) (*api.Banner, error) {
	conn := rpc.DB()
	err := conn.Model(banner).Where("name = ?", banner.Name).First(&banner).Error
	if err != nil {
		return nil, err
	}
	return banner, nil
}

func ListBanners(ctx context.Context, params *api.FilterParams) (*[]api.Banner, error) {
	offset := (params.Page - 1) * params.PerPage
	if offset < 0 {
		offset = 0
	}

	conn := rpc.DB()
	banners := []api.Banner{}
	err := conn.Model(&api.Banner{}).Limit(int(params.PerPage)).Offset(int(offset)).Find(&banners).Error
	if err != nil {
		return nil, err
	}
	return &banners, nil
}

func DeleteBannerByName(ctx context.Context, banner *api.Banner) (*api.Banner, error) {
	conn := rpc.DB()
	err := conn.Model(banner).Where("name = ?", banner.Name).Delete(&banner).Error
	if err != nil {
		return nil, err
	}
	return banner, nil
}
