package utils

func Panic(err interface{}) {
	if err != nil {
		panic(err)
	}
}
