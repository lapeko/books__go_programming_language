package utils

import (
	"fmt"
)

const maxLimit = 64

func binaryEncode(num uint64) (uint64, error) {
	if num >= maxLimit {
		return 0, fmt.Errorf("too big payload %d. Should be less %d", num, maxLimit)
	}
	return 1 << num, nil
}

func putIntoStorage(storage []uint64, num uint64) ([]uint64, error) {
	idx, rest := int(num/64), num%64
	if len(storage)-1 < idx {
		for i := len(storage) - 1; i < idx; i++ {
			storage = append(storage, 0)
		}
	}
	bin, err := binaryEncode(rest)
	if err != nil {
		return nil, err
	}
	storage[idx] |= bin
	return storage, nil
}
