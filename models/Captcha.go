// Package models provides 
package models

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

var cpt *captcha.Captcha

func init(){
  store := cache.NewMemoryCache()
  cpt = captcha.NewWithFilter("/captcha",store)
  cpt.ChallengeNums = 4
  cpt.StdWidth = 100
  cpt.StdHeight = 40
}