package intset

//func TestNew(t *testing.T) {
//	i := New()
//	if !utils.EqualType(i, &intSet{}) {
//		t.Errorf("New() should return %T type instead of %T", &intSet{}, i)
//	}
//}

//func TestAdd(t *testing.T) {
//	tests := []struct{
//		addValues []uint64
//		expectedStorage []uint64
//	}{
//		{addValues: []uint64{1}, expectedStorage: []uint64{1 << 1 - 1}},
//	}
//	i := New()
//	testValues :=
//	for _, val := range testValues {
//		i.Add(val)
//	}
//	iSet, ok := i.(*intSet)
//	if ok {
//		t.Errorf("type casting %v to %v failure", reflect.TypeOf(i), reflect.TypeOf(intSet{}))
//	}
//	expected := []uint64{1}
//	if !testutils.Compare(iSet.storage, []uint64{1}) {
//		t.Errorf("TestAdd. Expected: %v, got: %v", expected, iSet.storage)
//	}
//}
