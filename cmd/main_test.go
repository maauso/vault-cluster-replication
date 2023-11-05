package main

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestDurationToSeconds(t *testing.T) {
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "10 second",
			args: args{
				duration: 10 * time.Second,
			},
			want: 10 * time.Second,
		},
		{
			name: "10 minutes",
			args: args{
				duration: 10 * time.Minute,
			},
			want: 10 * time.Minute,
		},
		{
			name: "10 hours",
			args: args{
				duration: 10 * time.Hour,
			},
			want: 10 * time.Hour,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DurationToSeconds(tt.args.duration); got != tt.want {
				t.Errorf("DurationToSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    Config
		wantErr bool
	}{
		{
			name: "Default config",
			want: Config{
				ConfigFilePath:             "configs/default-conf.yaml",
				ScheduledExecutionInterval: 5 * time.Minute,
			},
		},
		{
			name: "Custom config",
			want: Config{
				ConfigFilePath:             "configs/custom-conf.yaml",
				ScheduledExecutionInterval: 10 * time.Minute,
			},
		},
		{
			name:    "Custom config with invalid interval",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Custom config" {
				err := os.Setenv("CONFIG_FILE_PATH", "configs/custom-conf.yaml")
				if err != nil {
					t.Errorf("parseConfig() = %v", err)
				}
				err = os.Setenv("SCHEDULED_EXECUTION_INTERVAL", "10m")
				if err != nil {
					t.Errorf("parseConfig() = %v", err)
				}
			}
			if tt.name == "Custom config with invalid interval" {
				err := os.Setenv("SCHEDULED_EXECUTION_INTERVAL", "10y")
				if err != nil {
					t.Errorf("parseConfig() = %v", err)
				}
			}

			if got, _ := parseConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
