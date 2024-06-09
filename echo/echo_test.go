package echo_test

import (
	"testing"

	"github.com/hinscx/goecho/echo"
)

func TestEcho(t *testing.T) {

	want := "testin"
	got := echo.Echo(want)

	if want != got {
		t.Errorf("got %s want %s given, %s", got, want, want)
	}
}
