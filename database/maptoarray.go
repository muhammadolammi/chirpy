package database

func MapToArray(m map[int]Chirp) []Chirp {
	arr := make([]Chirp, len(m))
	for i, v := range m {
		arr[i-1] = v

	}
	return arr
}
