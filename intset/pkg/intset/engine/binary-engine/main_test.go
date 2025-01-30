package binary_engine

import (
	"errors"
	"fmt"
	"github.com/lapeko/books__go_programming_language/intset/pkg/intset/utils"
	"sync"
	"testing"
)

const (
	exp0 = 1 << (iota * 10)
	exp10
	exp20
	exp30
	exp40
	exp50
	exp60
	exp63 = exp60 << 3
)

var mut sync.Mutex

func TestBinaryEncode(t *testing.T) {
	tests := []struct{ input, output uint64 }{
		{input: 0, output: exp0},
		{input: 10, output: exp10},
		{input: 20, output: exp20},
		{input: 30, output: exp30},
		{input: 40, output: exp40},
		{input: 50, output: exp50},
		{input: 60, output: exp60},
		{input: 63, output: exp63},
	}
	for _, tt := range tests {
		result, err := binaryEncode(tt.input)
		if err != nil {
			t.Errorf("Not expected error %v", err)
		}
		if result != tt.output {
			t.Errorf("Expected %d. Received %d", tt.output, result)
		}
	}
}

func TestBinaryEncodeError(t *testing.T) {
	_, err := binaryEncode(maxLimit)
	if err == nil {
		t.Error("Expected to get an error")
	}
}

func TestNew(t *testing.T) {
	eng := New()
	if !utils.EqualType(eng, &binaryEngine{}) {
		t.Errorf("New() should return %T type instead of %T", &binaryEngine{}, eng)
	}
}

func TestBinarySetEngine_Put(t *testing.T) {
	tests := []struct {
		name            string
		storage         []uint64
		num             uint64
		expectedStorage []uint64
	}{
		{name: "Insert 0", storage: []uint64{0}, num: 0, expectedStorage: []uint64{1}},
		{name: "Insert 1", storage: []uint64{0}, num: 1, expectedStorage: []uint64{2}},
		{name: "Insert 1 into existing", storage: []uint64{1}, num: 1, expectedStorage: []uint64{3}},
		{name: "Insert 3", storage: []uint64{0}, num: 3, expectedStorage: []uint64{8}},
		{name: "Insert 64 (creates second uint64)", storage: []uint64{0}, num: 64, expectedStorage: []uint64{0, 1}},
		{name: "Insert 63 into populated storage", storage: []uint64{1, 1}, num: 63, expectedStorage: []uint64{exp63 + 1, 1}},
		{name: "Insert 128 (creates third uint64)", storage: []uint64{0}, num: 128, expectedStorage: []uint64{0, 0, 1}},
	}

	for _, tt := range tests {
		engine := binaryEngine{storage: tt.storage}
		t.Run(tt.name, func(t *testing.T) {
			srcStorage := fmt.Sprint(tt.storage)
			engine.Put(tt.num)
			if !utils.Compare(engine.storage, tt.expectedStorage) {
				t.Errorf("Put %d into %s = %v. Not equal to %v", tt.num, srcStorage, engine.storage, tt.expectedStorage)
			}
		})
	}
}

func TestBinarySetEngine_Put_Error(t *testing.T) {
	testErrMsg := "binaryEncode error"
	mut.Lock()
	orig := binaryEncode
	binaryEncode = func(num uint64) (uint64, error) {
		return 0, errors.New(testErrMsg)
	}
	defer func() {
		binaryEncode = orig
		mut.Unlock()
		if r := recover(); r != nil {
			if fmt.Sprintf("%s", r) != testErrMsg {
				t.Errorf("not expected received error: %q. Expected: %q", r, testErrMsg)
			}
		} else {
			t.Errorf("expected panic but got none")
		}
	}()

	engine := binaryEngine{storage: []uint64{}}
	engine.Put(1)
}

var substracts = []struct{ subtrahend, subtractor, expect uint64 }{
	{0b00000001, 0, 0b00000000},
	{0b11111111, 0, 0b11111110},
	{0b11111110, 0, 0b11111110},
	{0b11111111, 0, 0b11111110},
	{0b11111110, 1, 0b11111100},
	{0b10000000, 7, 0b00000000},
	{0b10010000, 4, 0b10000000},
}

