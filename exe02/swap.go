package main
import "fmt"


func swap0 (x,y int)(int, int){
	return y,x
}

func swap1 (x,y int){
	x,y = y,x
}

func swap2(x *int,y *int){
	*x, *y = *y, *x
}
func swap3(x **int, y **int){
	**x , **y = **y, **x
}



func main(){
	var a,b = 1,2
	fmt.Printf("Initial: a=%d, b=%d\n", a,b)
	a,b = b,a
	fmt.Printf("after a,b = b,a: a=%d, b=%d\n", a,b)

	a,b = swap0(a,b)
	fmt.Printf("after swap0: a=%d, b=%d\n", a,b)
	swap1(a,b)
	fmt.Printf("after swap1: a=%d, b=%d\n", a,b)
	swap2(&a,&b)
	fmt.Printf("after swap2: a=%d, b=%d\n", a,b)
	pa,pb := &a,&b
	swap3(&pa,&pb)
	fmt.Printf("after swap3: a=%d, b=%d\n", a,b)

}