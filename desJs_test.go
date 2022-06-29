package hdu_cas_helper

import "testing"

func Test_strEnc(t *testing.T) {
	type args struct {
		data      string
		firstKey  string
		secondKey string
		thirdKey  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "password",
			args: args{
				data:      "a" + "b" + "c",
				firstKey:  "1",
				secondKey: "2",
				thirdKey:  "3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(strEnc(tt.args.data, tt.args.firstKey, tt.args.secondKey, tt.args.thirdKey))
		})
	}
}

func Test_strDec(t *testing.T) {
	type args struct {
		data      string
		firstKey  string
		secondKey string
		thirdKey  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "dec test",
			args: args{
				data:      "39644174795FB4D0",
				firstKey:  "1",
				secondKey: "2",
				thirdKey:  "3",
			},
			want: "a" + "b" + "c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strDec(tt.args.data, tt.args.firstKey, tt.args.secondKey, tt.args.thirdKey); got != tt.want {
				t.Errorf("strDec() = %v, want %v", got, tt.want)
			}
		})
	}
}
