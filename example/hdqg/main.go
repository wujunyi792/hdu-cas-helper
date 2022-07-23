package main

import (
	hducashelper "github.com/wujunyi792/hdu-cas-helper"
	"log"
)

func main() {
	ticker := hducashelper.CasPasswordLogin("", "") // 杭电 CAS 账号密码
	hduHelpLogin := hducashelper.HduHelpLogin(ticker)
	if hduHelpLogin.Error() != nil {
		log.Fatalln(hduHelpLogin.Error())
	}
	hdqgLogin := hducashelper.HDQGLogin(hduHelpLogin)
	log.Println(hdqgLogin.GetToken())
}
