package utils


func RemoveStringFromSlice(slc []string, target string) []string {
	var index = -1

	for i, str := range slc {
		if str == target {
			index = i
			break
		}
	}

	if index == -1 { // not found
		return slc
	}

	slc[index] = slc[len(slc)-1]
	slc = slc[:len(slc)-1]

	return slc
}
