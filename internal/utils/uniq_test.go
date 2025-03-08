package utils_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/suger-131997/dein/internal/utils"
	"testing"
)

func TestUniq(t *testing.T) {
	t.Run("uniq for int", func(tt *testing.T) {
		uniq := utils.Uniq([]int{1, 2, 3, 1, 2, 3, 4})
		if diff := cmp.Diff([]int{1, 2, 3, 4}, uniq, cmpopts.SortSlices(func(i, j string) bool {
			return i < j
		})); diff != "" {
			tt.Errorf("Uniq failed (-got +want):\n%s", diff)
		}
	})

	t.Run("uniq for struct", func(tt *testing.T) {
		type t struct {
			Num int
		}
		uniq := utils.Uniq([]t{{1}, {2}, {3}, {1}, {2}, {3}, {4}})
		if diff := cmp.Diff([]t{{1}, {2}, {3}, {4}}, uniq, cmpopts.SortSlices(func(i, j t) bool {
			return i.Num < j.Num
		})); diff != "" {
			tt.Errorf("Uniq failed (-got +want):\n%s", diff)
		}
	})
}
