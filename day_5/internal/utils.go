package internal

// ChunkIntSliceToPairs - chunks an int slice into a slice of slices, each containing two elements
func ChunkIntSliceToPairs(inputSlice []int) [][]int {
	var chunked [][]int
	for i := 0; i < len(inputSlice); i += 2 {
		if i+1 < len(inputSlice) {
			chunked = append(chunked, []int{inputSlice[i], inputSlice[i+1]})
		} else {
			// Handle the case where the length of inputSlice is odd
			chunked = append(chunked, []int{inputSlice[i]})
		}
	}
	return chunked
}
