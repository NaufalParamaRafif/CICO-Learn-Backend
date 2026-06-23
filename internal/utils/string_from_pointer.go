package utils

func StringFromPtr(str *string) string {
	if str == nil {
		return ""
	}

	return *str

}
