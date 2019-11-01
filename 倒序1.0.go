package main

import (
	"fmt"
)

func main() {
	var hel string
	hel ="hello"
	fmt.Scanf("%s", &hel)
	res := reverse(hel)
	fmt.Println(res)
}

func reverse(hel string) (res string) {
	for i := len(hel)-1 ; i >= 0 ; i-- {
		res+=string(hel[i])
	}
	return
}

