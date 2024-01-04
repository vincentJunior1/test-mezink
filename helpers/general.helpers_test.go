package helpers

import (
	"reflect"
	"testing"
	"time"
)

func TestDiffTime(t *testing.T) {
	type args struct {
		a time.Time
		b time.Time
	}
	a := time.Now().UTC()
	b := a.Add(time.Hour * 1)
	c := a.Add(time.Hour * 25)
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Hour",
			args: args{
				a: a,
				b: b,
			},
			want: "1 hours",
		},
		{
			name: "Test Day",
			args: args{
				a: a,
				b: c,
			},
			want: "1 days 1 hours",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Logf("a = %v b = %v c = %v", a, b, c)
			if got := DiffTime(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("DiffTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiffTimes(t *testing.T) {
	type args struct {
		a time.Time
		b time.Time
	}
	a := time.Now().UTC()
	b := a.Add(time.Hour * 1)
	// c := a.Add(time.Hour * 25)
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
	}{
		{
			name: "Test True",
			args: args{
				a: a,
				b: b,
			},
			want:  a.Format("2006-01-02 15:04"),
			want1: b.Format("2006-01-02 15:04"),
			want2: "1h0m0s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := DiffTimes(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("DiffTimes() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DiffTimes() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("DiffTimes() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestConvertStringToDate(t *testing.T) {
	type args struct {
		s         string
		layoutISO string
	}
	data := "2023-02-02"
	layout := "2006-01-02"
	value, _ := time.Parse(layout, data)
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "Test Convert",
			args: args{
				s:         data,
				layoutISO: layout,
			},
			want:    value,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertStringToDate(tt.args.s, tt.args.layoutISO)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertStringToDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertStringToDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckArray(t *testing.T) {
	type args struct {
		data       string
		arrayCheck []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Exist",
			args: args{
				data:       "4",
				arrayCheck: []string{"1", "2", "3", "4"},
			},
			want: true,
		},
		{
			name: "Test Not Exist",
			args: args{
				data:       "5",
				arrayCheck: []string{"1", "2", "3", "4"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckArray(tt.args.data, tt.args.arrayCheck); got != tt.want {
				t.Errorf("CheckArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCurrency(t *testing.T) {
	type args struct {
		current float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test get currency",
			args: args{
				current: 3000,
			},
			want: "Rp. 3.000,00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCurrency(tt.args.current); got != tt.want {
				t.Errorf("GetCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayStringToArrayInt(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "Test Array",
			args: args{
				data: []string{"1", "2"},
			},
			want:    []int{1, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ArrayStringToArrayInt(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArrayStringToArrayInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayStringToArrayInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimZero(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Trim Zero",
			args: args{
				data: "123000",
			},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimZero(tt.args.data); got != tt.want {
				t.Errorf("TrimZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimiliarTo(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Similar To",
			args: args{
				data: []string{"a", "b", "c"},
			},
			want: "(a|b|c)%",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimiliarTo(tt.args.data); got != tt.want {
				t.Errorf("SimiliarTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueArray(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test unirue array",
			args: args{
				data: []string{"a", "a", "b"},
			},
			want: []string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueArray(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBeginningOfMonth(t *testing.T) {
	type args struct {
		date time.Time
	}
	data, _ := time.Parse("20060102", "20230202")
	result, _ := time.Parse("20060102", "20230201")
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Beginning of month",
			args: args{
				date: data,
			},
			want: result,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BeginningOfMonth(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BeginningOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndOfMonth(t *testing.T) {
	type args struct {
		date time.Time
	}
	data, _ := time.Parse("20060102", "20230202")
	result, _ := time.Parse("20060102", "20230228")
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Test end of month",
			args: args{
				date: data,
			},
			want: result,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndOfMonth(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleNan(t *testing.T) {
	type args struct {
		data float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Handle Nan",
			args: args{
				data: 0 / 100,
			},
			want: 0,
		},
		{
			name: "Handle Non Nan",
			args: args{
				data: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleNan(tt.args.data); got != tt.want {
				t.Errorf("HandleNan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundFloat(t *testing.T) {
	type args struct {
		input     float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Round Float",
			args: args{
				input:     29.56,
				precision: 1,
			},
			want: 29.6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundFloat(tt.args.input, tt.args.precision); got != tt.want {
				t.Errorf("RoundFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
