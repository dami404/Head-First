// Функции
package Chapter3

import (
	"fmt"
)

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}
func task1() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}

func task2() {
	j := 1
	jPointer := &j
	k := &jPointer
	fmt.Println(jPointer)
	fmt.Println(*jPointer)
	fmt.Println(k)
	fmt.Println(*k)
	fmt.Println(**k)

}

// doc example 2
func Task3() {
	var myInt int = 42
	myIntPointer := &myInt
	fmt.Println(*myIntPointer)
}

// func divide(dividend float64, divisor float64) (float64, error) {
// 	if divisor == 0.0 {
// 		return 0, errors.New("can't divide by 0")
// 	}
// 	return dividend / divisor, nil
// }

// func main() {
// 	quotient, err := divide(5.6, 0.0)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Printf("%0.2f\n", quotient)
// 	}
// }

// doc example 3
func Main() {
	task1()
	task2()
	Task3()
}
