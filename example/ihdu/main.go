package main

import (
	hducashelper "github.com/wujunyi792/hdu-cas-helper"
	"log"
)

func main() {
	ticker := hducashelper.CasPasswordLogin("", "") // 杭电 CAS 账号密码
	//ticker := hducashelper.CasQrCodeLogin().PrintScannerUrl().AsyncLogin(5 * time.Second, 10) // 使用二维码登录， 控制台会输出链接， 使用微信打开连接即可登录
	iHduLogin := hducashelper.IHduLogin(ticker)
	if iHduLogin.Error() != nil {
		log.Fatalln(iHduLogin.Error())
	}
	log.Println(iHduLogin.GetCookie())
}
