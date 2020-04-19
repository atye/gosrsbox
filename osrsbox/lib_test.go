package osrsbox

import "testing"

func Test_makeValidItemName(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"abyssal Whip",
			"Abyssal whip",
		},
	}

	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			result := makeValidItemName(tc.in)

			if result != tc.out {
				t.Errorf("expected %s, got %s", tc.out, result)
			}
		})
	}
}
