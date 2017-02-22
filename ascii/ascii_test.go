package ascii

import (
	"testing"
)

func TestGreetingASCII(t *testing.T) {
	if !(isASCII(GreetingASCII())){
		t.Fail()
	}
}



func isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}

