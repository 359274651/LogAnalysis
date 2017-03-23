package CommonLibrary

//import "log"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
