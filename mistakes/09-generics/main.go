func getKeys(m any) ([]any , error) {
	switch t := m.(type) {
	default:
		return nil , fmt.Errorf("unknown type: %T", t)
	case map[string]int:
		var keys []any
		for k := range t {
			keys = append(keys, k)
		}

		return keys, nil
	case map[int]string:
		var keys []any
for k := range t {
keys = append(keys, k)
}
		
	}
}
