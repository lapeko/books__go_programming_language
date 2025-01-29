package intset

type intSet struct {
}

type stringer interface {
	String() string
}

type IntSet interface {
	stringer
	Add(num int)
	Delete(num int)
	Has(num int) bool
	Fill(nums []int)
}

func New() IntSet {
	return &intSet{}
}

func (i *intSet) Add(num int) {}

func (i *intSet) Delete(num int) {}

func (i *intSet) Has(num int) bool {
	return false
}

func (i *intSet) Fill(nums []int) {}

func (i *intSet) String() string {
	return ""
}
