package intset

import (
	"github.com/lapeko/books__go_programming_language/intset/pkg/testutils"
	"testing"
)

func TestNew(t *testing.T) {
	i := New()
	testutils.EqualType(t, i, &intSet{})
}
