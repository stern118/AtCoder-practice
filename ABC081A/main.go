package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	marbleCount := 0 // ビー玉の合計数

	// 読み込み
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	// ビー玉の数を数える
	for _, v := range strings.Split(line, "") {
		num, _ := strconv.Atoi(v)
		if num == 1 {
			marbleCount++
		}
	}

	// 出力
	fmt.Println(marbleCount)
}
