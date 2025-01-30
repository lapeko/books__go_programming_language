package binary_engine

import (
	"fmt"
	"github.com/lapeko/books__go_programming_language/intset/pkg/intset"
	"strings"
)

const maxLimit = 64

type BinarySet interface {
	intset.Engine
}

type binaryEngine struct {
	storage []uint64
}

func New() intset.Engine {
	return &binaryEngine{storage: []uint64{0}}
}

func (b *binaryEngine) Put(num uint64) {
	idx, rest := int(num/maxLimit), num%maxLimit
	if len(b.storage)-1 < idx {
		for i := len(b.storage) - 1; i < idx; i++ {
			b.storage = append(b.storage, 0)
		}
	}
	bin, err := binaryEncode(rest)
	if err != nil {
		panic(err)
	}
	b.storage[idx] = b.storage[idx] | bin
}

func (b *binaryEngine) Delete(num uint64) {
	idx, rest := int(num/maxLimit), num%maxLimit
	if idx >= len(b.storage) {
		return
	}
	res, err := subtractCandidate(b.storage[idx], rest)
	if err != nil {
		panic(err)
	}
	b.storage[idx] = res
}

func (b *binaryEngine) Has(num uint64) bool {
	idx, rest := int(num/maxLimit), num%maxLimit
	if idx >= len(b.storage) {
		return false
	}
	return b.storage[idx]&(1<<rest) != 0
}

func (b *binaryEngine) String() string {
	sb := strings.Builder{}
	sb.WriteString("{ ")
	for idx := range b.storage {
		for i := uint64(0); i < maxLimit; i++ {
			currentNum := uint64(maxLimit*idx) + i
			if b.Has(currentNum) {
				sb.WriteString(fmt.Sprintf("%d ", currentNum))
			}
		}
	}
	sb.WriteString("}")
	return sb.String()
}

var binaryEncode = func(num uint64) (uint64, error) {
	if num >= maxLimit {
		return 0, fmt.Errorf("too big payload %d. Should be less %d", num, maxLimit)
	}
	return 1 << num, nil
}

var subtractCandidate = func(storage, bt uint64) (uint64, error) {
	if bt >= maxLimit {
		return 0, fmt.Errorf("too big payload %d. Should be less %d", bt, maxLimit)
	}
	return storage & ^(1 << bt), nil
}
