package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ
	scanner.Scan() // １行分の入力を取得する
	fmt.Println(scanner.Text())
}
