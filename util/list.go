package util

func RemoveDuplicatesAndEmpty(list []string) (ret []string) {
	length := len(list)
	for i := 0; i < length; i++ {
		if (i > 0 && list[i-1] == list[i]) || len(list[i]) == 0 {
			continue
		}
		ret = append(ret, list[i])
	}
	return
}
