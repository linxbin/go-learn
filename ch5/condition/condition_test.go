package condition

import "testing"

func TestCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i%2 == 0 {
		case true:
			t.Logf("%d is Even", i)
		case false:
			t.Logf("%d is Odds", i)
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Logf("%d is Even", i)
		case i%2 != 0:
			t.Logf("%d is Odds", i)
		default:
			t.Log("unknown")
		}
	}
}
