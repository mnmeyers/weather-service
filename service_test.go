package main

import (
	"context"
	"reflect"
	"testing"
)

func TestGetService(t *testing.T) {
	tests := []struct {
		name string
		want Service
	}{
		{"test", &ServiceImpl{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceImpl_GetWeather(t *testing.T) {
	type args struct {
		ctx context.Context
		lat float64
		lon float64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "invalid coordinates",
			args: args{
				ctx: context.Background(),
				lat: 91,  // invalid latitude
				lon: 181, // invalid longitude
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServiceImpl{}
			got, err := s.GetWeather(tt.args.ctx, tt.args.lat, tt.args.lon)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWeather() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetWeather() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWeatherCharacterization(t *testing.T) {
	type args struct {
		temp int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hot temperature",
			args: args{
				temp: 95,
			},
			want: "hot",
		},
		{
			name: "cold temperature",
			args: args{
				temp: 35,
			},
			want: "cold",
		},
		{
			name: "moderate temperature",
			args: args{
				temp: 60,
			},
			want: "moderate",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWeatherCharacterization(tt.args.temp); got != tt.want {
				t.Errorf("getWeatherCharacterization() = %v, want %v", got, tt.want)
			}
		})
	}
}
