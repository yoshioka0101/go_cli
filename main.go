package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data := make(map[string]string)

	for {
		fmt.Println("1: データを登録")
		fmt.Println("2: 保存されたデータを表示")
		fmt.Println("3: キーを検索して値を取得")
		fmt.Println("4: 終了")
		fmt.Print("選択してください: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("キーを入力してください: ")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)

			fmt.Print("値を入力してください: ")
			value, _ := reader.ReadString('\n')
			value = strings.TrimSpace(value)

			data[key] = value
			fmt.Printf("登録されました: %s - %s\n", key, value)
		case "2":
			if len(data) == 0 {
				fmt.Println("保存されたデータはありません")
			} else {
				fmt.Println("保存されたデータ:")
				for key, value := range data {
					fmt.Printf("%s: %s\n", key, value)
				}
			}
		case "3":
			fmt.Print("検索するキーを入力してください: ")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)

			if value, exists := data[key]; exists {
				fmt.Printf("キー「%s」の値は「%s」です\n", key, value)
			} else {
				fmt.Printf("キー「%s」は登録されていません\n", key)
			}
		case "4":
			fmt.Println("終了します")
			return
		default:
			fmt.Println("無効な選択です。もう一度試してください")
		}
	}
}
