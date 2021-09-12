package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gomarkdown/markdown"
	"github.com/hunterhug/go_image"
	_ "github.com/jinzhu/gorm"
	"io/ioutil"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//time to date
func TimestampToDate(timestamp int) string {
	t := time.Unix(int64(timestamp),0)
	return t.Format("2021-09-12 08:39:01")
}

//get current time
func GetUnix() int64 {
	curtime := time.Now().Unix()

	fmt.Println(curtime)
	return curtime
}

//get nano time
func GetUnixNano() int64 {
	nanotime := time.Now().UnixNano()
	return nanotime
}

//get date time
func GetDate() string {
	template := "2021-09-12 08:39:01"
	return time.Now().Format(template)
}

//get formatday
func FormatDay() string {
	template := "20210912"
	return time.Now().Format(template)
}


//Md5 encrypt
func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return string(hex.EncodeToString(m.Sum(nil)))
}

//verify email
func VerifyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` 
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//generate order
func GenerateOrderId() string {
	template := "200601021504"
	return time.Now().Format(template)
}

//send message
func SendMsg(str string)  {
	ioutil.WriteFile("test_send.text",[]byte(str),12345)	
}

//picture
func ResizeImage(filename string)  {
	extName := path.Ext(filename)
	resizeImage := strings.Split(beego.AppConfig.String("resizeImageSize"), ",")

	for i := 0; i < len(resizeImage); i++ {
		w := resizeImage[i]
		width, _ := strconv.Atoi(w)
		savepath := filename + "_" + w + "x" + w + extName
		err := go_image.ThumbnailF2F(filename, savepath, width, width)
		if err != nil {
			beego.Error(err)
		}
	}
}

//format image
func FormatImage(picName string) string {
	ossStatus, err := beego.AppConfig.Bool("ossStatus")
	if err != nil {
		//判断目录前面是否有/
		flag := strings.Contains(picName, "/static")
		if flag {
			return picName
		}
		return "/" + picName
	}
	if ossStatus {
		return beego.AppConfig.String("ossDomain") + "/" + picName
	} else {
		flag := strings.Contains(picName, "/static")
		if flag {
			return picName
		}
		return "/" + picName

	}
}

//format title
func FormatAttribute(str string) string {
	md := []byte(str)
	htmlByte := markdown.ToHTML(md, nil, nil)
	return string(htmlByte)
}

//mul
func Mul(price float64, num int) float64 {
	return price * float64(num)
}
