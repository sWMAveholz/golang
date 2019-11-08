package main

import "testing"
import "fmt"

func Test_MakeIndex(t *testing.T){
	p1 := Page{"Bla1","Bla2","Bla3"}
	p2 := Page{"Bla1","Bla2"}

	b := Book{p1,p2}

	i := MakeIndex(b)

	fmt.Println(b)
	fmt.Println(i)


}