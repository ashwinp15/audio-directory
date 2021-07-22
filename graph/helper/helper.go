package helper

func MaybeNullString(s string) interface{} {
	if len(s) == 0 {
		return
	}

}