func TestBinarySetEngine_Delete(t *testing.T) {
	for _, tt := range substracts {
		engine := binaryEngine{storage: []uint64{tt.subtrahend}}
		name := fmt.Sprintf("%d Delete(%d) expects %d", engine.storage[0], tt.subtractor, tt.expect)
		t.Run(name, func(t *testing.T) {
			engine.Delete(tt.subtractor)
			if engine.storage[0] != tt.expect {
				t.Errorf("%s. Received %d", name, engine.storage[0])
			}
		})
	}

	engine := binaryEngine{storage: []uint64{0}}
	engine.Delete(maxLimit)
	if !utils.Compare(engine.storage, []uint64{0}) {
		t.Errorf("%v Delete(%d) expects %d", []uint64{0}, maxLimit, []uint64{0})
	}
}

func TestBinarySetEngine_Delete_Error(t *testing.T) {
	testErrMsg := "subtractCandidate error"
	mut.Lock()
	orig := subtractCandidate
	subtractCandidate = func(storage, bt uint64) (uint64, error) {
		return 0, errors.New(testErrMsg)
	}
	defer func() {
		subtractCandidate = orig
		mut.Unlock()
		if r := recover(); r != nil {
			if fmt.Sprintf("%s", r) != testErrMsg {
				t.Errorf("not expected received error: %q. Expected: %q", r, testErrMsg)
			}
		} else {
			t.Errorf("expected panic but got none")
		}
	}()

	engine := binaryEngine{storage: []uint64{0}}
	engine.Delete(1)
}

func TestBinarySetEngine_Has(t *testing.T) {
	tests := []struct {
		initStorage []uint64
		checkNum    uint64
		expected    bool
	}{
		{[]uint64{0b00000000}, 0, false},
		{[]uint64{0b00000001}, 0, true},
		{[]uint64{0b00000000}, 64, false},
		{[]uint64{0b10000000}, 7, true},
		{[]uint64{0, 0, 1}, 128, true},
	}

	for _, tt := range tests {
		e := binaryEngine{storage: tt.initStorage}
		if res := e.Has(tt.checkNum); res != tt.expected {
			t.Errorf("binaryEngine{%v}.Has(%d) = %t when expected %t", tt.initStorage, tt.checkNum, res, tt.expected)
		}
	}
}

func TestBinarySetEngine_String(t *testing.T) {
	tests := []struct {
		storage []uint64
		expect  string
	}{
		{[]uint64{0b00000000}, "{ }"},
		{[]uint64{0b00000001}, "{ 0 }"},
		{[]uint64{0b00000011}, "{ 0 1 }"},
		{[]uint64{0b00001010}, "{ 1 3 }"},
		{[]uint64{0b11111111, 0, 1}, "{ 0 1 2 3 4 5 6 7 128 }"},
	}

	for _, tt := range tests {
		e := binaryEngine{storage: tt.storage}
		if res := e.String(); res != tt.expect {
			t.Errorf("binaryEngine{%v}.String() = %q. Expected: %q", tt.storage, res, tt.expect)
		}
	}
}

func TestSubtractCandidate(t *testing.T) {
	for _, tt := range substracts {
		res, err := subtractCandidate(tt.subtrahend, tt.subtractor)
		if err != nil {
			t.Errorf("unexpected error %q when subtractCandidate(%d, %d) expects %d", err, tt.subtrahend, tt.subtractor, tt.expect)
		}
		if res != tt.expect {
			t.Errorf("subtractCandidate(%d, %d) = %d. Expected %d", tt.subtrahend, tt.subtractor, res, tt.expect)
		}
	}
}

func TestSubtractCandidateError(t *testing.T) {
	if _, err := subtractCandidate(0, maxLimit); err == nil {
		t.Errorf("subtractCandidate(%d) should return an error", maxLimit)
	}
}
