package ascii

import (
	"testing"
)

func TestGreetingASCII(t *testing.T) {
	v := GreetingASCII()
	for i := 0; i < len(v); i++ {
		if (isASCII(string(v[i])) == false){
			// throws error
			t.Fail()
			break
		}
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

