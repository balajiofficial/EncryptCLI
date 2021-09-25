package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func generateRandomKey() string {
	var key = ""
	for i := 0; i < 10; i++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		time.Sleep(time.Duration(1))
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

func decimalTobase77(n int) string {
	var str = ""
	r := n % 77
	for n > 0 {
		c := ""
		if r < 10 {
			c = fmt.Sprint(r)
		} else if r < 36 {
			c = string(rune(r - 10 + int('A')))
		} else if r < 62 {
			c = string(rune(r - 36 + int('a')))
		} else {
			c = string(rune(r - 62 + int('!')))
		}
		str += fmt.Sprint(c)
		n /= 77
		r = n % 77
	}
	return reverse(str)
}

func base77ToDecimal(str string) int {
	var n = 0
	str = reverse(str)
	for i := 0; i < len(str); i++ {
		if int(str[i]) >= int('a') {
			n += (int(str[i]) - int('a') + 36) * int(math.Pow(77, float64(i)))
		} else if int(str[i]) >= int('A') {
			n += (int(str[i]) - int('A') + 10) * int(math.Pow(77, float64(i)))
		} else if int(str[i]) >= int('0') {
			n += (int(str[i]) - int('0')) * int(math.Pow(77, float64(i)))
		} else {
			n += (int(str[i]) - int('!') + 62) * int(math.Pow(77, float64(i)))
		}
	}
	return n
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
			fmt.Print("Do you wish to have your own password?")
			scanner.Scan()
			response := scanner.Text()
			var key string
			if response == "y" {
				fmt.Println("Enter your new password : ")
				scanner.Scan()
				key = scanner.Text()
				fmt.Println(key)
			} else {
				key = generateRandomKey()
			}
			for i := 0; i < 10; i++ {
				ascii_value := int(key[i])
				ascii_str := fmt.Sprint(decimalTobase77(ascii_value * 17))
				encrypted += fmt.Sprint(len(ascii_str)) + reverse(ascii_str)
			}

			//  Encrypt Main Document
			for i := 0; i < len(file); i++ {
				ascii_value := int(key[i%10])
				encrypted_str := fmt.Sprint(decimalTobase77(ascii_value * int(file[i])))
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
			for i = 0; j < 10; j++ {
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
				n := base77ToDecimal(reverse(temp))
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
				n := base77ToDecimal(reverse(temp))
				key_value := int(key[j])
				text += string(rune(n / key_value))
				j = (j + 1) % 10
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
