package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "TestAddSuccess",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "TestAddSuccess2",
			args: args{
				a: 4,
				b: 5,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divide(t *testing.T) {
	type args struct {
		dividend float64
		divisor  float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestDivideSuccess",
			args: args{
				dividend: 1,
				divisor: 2,
			},
			want: 0.5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide(tt.args.dividend, tt.args.divisor)
			if (err != nil) != tt.wantErr {
				t.Errorf("divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Area(t *testing.T) {
	type fields struct {
		Radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
		{
			name: "TestCircleSuccess",
			fields: fields{Radius: 5.0},
			want: 78.53981633974483,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.Radius,
			}
			if got := c.Area(); got != tt.want {
				t.Errorf("Circle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}


func BenchmarkGenTestAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(1,2)
	}
}
func BenchmarkGenTestDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		divide(1,2)
	}
}
func BenchmarkGenTestCircle(b *testing.B) {
	c := Circle{
		Radius: 5,
	}
	for i := 0; i < b.N; i++ {
		c.Area()
	}
}