package helper

func SetProperty(property *string, status int) {
	if status >= 400 {
		*property = "error"
	} else {
		*property = "message"
	}
}