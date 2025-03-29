package utils_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/suger-131997/dein/internal/utils"
)

func TestPriorityQueue(t *testing.T) {
	type testStruct struct {
		Num int
	}

	pq := utils.NewPriorityQueue(func(i, j testStruct) bool {
		return i.Num < j.Num
	})

	utils.Push(pq, testStruct{Num: 5})
	utils.Push(pq, testStruct{Num: 2})
	utils.Push(pq, testStruct{Num: 3})
	utils.Push(pq, testStruct{Num: 1})
	utils.Push(pq, testStruct{Num: 6})
	utils.Push(pq, testStruct{Num: 4})

	got := make([]testStruct, 0, pq.Len())
	for pq.Len() != 0 {
		got = append(got, utils.Pop(pq))
	}

	want := []testStruct{{Num: 1}, {Num: 2}, {Num: 3}, {Num: 4}, {Num: 5}, {Num: 6}}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("unexpected result (-got +want):\n%s", diff)
	}
}
