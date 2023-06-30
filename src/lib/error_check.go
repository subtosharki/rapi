package lib

func ErrorCheck(err error) {
	if err != nil {
		Error(err.Error())
	}
}
