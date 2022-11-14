package main

//int sum(int a, int b);
import "C"

func sum(a, b C.int) C.int {
	return a + b
}

func main() {

}
