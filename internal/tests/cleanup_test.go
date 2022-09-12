package tests

import (
	"reflect"
	"runtime"
	"testing"
)

func ATestCleanup(t *testing.T) {
	var cleanups []int
	t.Run("test", func(t *testing.T) {
		t.Cleanup(func() { cleanups = append(cleanups, 1) })
		t.Cleanup(func() { cleanups = append(cleanups, 2) })
	})
	if got, want := cleanups, []int{2, 1}; !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected cleanup record: got %v want %v", got, want)
	}
}

func ATestConcurrentCleanup(t *testing.T) {
	cleanups := 0
	t.Run("test", func(t *testing.T) {
		done := make(chan struct{})
		for i := 0; i < 2; i++ {
			i := i
			go func() {
				t.Cleanup(func() {
					cleanups |= 1 << i
				})
				done <- struct{}{}
			}()
			<-done
			<-done
		}
	})
	if cleanups != 1|2 {
		t.Errorf("unexpected cleanup: got %d want 3", cleanups)
	}
}

func ATestCleanupCalledEvenAfterGoexit(t *testing.T) {
	cleanups := 0
	t.Run("test", func(t *testing.T) {
		t.Cleanup(func() {
			cleanups++
		})
		t.Cleanup(func() {
			runtime.Goexit()
		})
	})
	if cleanups != 1 {
		t.Errorf("unexpected cleanup count: got %d want 1", cleanups)
	}
}

func ATestRunCleanup(t *testing.T) {
	outerCleanup := 0
	innerCleanup := 0
	t.Run("test", func(t *testing.T) {
		t.Cleanup(func() { outerCleanup++ })
		t.Run("x", func(t *testing.T) {
			t.Cleanup(func() { innerCleanup++ })
		})
	})
	if innerCleanup != 1 {
		t.Errorf("unexpected inner cleanup count; got %d want 1", innerCleanup)
	}
	if outerCleanup != 1 {
		t.Errorf("unexpected outer cleanup count; got %d want 0", outerCleanup)
	}
}
