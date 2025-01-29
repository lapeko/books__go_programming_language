package engine

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
	tests := []struct {
		input  uint64
		output uint64
	}{
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
	if !utils.EqualType(eng, &binarySetEngine{}) {
		t.Errorf("New() should return %T type instead of %T", &binarySetEngine{}, eng)
	}
}

func TestPutIntoStorage(t *testing.T) {
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
		engine := binarySetEngine{storage: tt.storage}
		t.Run(tt.name, func(t *testing.T) {
			srcStorage := fmt.Sprint(tt.storage)
			if err := engine.Put(tt.num); err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !utils.Compare(engine.storage, tt.expectedStorage) {
				t.Errorf("Put %d into %s = %v. Not equal to %v", tt.num, srcStorage, engine.storage, tt.expectedStorage)
			}
		})
	}
}

func TestPutIntoStorageError(t *testing.T) {
	mut.Lock()
	orig := binaryEncode
	testErrMsg := "binaryEncode error"
	binaryEncode = func(num uint64) (uint64, error) {
		return 0, errors.New(testErrMsg)
	}
	defer func() {
		binaryEncode = orig
		mut.Unlock()
	}()

	engine := binarySetEngine{storage: []uint64{}}
	err := engine.Put(1)
	if err == nil || err.Error() != testErrMsg {
		t.Errorf("Expected error: %q have not cought", testErrMsg)
	}
}
