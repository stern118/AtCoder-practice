package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0 // 合計値
	scanner := bufio.NewScanner(os.Stdin)

	// 1行目を読み込む
	scanner.Scan()
	num1, _ := strconv.Atoi(scanner.Text())
	sum += num1

	// 2行目を読み込む
	scanner.Scan()
	line2 := strings.Split(scanner.Text(), " ")
	for _, v := range line2 {
		num, _ := strconv.Atoi(v)
		sum += num
	}

	// 3行目を読み込む
	scanner.Scan()

	// 結果出力
	fmt.Println(sum, scanner.Text())
}
