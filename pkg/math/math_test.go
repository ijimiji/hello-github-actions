package math

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{
			a:    1,
			b:    2,
			want: 3,
		},
		{
			a:    -1,
			b:    1,
			want: 0,
		},
	}
	for _, tc := range tests {
		if Add(tc.a, tc.b) != tc.want {
			t.Errorf("Add(%d, %d) != %d", tc.a, tc.b, tc.want)
		}
	}
}
