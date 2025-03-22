package helper

// Function PanicIfError
// Use to handle panic if error is not nil
// @Parameter, err error
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
