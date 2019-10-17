package app

import "testing"

func TestStartTrack_Validate(t *testing.T) {
	type args struct {
		englishName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				englishName: "Nyota Uhura",
			},
			wantErr: false,
		},
		{
			name: "invalid character",
			args: args{
				englishName: "Nyota-Uhura",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &StartTrack{}
			if err := st.Validate(tt.args.englishName); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStartTrack_Translate(t *testing.T) {
	type args struct {
		englishName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path 1",
			args: args{
				englishName: "Uhura",
			},
			want: "0xF8E5 0xF8D6 0xF8E5 0xF8E1 0xF8D0",
		},
		{
			name: "happy path 2",
			args: args{
				englishName: "Leonard James Akaar",
			},
			want: "0xF8D9 0xF8D4 0xF8DD 0xF8DB 0xF8D0 0xF8E1 0xF8D3 0x0020 0xF8D8 0xF8D0 0xF8DA 0xF8D4 0xF8E2 0x0020 0xF8D0 0xF8DF 0xF8D0 0xF8D0 0xF8E1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &StartTrack{}
			if got := st.Translate(tt.args.englishName); got != tt.want {
				t.Errorf("Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
