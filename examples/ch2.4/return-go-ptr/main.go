package main

/*
//extern int* getGoPtr();
//
//static void Main() {
//	int * p =  getGoPtr();
//	*p = 42;

*/
//#include

/*
extern int* getGoPtr();

static void Main() {
	int* p = getGoPtr();
	*p = 42;
}
*/
import "C"

func main() {
	C.Main()
}

/*func main() {
	C.Main()
}
*/
//export getGoPtr
func getGoPtr() *C.int {
	return new(C.int)
}
