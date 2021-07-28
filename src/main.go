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

func encode_zeroes(str string) string {
	encoded_str := str
	for i := 0; i < 5-len(str); i++ {
		encoded_str = "0" + encoded_str
	}
	return encoded_str
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
				encrypted += reverse(encode_zeroes(fmt.Sprintf("%d", ascii_value*17)))
			}

			//  Encrypt Main Document
			for i := 0; i < len(file); i++ {
				ascii_value := int(key[i%15])
				encrypted += reverse(encode_zeroes(fmt.Sprintf("%d", ascii_value*int(file[i]))))
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
			for i = 0; j < 15; i++ {
				temp += string(file[i])
				if (i+1)%5 == 0 {
					n, _ := strconv.Atoi(reverse(temp))
					key += string(rune((n / 17)))
					temp = ""
					j++
				}
			}

			// Decrypt Main Document
			text := ""
			j = 0
			k := 0
			for ; i < len(file); i++ {
				k++
				temp += string(file[i])
				if k%5 == 0 {
					n, _ := strconv.Atoi(reverse(temp))
					text += string(rune(n / int(key[j])))
					j = (j + 1) % 15
					temp = ""
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
