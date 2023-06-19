package lib

func RapiErrorCheck(err error) {
	if err != nil {
		RapiError(err.Error())
	}
}
