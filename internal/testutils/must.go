package testutils

import "testing"

func Must[T any](t *testing.T) func(r T, err error) T {
	return func(r T, err error) T {
		t.Helper()

		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}

		return r
	}
}
