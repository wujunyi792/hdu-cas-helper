package hdu_cas_helper

import (
	"testing"
)

func TestLogin(t *testing.T) {
	type args struct {
		cas *LoginStatus
	}
	tests := []struct {
		name string
		args args
		want *Skl
	}{
		{
			name: "test skl login",
			args: args{cas: CasPasswordLogin("", "")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Login(tt.args.cas); got.Error() != nil {
				t.Errorf("Login() Get An Err: %v", got.Error())
			} else {
				t.Logf("Successfully Get Skl Token: %s", got.GetToken())
			}
		})
	}
}
