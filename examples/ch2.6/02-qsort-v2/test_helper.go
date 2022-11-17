package qsort

//extern int t_get_go_qsort_compare(void* a, void* b)
import "C"

import "unsafe"

func t_get_go_qsort_compare() CompareFunc {
	return CompareFunt(c.t_get_go_qsort_compare)
}

//export t_go_qsort_compare
func t_go_qsort_compare(a, b unsafe.Pointer) C.int {
	pa, pb := (*C.int)(a), (*C.int)(b)
	return C.int(*pa - *pb)
}
