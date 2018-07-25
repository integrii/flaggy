package flaggy_test

import (
	"os"
	"testing"

	"github.com/integrii/flaggy"
)

func TestMain(m *testing.M) {
	flaggy.PanicInsteadOfExit = true
	os.Exit(m.Run())
}
