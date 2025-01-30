package engine

import (
	"fmt"
	"github.com/lapeko/books__go_programming_language/intset/pkg/intset"
	"strings"
)

const maxLimit = 64

type BinarySet interface {
	intset.Engine
}

type binarySetEngine struct {
	storage []uint64
}

func New() intset.Engine {
	return &binarySetEngine{storage: []uint64{0}}
}

func (b *binarySetEngine) Put(num uint64) {
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

func (b *binarySetEngine) Delete(num uint64) {
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

func (b *binarySetEngine) Has(num uint64) bool {
	idx, rest := int(num/maxLimit), num%maxLimit
	if idx >= len(b.storage) {
		return false
	}
	return b.storage[idx] == b.storage[idx]|1<<rest
}

func (b *binarySetEngine) String() string {
	sb := strings.Builder{}
	sb.WriteString("{ ")
	for idx, encodedNum := range b.storage {
		for i := uint64(1); i <= maxLimit; i++ {
			if encodedNum == encodedNum&i {
				sb.WriteString(fmt.Sprintf("%d ", uint64(idx)*maxLimit+i))
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
