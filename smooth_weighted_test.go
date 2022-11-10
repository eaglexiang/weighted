package weighted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSW_Next(t *testing.T) {
	w := &SW{}
	w.Add("server1", 5)
	w.Add("server2", 2)
	w.Add("server3", 3)

	results := make(map[string]int)

	for i := 0; i < 100; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if results["server1"] != 50 || results["server2"] != 20 || results["server3"] != 30 {
		t.Error("the algorithm is wrong")
	}

	w.Reset()
	results = make(map[string]int)

	for i := 0; i < 100; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if results["server1"] != 50 || results["server2"] != 20 || results["server3"] != 30 {
		t.Error("the algorithm is wrong")
	}

	w.RemoveAll()
	w.Add("server1", 7)
	w.Add("server2", 9)
	w.Add("server3", 13)

	results = make(map[string]int)

	for i := 0; i < 29000; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if results["server1"] != 7000 || results["server2"] != 9000 || results["server3"] != 13000 {
		t.Error("the algorithm is wrong")
	}
}

func TestSW_SetWeight(t *testing.T) {
	w := &SW{}
	w.Add("server1", 5)
	w.Add("server2", 2)
	w.Add("server3", 3)
	w.Add("server4", 4)

	w.SetWeight([]WeightItem{
		{
			ID:     "server1",
			Weight: 1,
		},
		{
			ID:     "server2",
			Weight: 0,
		},
		{
			ID:     "server4",
			Weight: 2,
		},
		{
			ID:     "server5",
			Weight: 5,
		},
	})

	expected := &SW{}
	expected.Add("server1", 1)
	expected.Add("server4", 2)
	expected.Add("server5", 5)

	expectedBuf, err := expected.MarshalJSON()
	if !assert.NoError(t, err) {
		return
	}
	wBuf, err := w.MarshalJSON()
	if !assert.NoError(t, err) {
		return
	}

	if !assert.Equal(t, string(expectedBuf), string(wBuf)) {
		return
	}
}
