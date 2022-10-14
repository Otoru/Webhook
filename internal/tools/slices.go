package tools

func Contains(value any, slice []any) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}

	return false
}
