package main

import (
	"fmt"
	//	"hunterkim.com/kim/stringutil"
)

type T struct {
	v1 int
	v2 []int
}

/*func (o *T) fn(a int, b int, c chan T, exit chan bool) {
	go func() {
		oo := <-c
		fmt.Println(oo.v1, oo.v2, a, b)
		exit <- true
	}()
}*/
func main() {
	//str := stringutil.Reverse("ABC中")
	//length := len("1234567890中；")
	//fmt.Println(str, length)

	//v := []rune("abd中文efg")
	//for i,s := range v{
	//	fmt.Println(i,string(s))
	//}
	str := "abc"
	s := []rune(str)
	s[1] = 0xFFFFD
	for _, c := range s {
		fmt.Printf("%#U\n", c)
	}
	fmt.Println()
	fmt.Println(str)

	for i := 0; i < 3; i++ {

		v := 10
		switch v {
		case 10:
			fmt.Println("10")
			break
		case 11:
			fmt.Println("11")
		}
	}
	//c := make(chan T, 1)
	//exit := make(chan bool)
	//c <- o
	//o.fn(2, 4, c, exit)
	//<-exit
}
