package exception

// PanicIf panic if error exists
func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
