package num

import "testing"

func TestCheck(t *testing.T) {
	result := ToInt("1234")
	if result != 1234 {
		t.Errorf("Not equal = %d; expected 1234", result)
	}
}
