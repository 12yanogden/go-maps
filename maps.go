package maps

// Return true if the key given is in the map given. Else, false.
func HasKey[KEY comparable, A any](mapp map[KEY]A, key KEY) bool {
	for mappKey := range mapp {
		if key == mappKey {
			return true
		}
	}

	return false
}
