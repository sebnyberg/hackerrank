package taumandbday

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	// Write your code here
	return min(int64(b)*(int64(z)+int64(wc)), int64(b)*int64(bc)) + min(int64(w)*(int64(z)+int64(bc)), int64(w)*int64(wc))
}
