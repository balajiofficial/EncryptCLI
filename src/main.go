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

func reverse(str string) string {
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return reversed
}

func main() {
	fmt.Print("Enter 1 for encryption or 2 for decryption : ")
	var num int
	fmt.Scanln(&num)
	if num == 1 {
		fmt.Println("Enter the path of the file to be encrypted -")
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
				ascii_str := fmt.Sprint(ascii_value * 17)
				encrypted += fmt.Sprint(len(ascii_str)) + reverse(ascii_str)
			}

			//  Encrypt Main Document
			for i := 0; i < len(file); i++ {
				ascii_value := int(key[i%15])
				encrypted_str := fmt.Sprint(ascii_value * int(file[i]))
				encrypted += fmt.Sprint(len(encrypted_str)) + reverse(encrypted_str)
			}

			os.WriteFile(fileName, []byte(encrypted), 0644)
			fmt.Println("File has been encrypted")
		}
	} else if num == 2 {
		fmt.Println("Enter the path of the file to be decrypted -")
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
			for i = 0; j < 15; j++ {
				ind, err := strconv.Atoi(string(rune(file[i])))
				if err != nil {
					fmt.Println("Invalid Decryption Format")
					os.Exit(0)
				}
				temp = ""
				for ind > 0 {
					i++
					ind--
					temp += string(rune(file[i]))
				}
				n, _ := strconv.Atoi(reverse(temp))
				key += string(rune((n / 17)))
				i++
			}

			// Decrypt Main Document
			text, temp := "", ""
			j = 0
			for i < len(file) {
				ind, _ := strconv.Atoi(string(rune(file[i])))
				for ; ind > 0; ind-- {
					i++
					temp += string(rune(file[i]))
				}
				n, err := strconv.Atoi(reverse(temp))
				if err != nil {
					fmt.Println("Invalid Decryption Format")
					os.Exit(0)
				}
				key_value := int(key[j])
				text += string(rune(n / key_value))
				j = (j + 1) % 15
				temp = ""
				i++
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
