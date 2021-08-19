package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 2, 2: 1, 3: 3}
	t.Log(m1[1])
	m1[3] = 4
	t.Logf("len m1=%d", len(m1))

}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[1]; ok {
		t.Logf("Key 1's value is %d", v)
	} else {
		t.Log("key 1 is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	m := map[int]int{1: 2, 2: 3, 3: 4, 4: 5}
	for k, v := range m {
		t.Log(k, v)
	}
}
