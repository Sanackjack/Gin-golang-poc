package helper

func ErrorPanic(err error) {
	if err != nil {

		//errors.Is(er)
		panic(err)
	}
}
