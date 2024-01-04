package helpers

import (
	"testing"
)

func TestAuthenticate(t *testing.T) {
	type args struct {
		encryptedPassword string
		reqPassword       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test True Hash",
			args: args{
				encryptedPassword: "$2a$10$SBiQdDRczTLR14KZ88HUleyDnF.tlFaSziAxg9nJ35IcSUjSVJNfW",
				reqPassword:       "1234",
			},
			want: true,
		},
		{
			name: "Test False Hash",
			args: args{
				encryptedPassword: "$2a$10$SBiQdDRczTLR14KZ88HUleyDnF.tlFaSziAxg9nJ35IcSUjSVJNfW",
				reqPassword:       "12345",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Authenticate(tt.args.encryptedPassword, tt.args.reqPassword); got != tt.want {
				t.Errorf("Authenticate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasingPassword(t *testing.T) {
	type args struct {
		reqPassword string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Hashing",
			args: args{
				reqPassword: "1234",
			},
			want: "$2a$10$OHSpPvj6lnaQmzEnCIWlFO.etDdlNgGEIjwRJk.kOO9RcrYJQ7dam",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasingPassword(tt.args.reqPassword); got != tt.want {
				t.Errorf("HasingPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
