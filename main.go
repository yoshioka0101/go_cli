package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const dataFile = "data.json" 

func main() {
	reader := bufio.NewReader(os.Stdin)
	data := loadData()

	for {
		showMenu()
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			registerData(reader, data)
		case "2":
			showData(data)
		case "3":
			searchData(reader, data)
		case "4":
			saveData(data)
			fmt.Println("終了します")
			return
		default:
			fmt.Println("無効な選択です。もう一度試してください")
		}
	}
}

func showMenu() {
	fmt.Println("\n1: データを登録")
	fmt.Println("2: 保存されたデータを表示")
	fmt.Println("3: キーを検索して値を取得")
	fmt.Println("4: 終了")
	fmt.Print("選択してください: ")
}

func registerData(reader *bufio.Reader, data map[string]string) {
	key, _ := readInput(reader, "キーを入力してください: ")
	value, _ := readInput(reader, "値を入力してください: ")
	data[key] = value
	fmt.Printf("登録されました: %s - %s\n", key, value)
}

func showData(data map[string]string) {
	if len(data) == 0 {
		fmt.Println("保存されたデータはありません")
		return
	}
	fmt.Println("保存されたデータ:")
	for key, value := range data {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func searchData(reader *bufio.Reader, data map[string]string) {
	key, _ := readInput(reader, "検索するキーを入力してください: ")
	if value, exists := data[key]; exists {
		fmt.Printf("キー「%s」の値は「%s」です\n", key, value)
	} else {
		fmt.Printf("キー「%s」は登録されていません\n", key)
	}
}

func readInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func loadData() map[string]string {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			// ファイルが存在しない場合は空のマップを返す
			return make(map[string]string)
		}
		fmt.Println("データの読み込み中にエラーが発生しました:", err)
		os.Exit(1)
	}
	defer file.Close()

	data := make(map[string]string)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("データのデコード中にエラーが発生しました:", err)
		os.Exit(1)
	}

	return data
}

func saveData(data map[string]string) {
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("データの保存中にエラーが発生しました:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		fmt.Println("データのエンコード中にエラーが発生しました:", err)
	}
}