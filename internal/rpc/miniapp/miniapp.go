package miniapp

import (
	"Open_IM/internal/rpc"
	api "Open_IM/pkg/base_info"
	"context"
)

func SaveMiniApp(ctx context.Context, minapp *api.MiniApp) (*api.MiniApp, error) {
	conn := rpc.DB()
	err := conn.Model(&api.MiniApp{}).Create(&minapp).Error
	if err != nil {
		return nil, err
	}

	return minapp, nil
}

func UpdateMiniApp(ctx context.Context, minapp *api.MiniApp) (*api.MiniApp, error) {
	conn := rpc.DB()
	err := conn.Model(&minapp).Where("app_name=? ", minapp.AppName).Updates(&minapp).Error
	if err != nil {
		return nil, err
	}

	return minapp, nil
}

func GetMiniApp(ctx context.Context, minapp *api.MiniApp) (*api.MiniApp, error) {
	conn := rpc.DB()
	err := conn.Model(&minapp).Where("app_name=? ", minapp.AppName).First(&minapp).Error
	if err != nil {
		return nil, err
	}

	return minapp, nil
}

func ListMiniApp(ctx context.Context, params *api.FilterParams) (*[]api.MiniApp, error) {
	offset := (params.Page - 1) * params.PerPage
	if offset < 0 {
		offset = 0
	}

	conn := rpc.DB()
	miniApps := []api.MiniApp{}
	err := conn.Model(&api.MiniApp{}).Limit(int(params.PerPage)).Offset(int(offset)).Find(&miniApps).Error
	if err != nil {
		return nil, err
	}
	return &miniApps, nil
}

func DeleteMiniApp(ctx context.Context, minapp *api.MiniApp) (*api.MiniApp, error) {
	conn := rpc.DB()
	err := conn.Model(minapp).Where("app_name = ?", minapp.AppName).Delete(&minapp).Error
	if err != nil {
		return nil, err
	}
	return minapp, nil
}
