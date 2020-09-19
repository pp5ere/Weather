package util

import (
	"reflect"
	"testing"
	"time"
)

func TestStrToDate(t *testing.T) {
	type args struct {
		dateStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "TestStrToDate 1",
			args: args{
				dateStr: "2020-05-21",
			},
			want: time.Date(2020, 05, 21, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "TestStrToDate 2",
			args: args{
				dateStr: "2093-01-21",
			},
			want: time.Date(2093, 01, 21, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "TestStrToDate 3",
			args: args{
				dateStr: "1919-12-12",
			},
			want: time.Date(1919, 12, 12, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StrToDate(tt.args.dateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrToDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrToDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateHeatIndex(t *testing.T) {
	type args struct {
		tempF  float64
		humRel float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Case 1",
			args: args{
				tempF: 62.49,
				humRel: 59.22,
			},
			want: 61.22,//CelsiusToFahrenheit(16.22),
		},
		{
			name: "Case 2",
			args: args{
				tempF: 61.47,
				humRel: 59.54,
			},
			want: CelsiusToFahrenheit(26.88),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateHeatIndex(tt.args.tempF, tt.args.humRel); got != tt.want {
				t.Errorf("CalculateHeatIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
