package utils

func Uniq[T comparable](l []T) []T {
	m := make(map[T]struct{})
	r := make([]T, 0)
	for _, v := range l {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}
