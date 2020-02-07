package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// 読み込み
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")

	// 積の合計を計算
	var nums [2]int
	for i, v := range line {
		num, _ := strconv.Atoi(v)
		nums[i] = num
	}
	total := nums[0] * nums[1]

	// 判定
	if total%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
