package hdu_cas_helper

import (
	"testing"
	"time"
)

func TestQrCodeLogin(t *testing.T) {
	c := CasQrCodeLogin()
	t.Log(c.GetScannerUrl())
	t.Logf(c.AsyncLogin(5*time.Second, 10).tgc)
}
