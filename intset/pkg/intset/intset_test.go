package intset

import (
	"testing"

	"github.com/lapeko/books__go_programming_language/intset/pkg/intset/utils"
)

type engineMock struct {
	stringCalled     bool
	putCalledWith    []uint64
	deleteCalledWith uint64
	hasCalledWith    uint64
	Engine
}

func (e *engineMock) String() string {
	e.stringCalled = true
	return ""
}

func (e *engineMock) Put(num uint64) {
	e.putCalledWith = append(e.putCalledWith, num)
}

func (e *engineMock) Delete(num uint64) {
	e.deleteCalledWith = num
}

func (e *engineMock) Has(num uint64) bool {
	e.hasCalledWith = num
	return false
}

func TestNew(t *testing.T) {
	set := New(&engineMock{})
	if !utils.EqualType(set, &intSet{}) {
		t.Errorf("New() should return %T type instead of %T", &intSet{}, set)
	}
}

func TestAdd(t *testing.T) {
	tests := []uint64{1, 3, 4, 5, 3, 5, 6, 0}
	em := engineMock{}
	set := New(&em)
	for idx, testPayload := range tests {
		set.Add(testPayload)
		if em.putCalledWith[idx] != testPayload {
			t.Errorf("engineMock.Put was not called with %d", testPayload)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []uint64{1, 3, 4, 5, 3, 5, 6, 0}
	em := engineMock{}
	set := New(&em)
	for _, testPayload := range tests {
		set.Delete(testPayload)
		if em.deleteCalledWith != testPayload {
			t.Errorf("engineMock.Delete was not called with %d", testPayload)
		}
	}
}

func TestHas(t *testing.T) {
	tests := []uint64{1, 3, 4, 5, 3, 5, 6, 0}
	em := engineMock{}
	set := New(&em)
	for _, testPayload := range tests {
		set.Has(testPayload)
		if em.hasCalledWith != testPayload {
			t.Errorf("engineMock.Has was not called with %d", testPayload)
		}
	}
}

func TestFill(t *testing.T) {
	fillData := []uint64{1, 3, 4, 5, 3, 5, 6, 0}
	em := engineMock{}
	set := New(&em)
	set.Fill(fillData)

	for idx, testPayload := range fillData {
		if em.putCalledWith[idx] != testPayload {
			t.Errorf("engineMock.Put was not called with %d", testPayload)
		}
	}
}

func TestString(t *testing.T) {
	em := engineMock{}
	set := New(&em)
	_ = set.String()
	if !em.stringCalled {
		t.Errorf("engineMock.String was not called")
	}
}
