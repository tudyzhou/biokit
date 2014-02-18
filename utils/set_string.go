package utils

func SetString(org []string) []string {
	m := map[string]bool{}
	for _, v := range org {
		if _, seen := m[v]; !seen {
			org[len(m)] = v
			m[v] = true
		}
	}
	return org[:len(m)]
}
