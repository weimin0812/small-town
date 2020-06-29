package en

import (
	"strconv"
)

//Write a program that outputs the string representation of numbers from 1 to n.
// 
//
// But for multiples of three it should output “Fizz” instead of the number and 
//for the multiples of five output “Buzz”. For numbers which are multiples of both
// three and five output “FizzBuzz”. 
//
// Example:
// 
//n = 15,
//
//Return:
//[
//    "1",
//    "2",
//    "Fizz",
//    "4",
//    "Buzz",
//    "Fizz",
//    "7",
//    "8",
//    "Fizz",
//    "Buzz",
//    "11",
//    "Fizz",
//    "13",
//    "14",
//    "FizzBuzz"
//]
// 
//
//leetcode submit region begin(Prohibit modification and deletion)
func fizzBuzz(n int) []string {
	var ret []string
	fb := []struct {
		Key   int
		Value string
	}{
		{Key: 3, Value: "Fizz"},
		{Key: 5, Value: "Buzz"},
	}
	for i := 1; i <= n; i++ {
		// todo 使用map的方式是怎么保证顺序的
		ret = append(ret, convert(i, fb))
	}
	return ret
}

func convert(i int, fb []struct {
	Key   int
	Value string
}) string {
	ret := ""
	for _, t := range fb {
		if i%t.Key == 0 {
			ret += t.Value
		}
	}
	if ret == "" {
		ret = strconv.Itoa(i)
	}
	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
