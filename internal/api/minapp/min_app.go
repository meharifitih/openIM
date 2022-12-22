package minapp

import (
	"Open_IM/internal/rpc/miniapp"
	api "Open_IM/pkg/base_info"
	"Open_IM/pkg/common/log"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

var (
	dst string = "./images"
)

const DefaultPageSize = 10

func AddMiniApp(c *gin.Context) {
	params := api.MiniAppDto{}
	if err := c.Bind(&params); err != nil {
		log.NewError("BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	icons, err := json.Marshal(params.Icons)
	if err != nil {
		c.JSON(http.StatusBadRequest, "unable to marshall icons")
		return
	}

	pages, err := json.Marshal(params.Pages)
	if err != nil {
		c.JSON(http.StatusBadRequest, "unable to marshall pages")
		return
	}

	widgets, err := json.Marshal(params.Widget)
	if err != nil {
		c.JSON(http.StatusBadRequest, "unable to marshall widgets")
		return
	}

	permission, err := json.Marshal(params.ReqPermission)
	if err != nil {
		c.JSON(http.StatusBadRequest, "unable to marshall permissions")
		return
	}

	app := api.MiniApp{
		Dir:             params.Dir,
		Lang:            params.Lang,
		AppName:         params.AppName,
		ShortName:       params.ShortName,
		Description:     params.Description,
		VerName:         params.VerName,
		VerCode:         params.VerCode,
		MiniPlatformVer: params.MiniPlatformVer,
		Icons:           datatypes.JSON(icons),
		Pages:           datatypes.JSON(pages),
		Widget:          datatypes.JSON(widgets),
		ReqPermission:   datatypes.JSON(permission),
	}

	miniApp, err := miniapp.SaveMiniApp(c, &app)
	if err != nil {
		log.NewError("SaveMiniApp failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	log.NewInfo("SaveMiniApp api return ", miniApp)
	c.JSON(http.StatusOK, miniApp)
}

func UpdateMiniApp(c *gin.Context) {
	params := api.UpdateMiniAPp{}
	if err := c.Bind(&params); err != nil {
		log.NewError("BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	name := c.Param("name")
	app := api.MiniApp{
		AppName:         name,
		VerCode:         params.VerCode,
		VerName:         params.VerName,
		MiniPlatformVer: params.MiniPlatformVer,
	}

	miniApp, err := miniapp.UpdateMiniApp(c, &app)
	if err != nil {
		log.NewError("UpdateMiniApp failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	log.NewInfo("UpdateMiniApp api return ", miniApp)
	c.JSON(http.StatusOK, miniApp)
}

func GetMiniApp(c *gin.Context) {
	name := c.Param("name")
	params := api.MiniApp{
		AppName: name,
	}

	miniApp, err := miniapp.GetMiniApp(c, &params)
	if err != nil {
		log.NewError("GetMiniApp failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	log.NewInfo("GetMiniApp api return ", miniApp)
	c.JSON(http.StatusOK, miniApp)
}

func ListMiniApp(c *gin.Context) {
	filterParam := api.QueryParams{}
	filter := api.FilterParams{}
	err := c.BindQuery(&filterParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	if filterParam.Page == "" {
		filter.Page = 1
	} else {

		p, err := toInt64(filterParam.Page)
		if err != nil {
			p = 1
		}

		filter.Page = p
	}

	if filterParam.PerPage == "" {
		filter.PerPage = DefaultPageSize
	} else {

		pp, err := toInt64(filterParam.PerPage)
		if err != nil {
			pp = DefaultPageSize
		}
		filter.PerPage = pp
	}

	miniApp, err := miniapp.ListMiniApp(c, &filter)
	if err != nil {
		log.NewError("ListMiniApp failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	log.NewInfo("ListMiniApp api return ", miniApp)
	c.JSON(http.StatusOK, miniApp)
}

func toInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func DeleteMiniApp(c *gin.Context) {
	params := api.MiniApp{}
	name := c.Param("name")
	params.AppName = name

	_, err := miniapp.DeleteMiniApp(c, &params)
	if err != nil {
		log.NewError("DeleteMiniApp failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	log.NewInfo("DeleteMiniApp api return ", "banner deleted")
	c.JSON(http.StatusOK, "app deleted")
}
