package internal

func SlicesDiff(slice1, slice2 []string) []string {
	m := make(map[string]bool)
	for _, item := range slice2 {
		m[item] = true
	}

	var diff []string
	for _, item := range slice1 {
		if _, found := m[item]; !found {
			diff = append(diff, item)
		}
	}
	return diff
}
