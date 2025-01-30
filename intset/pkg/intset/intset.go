package intset

type stringer interface {
	String() string
}

type Engine interface {
	stringer
	Put(num uint64)
	Delete(num uint64)
	Has(num uint64) bool
}

type intSet struct {
	Engine
}

type IntSet interface {
	stringer
	Add(num uint64)
	Delete(num uint64)
	Has(num uint64) bool
	Fill(nums []uint64)
}

func New(e Engine) IntSet {
	return &intSet{e}
}

func (i *intSet) Add(num uint64) {}

func (i *intSet) Delete(num uint64) {}

func (i *intSet) Has(num uint64) bool {
	return false
}

func (i *intSet) Fill(nums []uint64) {}

func (i *intSet) String() string {
	return ""
}
