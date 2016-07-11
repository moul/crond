package crond

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("Testing New", t, func() {
		manager := New()
		So(manager, ShouldHaveSameTypeAs, Manager{})
	})
}
