package utils

import (
	"fmt"
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

func TestPutIntoStorage(t *testing.T) {
	tests := []struct {
		storage         []uint64
		num             uint64
		expectedStorage []uint64
	}{
		{storage: []uint64{0}, num: 0, expectedStorage: []uint64{1}},
		{storage: []uint64{0}, num: 1, expectedStorage: []uint64{2}},
		{storage: []uint64{1}, num: 1, expectedStorage: []uint64{3}},
		{storage: []uint64{0}, num: 3, expectedStorage: []uint64{8}},
		{storage: []uint64{0}, num: 64, expectedStorage: []uint64{0, 1}},
		{storage: []uint64{1, 1}, num: 63, expectedStorage: []uint64{exp63 + 1, 1}},
		{storage: []uint64{0}, num: 128, expectedStorage: []uint64{0, 0, 1}},
	}

	for _, tt := range tests {
		srcStorage := fmt.Sprint(tt.storage)
		res, err := putIntoStorage(tt.storage, tt.num)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !Compare(res, tt.expectedStorage) {
			t.Errorf("Put %d into %s = %v. Not equal to %v", tt.num, srcStorage, res, tt.expectedStorage)
		}
	}
}
