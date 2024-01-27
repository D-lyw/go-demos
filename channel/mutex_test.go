package channel

import "testing"

func TestRunConcurrencyWithLock(t *testing.T) {
	RunConcurrencyWithLock()
}

func TestRunConcurrencyWithoutLock(t *testing.T) {
	RunConcurrencyWithoutLock()
}
