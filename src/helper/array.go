package helper

func ArrayMerge(args ...[]interface{}) []interface{} {
	count := len(args)
	if count == 0 {
		return nil
	}

	totalCount := 0
	for _, value := range args {
		totalCount += len(value)
	}

	res := make([]interface{}, totalCount)

	mergeCount := 0
	for _, value := range args {
		count = len(value)
		copy(res[mergeCount:], value)
		mergeCount += count
	}

	return res
}

func ArgsToInterfaceArray(args ...interface{}) []interface{} {
	count := len(args)
	if count == 0 {
		return nil
	}

	res := make([]interface{}, count)

	for index, value := range args {
		res[index] = value
	}

	return res
}
