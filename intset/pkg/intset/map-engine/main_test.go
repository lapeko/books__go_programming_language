package map_engine

import (
	"github.com/lapeko/books__go_programming_language/intset/pkg/intset/utils"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	e := New()
	if !utils.EqualType(e, &mapEngine{}) {
		t.Error("New engine should be a map engine")
	}
}

func TestMapEngine_String(t *testing.T) {
	tests := []struct {
		initValues []uint64
		expect     string
	}{
		{nil, "{ }"},
		{[]uint64{0}, "{ 0 }"},
		{[]uint64{0, 1}, "{ 0 1 }"},
		{[]uint64{0, 1, 2, 3, 4, 5, 6, 7, 128}, "{ 0 1 2 3 4 5 6 7 128 }"},
	}

	for _, tt := range tests {
		e := mapEngine{storage: make(map[uint64]struct{})}
		for _, v := range tt.initValues {
			e.storage[v] = struct{}{}
		}
		if got := e.String(); got != tt.expect {
			t.Errorf("MapEngine.String() = %v, want %v", got, tt.expect)
		}
	}
}

func TestMapEngine_Delete(t *testing.T) {
	tests := []struct {
		initValues []uint64
		delValues  []uint64
		expect     []uint64
	}{
		{[]uint64{1}, []uint64{}, []uint64{1}},
		{[]uint64{1}, []uint64{1}, []uint64{}},
		{[]uint64{0}, []uint64{1}, []uint64{0}},
		{[]uint64{1, 2, 128}, []uint64{1}, []uint64{2, 128}},
	}

	for _, tt := range tests {
		e := mapEngine{storage: make(map[uint64]struct{})}
		for _, v := range tt.initValues {
			e.storage[v] = struct{}{}
		}
		for _, v := range tt.delValues {
			e.Delete(v)
		}
		for _, v := range tt.delValues {
			if e.Has(v) {
				t.Errorf("MapEngine(%v).Delete(%v). %d should not exist anymore", tt.initValues, tt.delValues, v)
			}
		}
	}
}

func TestMapEngine_Has(t *testing.T) {
	tests := []struct {
		initValues []uint64
		checkValue uint64
		expect     bool
	}{
		{[]uint64{}, 0, false},
		{[]uint64{0}, 0, true},
		{[]uint64{12, 45}, 45, true},
		{[]uint64{140}, 45, false},
		{[]uint64{140}, 140, true},
	}

	for _, tt := range tests {
		e := mapEngine{storage: make(map[uint64]struct{})}
		for _, v := range tt.initValues {
			e.storage[v] = struct{}{}
		}
		if got := e.Has(tt.checkValue); got != tt.expect {
			t.Errorf("MapEngine(%v).Has(%v) = %v, expected %v", tt.initValues, tt.checkValue, got, tt.expect)
		}
	}
}

func TestMapEngine_Put(t *testing.T) {
	tests := []struct {
		storage     map[uint64]struct{}
		addValues   []uint64
		expectedMap map[uint64]struct{}
	}{
		{map[uint64]struct{}{}, []uint64{}, map[uint64]struct{}{}},
		{map[uint64]struct{}{}, []uint64{0}, map[uint64]struct{}{0: {}}},
		{map[uint64]struct{}{}, []uint64{0, 1, 2, 5}, map[uint64]struct{}{0: {}, 1: {}, 2: {}, 5: {}}},
		{map[uint64]struct{}{1: {}, 3: {}}, []uint64{1, 2, 5}, map[uint64]struct{}{1: {}, 2: {}, 3: {}, 5: {}}},
	}

	for _, tt := range tests {
		e := mapEngine{storage: tt.storage}
		for _, num := range tt.addValues {
			e.Put(num)
		}
		if !reflect.DeepEqual(tt.expectedMap, e.storage) {
			t.Errorf("Put(%d): got %v, want %v", tt.addValues, e.storage, tt.expectedMap)
		}
	}
}
