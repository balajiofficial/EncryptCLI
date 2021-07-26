package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func generateRandomKey() string {
	var key = ""
	for i := 0; i < 15; i++ {
		key += string(rune(rand.Intn(94) + 33))
	}
	return key
}

func main() {
	fmt.Print("Enter 1 for encryption or 2 for decryption : ")
	var num int
	fmt.Scanln(&num)
	if num == 1 {
		fmt.Print("Enter the path of the file to be encrypted : ")
		var fileName string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fileName = scanner.Text()
		file, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("File not found")
		} else {

			// Insert Key
			var encrypted = ""
			key := generateRandomKey()
			for i := 0; i < 15; i++ {
				ascii_value := int(key[i])
				encrypted += fmt.Sprintf("%d ", ascii_value*17)
			}

			//  Encrypt Main Document
			for i := 0; i < len(file); i++ {
				ascii_value := int(key[i%15])
				encrypted += fmt.Sprintf("%d ", ascii_value*int(file[i]))
			}
			os.WriteFile(fileName, []byte(encrypted), 0644)
			fmt.Println("File has been encrypted")
		}
	} else if num == 2 {
		fmt.Print("Enter the path of the file to decrypted : ")
		var fileName string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fileName = scanner.Text()
		file, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("File not found")
		} else {

			// Extract key
			var key, temp = "", ""
			var j, i = 0, 0
			for i = 0; j < 15; i++ {
				if string(file[i]) == " " {
					n, _ := strconv.Atoi(temp)
					key += string(rune((n / 17)))
					temp = ""
					j++
				} else {
					temp += string(file[i])
				}
			}

			// Decrypt Main Document
			text := ""
			j = 0
			for ; i < len(file); i++ {
				if string(file[i]) == " " {
					n, _ := strconv.Atoi(temp)
					text += string(rune(n / int(key[j])))
					j++
					j = j % 15
					temp = ""
				} else {
					temp += string(file[i])
				}
			}

			os.WriteFile(fileName, []byte(text), 0644)
			fmt.Println("File has been decrypted")
		}
	} else {
		fmt.Println("Invalid input")
	}
	fmt.Println()
	fmt.Print("Press enter to exit the program...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}
