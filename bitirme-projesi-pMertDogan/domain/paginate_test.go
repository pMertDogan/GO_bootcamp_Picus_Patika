package domain

import "testing"

func TestCalcPageAndSize(t *testing.T) {
	type args struct {
		page     int
		pageSize int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"max page size ", args{0, 1000}, 1, 100},
		{"zero as page size ", args{0, 0}, 1, 10},
		{"negative as page size ", args{0, -763}, 1, 10},
		{"zero as page no ", args{0, 10}, 1, 10},
		{"six as page size ", args{0, 6}, 1, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CalcPageAndSize(tt.args.page, tt.args.pageSize)
			if got != tt.want {
				t.Errorf("CalcPageAndSize() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CalcPageAndSize() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
