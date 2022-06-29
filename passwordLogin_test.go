package hdu_cas_helper

import "testing"

func TestPasswordLogin(t *testing.T) {
	type args struct {
		uname    string
		password string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "pwdLoginTest",
			args: args{
				uname:    "",
				password: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CasPasswordLogin(tt.args.uname, tt.args.password)
		})
	}
}

func Test_formReq_makeForm(t *testing.T) {
	type fields struct {
		rsa       string
		ul        int
		pl        int
		lt        string
		execution string
		_eventId  string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test form",
			fields: fields{
				rsa:       "1211",
				ul:        13,
				pl:        13131,
				lt:        "3131",
				execution: "qwq",
				_eventId:  "qwfqwq",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &formReq{
				Rsa:       tt.fields.rsa,
				Ul:        tt.fields.ul,
				Pl:        tt.fields.pl,
				Lt:        tt.fields.lt,
				Execution: tt.fields.execution,
				EventId:   tt.fields._eventId,
			}
			c.makeForm()
		})
	}
}
