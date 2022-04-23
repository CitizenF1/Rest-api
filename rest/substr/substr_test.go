package substr

import (
	"testing"
)

var Tests = []struct {
	str         string
	expectedInt int
	expectedStr string
}{
	{
		str:         "MaxSubstring",
		expectedInt: 12,
		expectedStr: "MSabginrstux",
	},
	{
		str:         "SomeTEXt12342",
		expectedInt: 12,
		expectedStr: "1234ESTXemot",
	},
}

func TestMaxSubstring(t *testing.T) {
	for _, tt := range Tests {
		s, i := MaxSubstring(tt.str)
		if i != tt.expectedInt || s != tt.expectedStr {
			t.Errorf("MaxSubstring(%s) => %s, %d, want %s, %d", tt.str, s, i, tt.expectedStr, tt.expectedInt)
		}
	}
}
