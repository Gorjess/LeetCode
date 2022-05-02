/*Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature. If there is no future day for which this is possible, put 0 instead.

For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].

Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/daily-temperatures
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateTemps(length int) []int {
	out := make([]int, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		out[i] = rand.Intn(100)
	}
	return out
}

func dailyTemperatures(T []int) []int {
	type elem struct {
		idx   int
		value int
	}

	tCnt := len(T)

	// result
	var (
		result = make([]int, tCnt)
	)

	// stack
	var (
		stack       = make([]*elem, tCnt)
		stackCursor = 0
	)

	// stack ops
	push := func(i, e int) {
		stack[stackCursor] = &elem{idx: i, value: e}
		stackCursor++
	}
	pop := func() {
		stack = stack[:len(stack)-1]
		stackCursor--
	}
	empty := func() bool {
		return stackCursor == 0
	}
	tail := func() *elem {
		if stackCursor == 0 {
			return nil
		}
		return stack[stackCursor-1]
	}

	for i := 0; i < tCnt; i++ {
		if empty() || T[i] <= tail().value { // carry on
			push(i, T[i])
		} else { // record result
			te := tail()
			for !empty() && T[i] > te.value {
				result[te.idx] = i - te.idx
				pop()
				te = tail()
			}
			push(i, T[i])
		}
	}

	return result
}

func main() {
	fmt.Println(dailyTemperatures(generateTemps(5)))
}
