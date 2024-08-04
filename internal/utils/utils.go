package utils

func Contains[T comparable](item T, list []T) bool {
	for _, content := range list {
		if item == content {
			return true
		}
	}
	return false
}
