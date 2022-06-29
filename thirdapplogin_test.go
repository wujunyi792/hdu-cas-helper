package hdu_cas_helper

import "testing"

func TestLogin(t *testing.T) {
	type args struct {
		cas *LoginStatus
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test login",
			args: args{cas: CasPasswordLogin("", "")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SklLogin(tt.args.cas); got.Error() != nil {
				t.Errorf("SKL Login() Get An Err: %v", got.Error())
			} else {
				t.Logf("SKL Successfully Get Skl Token: %s", got.GetToken())
			}
			if got := IHduLogin(tt.args.cas); got.Error() != nil {
				t.Errorf("Ihdu Login() Get An Err: %v", got.Error())
			} else {
				t.Logf("Ihdu Successfully Get IHdu Cookie: %s", got.GetCookie())
			}
			if got := NewJWLogin(tt.args.cas); got.Error() != nil {
				t.Errorf("newJW Login() Get An Err: %v", got.Error())
			} else {
				t.Logf("newJw Successfully Get IHdu Cookie: %s", got.GetCookie())
			}
		})
	}
}
