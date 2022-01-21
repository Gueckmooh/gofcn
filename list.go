package gofcn

func ListFilter[T any](list []T, predicate func(T) bool) []T {
	var filteredList []T
	for _, v := range list {
		if predicate(v) {
			filteredList = append(filteredList, v)
		}
	}
	return filteredList
}

func ListEqual[T comparable](l1, l2 []T) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := 0; i < len(l1); i += 1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}
