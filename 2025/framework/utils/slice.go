package utils

func RemoveIndex[T any](s []T, index int) []T {
	ret := []T{}
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
