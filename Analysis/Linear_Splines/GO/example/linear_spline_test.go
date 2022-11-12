package example

import (
	"testing"
)

func Test_getFuelPerHourLinear(t *testing.T) {

	type args struct {
		arg float64
	}

	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "Arg error",
			args:    args{arg: float64(-10)},
			want:    float64(0),
			wantErr: true,
		},
		{name: "Arg error 2",
			args:    args{arg: float64(100)},
			want:    float64(0),
			wantErr: true,
		},
		{name: "Arg error 2",
			args:    args{arg: float64(22)},
			want:    float64(310),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getValLinear(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("getValLinear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("getValLinear() got = %v, want %v", got, tt.want)
			}
		})
	}

	for i := 0; i <= 22; i++ {
		val, _ := getValLinear(float64(i))
		t.Logf("%d, %f", i, val)
	}
}

func TestPlotFuelPerHourLinear(t *testing.T) {
	type args struct {
		path  string
		start float64
		end   float64
		step  float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "step1",
			args: args{
				path:  "step1.png",
				start: 0.0,
				end:   22.0,
				step:  1.0,
			},
			wantErr: false,
		},
		{
			name: "step0-1",
			args: args{
				path:  "step0-1.png",
				start: 0.0,
				end:   22.0,
				step:  0.1,
			},
			wantErr: false,
		},
		{
			name: "step0-5",
			args: args{
				path:  "step0-5.png",
				start: 0.0,
				end:   22.0,
				step:  0.5,
			},
			wantErr: false,
		},
		{
			name: "step2",
			args: args{
				path:  "step2.png",
				start: 0.0,
				end:   22.0,
				step:  2.0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PlotValLinear(tt.args.path, tt.args.start, tt.args.end, tt.args.step); (err != nil) != tt.wantErr {
				t.Errorf("PlotValLinear() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
