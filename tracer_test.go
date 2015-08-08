package trace

import (
	"testing"
)

func TestNew(t *testing.T) {
	// go treats any function that starts with Test (taking a single *testing.T)
	// argument as a unit test and will execute it when we run our tests
	t.Error("No tests available. Yet!")
}
