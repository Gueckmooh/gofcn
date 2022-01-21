package gofcn_test

import (
	"gofcn"
	"testing"
)

func TestFilterIntList(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	filteredList := gofcn.ListFilter(list, func(v int) bool { return (v % 2) == 0 })
	if !gofcn.ListEqual(filteredList, []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20}) {
		t.Fail()
	}
}

func TestFilterStringList(t *testing.T) {
	list := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}
	filteredList := gofcn.ListFilter(list, func(v string) bool { return len(v) == 1 })
	if !gofcn.ListEqual(filteredList, []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}) {
		t.Fail()
	}
}
