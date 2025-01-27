package slice

func GetIntersection(slice1, slice2 []string) (slice []string) {
	m := ToMap(slice1)
	for _, i := range slice2 {
		if m[i] {
			slice = append(slice, i)
		}
	}
	return
}

func RmDuplicates(slice []string) (s []string) {
	m := make(map[string]bool)
	for _, i := range slice {
		if !m[i] {
			m[i] = true
			s = append(s, i)
		}
	}
	return
}

func ToMap(slice []string) map[string]bool {
	m := make(map[string]bool)
	for _, i := range slice {
		m[i] = true
	}
	return m
}
