# HDU CAS HELPER 杭电cas登录辅助

> 使用本工具，可以快速登录到杭电cas，并快速登录其他所有相关服务，如上课啦、杭电助手

## cas 登陆方式
- [x] 账号密码
- [x] 二维码
- [ ] 短信验证码

## 受支持的第三方校内应用
- [x] 上课啦
- [x] 杭电助手
- [x] 新正方教务
- [x] 智慧杭电
- [x] 杭电勤工三助

## Quick start

```shell
go get github.com/wujunyi792/hdu-cas-helper
```

无论登录哪个服务，必须先登录 cas

### cas login
```
// 杭电 CAS 账号密码
ticker := hducashelper.CasPasswordLogin("", "") 
```
```
// 使用二维码登录， 控制台会输出链接， 使用微信打开连接即可登录
ticker := hducashelper.CasQrCodeLogin().PrintScannerUrl().AsyncLogin(5 * time.Second, 10) 
```

获取到 ticket 以后，可以使用这个ticket来登录其他应用

### 上课啦登录示例
```Go
package main

import (
	hducashelper "github.com/wujunyi792/hdu-cas-helper"
	"log"
)

func main() {
	ticker := hducashelper.CasPasswordLogin("", "") // 杭电 CAS 账号密码
	//ticker := hducashelper.CasQrCodeLogin().PrintScannerUrl().AsyncLogin(5 * time.Second, 10)
	sklLogin := hducashelper.SklLogin(ticker)
	if sklLogin.Error() != nil {
		log.Fatalln(sklLogin.Error())
	}
	log.Println(sklLogin.GetToken())
}
```

### 智慧杭电登录示例
```Go
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
```


### 新正方教务登录示例
```Go
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

```


### 杭电助手登录示例
```Go
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
	log.Println(hduHelpLogin.GetToken())
}

```

### 杭电勤工登录示例
```Go
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

```


Now you obtained the necessary credentials to use these third-party applications, enjoy it!


