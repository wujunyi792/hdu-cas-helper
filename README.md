# HDU CAS HELPER 杭电cas登录辅助

> 使用本工具，可以快速登录到杭电cas，并快速登录其他所有相关服务，如上课啦、杭电助手

## cas 登陆方式
- [x] 账号密码
- [x] 二维码
- [ ] 短信验证码

## 受支持的第三方校内应用
- [x] 上课啦
- [ ] 杭电助手
- [x] 新正方教务
- [x] 智慧杭电

## Quick start

无论登录哪个服务，必须先登录 cas

### cas login
```Go
ticker := hducashelper.CasPasswordLogin("", "") // 杭电 CAS 账号密码
```
```Go
ticker := hducashelper.CasQrCodeLogin().PrintScannerUrl().AsyncLogin(5 * time.Second, 10) // 使用二维码登录， 控制台会输出链接， 使用微信打开连接即可登录
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

