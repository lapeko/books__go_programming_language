package engine

import (
	"fmt"
	"github.com/lapeko/books__go_programming_language/intset/pkg/intset"
)

const maxLimit = 64

type BinarySet interface {
	intset.Engine
}

type binarySetEngine struct {
	storage []uint64
}

func New() intset.Engine {
	return &binarySetEngine{}
}

func (b *binarySetEngine) Put(num uint64) error {
	idx, rest := int(num/64), num%64
	if len(b.storage)-1 < idx {
		for i := len(b.storage) - 1; i < idx; i++ {
			b.storage = append(b.storage, 0)
		}
	}
	bin, err := binaryEncode(rest)
	if err != nil {
		return err
	}
	b.storage[idx] |= bin
	return nil
}

func (b *binarySetEngine) Delete(num uint64) {

}

func (b *binarySetEngine) Has(num uint64) bool {
	return false
}

func (b *binarySetEngine) String() string {
	return ""
}

var binaryEncode = func(num uint64) (uint64, error) {
	if num >= maxLimit {
		return 0, fmt.Errorf("too big payload %d. Should be less %d", num, maxLimit)
	}
	return 1 << num, nil
}
