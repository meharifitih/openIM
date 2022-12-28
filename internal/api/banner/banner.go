package banner

import (
	"Open_IM/internal/rpc/banner"
	api "Open_IM/pkg/base_info"
	"Open_IM/pkg/common/log"
	"fmt"
	"io/ioutil"
	lg "log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// var (
// 	dst string = "./images"
// )

const (
	DefaultPageSize        = 10
	dst             string = "./images"
)

func AddBanner(c *gin.Context) {

	params := api.Banner{}
	if err := c.Bind(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	file := params.File

	split := strings.Split(file.Filename, ".")
	fileName := params.Name + "." + split[1]
	err := c.SaveUploadedFile(file, dst+`/`+fileName)

	if err != nil {
		c.Error(err)
		return
	}

	imageUrl := fmt.Sprintf("/manage/images/%s", fileName)
	params.ImgUrl = imageUrl

	banner, err := banner.SaveBanner(c, &params)
	if err != nil {
		log.NewError(params.Name, "AddBanner failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	log.NewInfo(params.Name, "AddBanner api return ", banner)
	c.JSON(http.StatusOK, banner)

}

func GetBanner(c *gin.Context) {

	params := api.Banner{}
	name := c.Param("name")
	params.Name = name

	banner, err := banner.GetBannerByName(c, &params)
	if err != nil {
		log.NewError(params.Name, "GetBannerByName failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	log.NewInfo(params.Name, "GetBannerByName api return ", banner)
	c.JSON(http.StatusOK, banner)

}

func toInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func ListBanners(c *gin.Context) {
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

	banners, err := banner.ListBanners(c, &filter)
	if err != nil {
		log.NewError("ListBanners failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	log.NewInfo("ListBanners api return ", banners)
	c.JSON(http.StatusOK, banners)
}

func DeleteBanner(c *gin.Context) {

	params := api.Banner{}
	name := c.Param("name")
	params.Name = name

	_, err := banner.DeleteBannerByName(c, &params)
	if err != nil {
		log.NewError(params.Name, "GetBannerByName failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	log.NewInfo(params.Name, "GetBannerByName api return ", "banner deleted")
	c.JSON(http.StatusOK, "banner deleted")

}

func GetImage(c *gin.Context) {
	fileName := c.Param("id")
	x, err := exists(fileName)
	lg.Println("the images status is ", x)

	if err != nil {
		fmt.Println(err)
		c.Error(err)
		return
	}

	c.File(dst + "/" + fileName)
}

func exists(fileName string) (bool, error) {

	files, err := ioutil.ReadDir(dst)
	if err != nil {
		// log.Error(err.Error())
		lg.Fatal(err)
	}

	for _, file := range files {

		if strings.Contains(file.Name(), fileName) {
			return true, nil
		}
	}

	_, err = os.Stat(dst + fileName)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
