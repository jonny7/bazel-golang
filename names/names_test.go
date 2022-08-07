package names

import "testing"

func TestName(t *testing.T) {
	name := Name()
	if len(name) < 1 {
		t.Error("expected a name with a length greater than 1")
	}
}
