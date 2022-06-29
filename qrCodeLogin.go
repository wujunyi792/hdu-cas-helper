package hdu_cas_helper

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	uuid "github.com/satori/go.uuid"
	"time"
)

type QrLogin struct {
	login     *LoginStatus
	confirmed bool
	uuid      string
	url       string
}

func (q *QrLogin) GetScannerUrl() string {
	return q.url
}

func (q *QrLogin) PrintScannerUrl() *QrLogin {
	fmt.Println(q.url)
	return q
}

func (q *QrLogin) GetUUID() string {
	return q.uuid
}

func (q *QrLogin) Login() *LoginStatus {
	type casRes struct {
		RedirectUrl string `json:"redirect_url"`
	}
	var res casRes
	resp, _, err := gorequest.New().
		Get("https://cas.hdu.edu.cn/cas/checkQRCodeScan?uuid=" + q.uuid).
		EndStruct(&res)
	if err != nil {
		q.login.err = err[0]
		return q.login
	}
	casTgc := getCookie(resp, "CASTGC")
	if len(casTgc) == 0 {
		q.login.err = errors.New("cas login failed, maybe hdu service is done, please check manually")
		return q.login
	}
	q.login.tgc = casTgc
	return q.login
}

func (q *QrLogin) AsyncLogin(interval time.Duration, retry int) *LoginStatus {
	for {
		q.login.err = nil
		s := q.Login()
		retry--
		if s.Error() == nil || retry <= 0 {
			return s
		}
		time.Sleep(interval)
	}
}

func CasQrLoginWithUUID(uuid string) *LoginStatus {
	qr := &QrLogin{
		login: &LoginStatus{
			method:  METHODQRCODE,
			tgc:     "",
			expired: false,
			service: "",
			err:     nil,
		},
		confirmed: false,
		uuid:      uuid,
	}
	return qr.Login()
}

func CasAsyncQrLoginWithUUID(uuid string, interval time.Duration, retry int) *LoginStatus {
	qr := &QrLogin{
		login: &LoginStatus{
			method:  METHODQRCODE,
			tgc:     "",
			expired: false,
			service: "",
			err:     nil,
		},
		confirmed: false,
		uuid:      uuid,
	}
	return qr.AsyncLogin(interval, retry)
}

func CasQrCodeLogin() *QrLogin {
	uid := uuid.NewV4().String()
	return &QrLogin{
		login: &LoginStatus{
			method:  METHODQRCODE,
			tgc:     "",
			expired: false,
			service: "",
			err:     nil,
		},
		confirmed: false,
		uuid:      uid,
		url:       "https://cas.hdu.edu.cn/cas/qyQrLogin?uuid=" + uid,
	}
}
