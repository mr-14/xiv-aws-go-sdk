package errorutil

// PanicIfError panic if error is not nil
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
