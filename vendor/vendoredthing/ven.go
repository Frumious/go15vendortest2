package vendoredthing

import (
	"fmt"
)

type Data struct {
	Msg         string
	Count       int
	ExtraString string
}

func DoSomething() {
	fmt.Println("I am a vendored something")
}

func GetData() Data {
	return Data{"hello from ventest2", 2, "extra string"}
}
