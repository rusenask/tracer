package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	// go treats any function that starts with Test (taking a single *testing.T)
	// argument as a unit test and will execute it when we run our tests
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		tracer.Trace("Trace package is here.")
		if buf.String() != "Trace package is here.\n" {
			t.Errorf("Trace should not write '%s.'", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("something")
}
