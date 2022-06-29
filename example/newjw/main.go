package main

import (
	hducashelper "github.com/wujunyi792/hdu-cas-helper"
	"log"
)

func main() {
	ticker := hducashelper.CasPasswordLogin("", "") // 杭电 CAS 账号密码
	//ticker := hducashelper.CasQrCodeLogin().PrintScannerUrl().AsyncLogin(5 * time.Second, 10) // 使用二维码登录， 控制台会输出链接， 使用微信打开连接即可登录
	newJwLogin := hducashelper.NewJWLogin(ticker)
	if newJwLogin.Error() != nil {
		log.Fatalln(newJwLogin.Error())
	}
	log.Println(newJwLogin.GetCookie())
}
