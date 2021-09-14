// Package frontend provides 
package frontend

import (
	"fmt"
	"time"
	"xinmall/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Get()  {
	
	c.BaseInit()

	startTime := time.Now().UnixNano()

	//Banner
	banner := []models.Banner{}
	if hasBanner := models.CacheDB.Get("banner", &banner); hasBanner == true {
		c.Data["bannerList"] = banner
	} else {
		models.DB.Where("status=1 AND banner_type=1").Order("sort desc").Find(&banner)
		c.Data["bannerList"] = banner
		models.CacheDB.Set("banner", banner)
	}

	//Phone list
	phone := models.Product{}
	if hasPhone := modes.CacheDB.Get("phone",&phone); hasPhone == true {
		c.Data["phoneList"] = phone
	} else {
		phone := models.GetProductByCategory(1,"hot",8)
		c.Data["phoneList"] = phone
		models.CacheDB.Set("phone",phone)
	}

	//Television list
	Tv := []models.Product{}
	if hasTv := models.CacheDB.Get("tv",&Tv); hasTv == true {
		c.Data["tvList"] == Tv
	} else {
		tv := models.GetProductByCategory(4, "best", 8)
		c.Data["tvList"] = tv
		models.CacheDb.Set("tv", tv)
	}

	//endtime
	endTime := time.Now().UnixNano()

	fmt.Println("执行时间",endTime-startTime)
	c.TplName = "frontend/index/index.html"
}	