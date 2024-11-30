package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1:単語入力")
		fmt.Println("2:辞書登録リスト")
		fmt.Println("3:終了")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		fmt.Println("単語入力")
	case "2":
		fmt.Println("選択した辞書リストです")
	case "3":
		fmt.Println("終了します")
		return
	default:
		fmt.Println("無効な数字が入力されました")
	}
  }
}
