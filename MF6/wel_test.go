package MF6

import (
	"testing"
	"time"
)

// tests of the wel file generator
type testStruct struct {
	dt      time.Time
	nd      int
	pumping float64
	well    int
}

func (t testStruct) date() time.Time {
	return t.dt
}

func (t testStruct) node() int {
	return t.nd
}

func (t testStruct) value() float64 {
	return t.pumping
}

var d0 = testStruct{dt: time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC), nd: 1,
	pumping: -12, well: 101}
var d1 = testStruct{dt: time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC), nd: 2,
	pumping: -14, well: 102}
var d2 = testStruct{dt: time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC), nd: 1,
	pumping: -16, well: 103}
var d3 = testStruct{dt: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC), nd: 2,
	pumping: -18, well: 104}

var data = []fileData{d0, d1, d2, d3}

func TestWel(t *testing.T) {
	if err := Wel("test", data); err != nil {
		t.Error("Wel function errored with", err)
	}
}

func Test_filterDataByDate(t *testing.T) {
	dt := time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC)

	d, err := filterDataByDate(dt, data)
	if err != nil {
		t.Error("function returned an error and shouldn't have")
	}

	if len(d) != 2 {
		t.Errorf("function should have returned two records but returned %d instead", len(d))
	}
}
