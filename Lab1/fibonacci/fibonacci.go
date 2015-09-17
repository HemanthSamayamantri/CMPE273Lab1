//package main
package fibpackage

import "fmt"

/* this is a comment
func main() {
	fmt.Println("Fibonacci sequence sum for the number:")
	var input int64
	fmt.Scanf("%d", &input)
	fmt.Println("The result is", ComputeFibSum(input))
}
*/

func ComputeFibSum(value int64) int64 {
	if value == 0 {
		return 0
	} else if value == 1 {
		return 1
	} else {
		return ComputeFibSum(value-1) + ComputeFibSum(value-2)
	}
}
