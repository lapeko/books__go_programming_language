package map_engine

import (
	"bytes"
	"fmt"
	"github.com/lapeko/books__go_programming_language/intset/pkg/intset"
	"slices"
)

type mapEngine struct {
	storage map[uint64]struct{}
}

func New() intset.Engine {
	return &mapEngine{storage: make(map[uint64]struct{})}
}

func (m mapEngine) String() string {
	s := make([]uint64, 0, len(m.storage))
	for num := range m.storage {
		s = append(s, num)
	}
	slices.Sort(s)

	buf := new(bytes.Buffer)
	for _, num := range s {
		buf.WriteString(fmt.Sprintf("%d ", num))
	}

	return fmt.Sprintf("{ %s}", buf)
}

func (m mapEngine) Put(num uint64) {
	m.storage[num] = struct{}{}
}

func (m mapEngine) Delete(num uint64) {
	delete(m.storage, num)
}

func (m mapEngine) Has(num uint64) bool {
	_, ok := m.storage[num]
	return ok
}
