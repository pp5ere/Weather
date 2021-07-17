package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStrToDate(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		input string
		want  time.Time
	}{
		{"Case 1", "2020-05-21", time.Date(2020, 05, 21, 0, 0, 0, 0, time.UTC)},
		{"Case 2", "2093-01-21", time.Date(2093, 01, 21, 0, 0, 0, 0, time.UTC)},
		{"Case 3", "1919-12-12", time.Date(1919, 12, 12, 0, 0, 0, 0, time.UTC)},
		{"Case 4", "2020-02-28", time.Date(2020, 02, 28, 0, 0, 0, 0, time.UTC)},
	}
	for _, test := range tests {
		got, err := StrToDate(test.input)
		if err != nil {
			t.Errorf("Error to run StrToDate() test case: %s, error: %v", err, test.name)
			return
		}
		if !assert.Equal(got, test.want) {
			t.Errorf("In test: %s, expected: %v, found: %v", test.name, test.want, got)
		}
	}
}

func TestCalculateHeatIndex(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		tempF  float64
		humRel float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Case 1", args{tempF: 62.49, humRel: 59.22}, 61.22},
		{"Case 2", args{tempF: 84.06, humRel: 51.44}, 85.52},
	}
	for _, test := range tests {
		got := CalculateHeatIndex(test.args.tempF, test.args.humRel)
		if !assert.Equal(got, test.want) {
			t.Errorf("CalculateHeatIndex() = %v, want: %v, fail: %s", got, test.want, test.name )
		}
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		tempF float64
		want  float64
	}{
		{"Case 1", 74.3, 23.5},
		{"Case 2", 0, -17.78},
		{"Case 3", 32, 0},
	}
	for _, test := range tests {
		got := FahrenheitToCelsius(test.tempF)
		if !assert.Equal(got, test.want) {
			t.Errorf("FahrenheitToCelsius() = %v, want %v, fail: %s", got, test.want, test.name)
		}
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		tempC float64
		want  float64
	}{
		{"Case 1", 0, 32},
		{"Case 2", -17.78, 0},
		{"Case 3", 23.5, 74.3},
	}
	for _, test := range tests {
		got := CelsiusToFahrenheit(test.tempC)
		if !assert.Equal(got, test.want) {
			t.Errorf("CelsiusToFahrenheit() = %v, want %v, fail: %s", got, test.want, test.name)
		}
	}
}

func TestDewPoint(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		tempC float64
		hum   float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Case 1", args{0, 70}, -4.88},
		{"Case 2", args{25.4, 82}, 22.1},
		{"Case 3", args{38.5, 90}, 36.58},
	}
	for _, test := range tests {
		got := DewPoint(test.args.tempC, test.args.hum)
		if !assert.Equal(got, test.want) {
			t.Errorf("DewPoint() = %v, want %v, fail: %s", got, test.want, test.name)
		}
	}
}

