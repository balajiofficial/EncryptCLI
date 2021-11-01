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

func generateRandomKey(n int) string {
	var key = ""
	for i := 0; i < n; i++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		time.Sleep(time.Duration(1))
		key += string(rune(rand.Intn(94) + 33))
	}
	return key
}

func generateRandomEvenChar() rune {
	var char rune = 'A'
	for int(char)%2 == 1 {
		rand.Seed(int64(time.Now().Nanosecond()))
		time.Sleep(time.Duration(1))
		char = rune(rand.Intn(94) + 33)
	}
	return char
}

func generateRandomOddChar() rune {
	var char rune = 'B'
	for int(char)%2 == 0 {
		rand.Seed(int64(time.Now().Nanosecond()))
		time.Sleep(time.Duration(1))
		char = rune(rand.Intn(94) + 33)
	}
	return char
}

func reverse(str string) string {
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return reversed
}

func decimalTobase95(n int) string {
	var str = ""
	r := n % 95
	for n > 0 {
		c := ""
		if r < 10 {
			c = fmt.Sprint(r)
		} else if r < 36 {
			c = string(rune(r - 10 + int('A')))
		} else if r < 62 {
			c = string(rune(r - 36 + int('a')))
		} else if r < 78 {
			c = string(rune(r - 62 + int(' ')))
		} else if r < 85 {
			c = string(rune(r - 78 + int(':')))
		} else if r < 91 {
			c = string(rune(r - 85 + int('[')))
		} else {
			c = string(rune(r - 91 + int('{')))
		}
		str += fmt.Sprint(c)
		n /= 95
		r = n % 95
	}
	return reverse(str)
}

func base95ToDecimal(str string) int {
	var n = 0
	str = reverse(str)
	for i := 0; i < len(str); i++ {
		if int(str[i]) >= int('{') {
			n += (int(str[i]) - int('{') + 91) * int(math.Pow(95, float64(i)))
		} else if int(str[i]) >= int('a') {
			n += (int(str[i]) - int('a') + 36) * int(math.Pow(95, float64(i)))
		} else if int(str[i]) >= int('[') {
			n += (int(str[i]) - int('[') + 85) * int(math.Pow(95, float64(i)))
		} else if int(str[i]) >= int('A') {
			n += (int(str[i]) - int('A') + 10) * int(math.Pow(95, float64(i)))
		} else if int(str[i]) >= int(':') {
			n += (int(str[i]) - int(':') + 78) * int(math.Pow(95, float64(i)))
		} else if int(str[i]) >= int('0') {
			n += (int(str[i]) - int('0')) * int(math.Pow(95, float64(i)))
		} else {
			n += (int(str[i]) - int(' ') + 62) * int(math.Pow(95, float64(i)))
		}
	}
	return n
}

func exitMessage(message string) {
	fmt.Println()
	fmt.Println(message)
	fmt.Print("Press enter to exit the program...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	os.Exit(0)
}

func stringToNumber(str string) int {
	num := 0
	if str == "1" {
		num = 5
	} else if str == "2" {
		num = 15
	} else {
		exitMessage("Invalid Input")
	}
	return num
}

func encryptText(key string, saveSpace bool, file []byte) string {
	text := decimalTobase95(len(key))
	test := ""
	for i := 0; i < len(key); i++ {
		ascii_value := int(key[i])
		ascii_str := ""
		if saveSpace {
			ascii_str = decimalTobase95(ascii_value)
		} else {
			ascii_str = decimalTobase95(ascii_value * 17)
		}
		test += fmt.Sprint(len(ascii_str)) + reverse(ascii_str)
		text += fmt.Sprint(len(ascii_str)) + reverse(ascii_str)
	}

	//  Encrypt Main Document
	for i := 0; i < len(file); i++ {
		ascii_value := int(key[i%len(key)])
		encrypted_str := fmt.Sprint(decimalTobase95(ascii_value * int(file[i])))
		text += fmt.Sprint(len(encrypted_str)) + reverse(encrypted_str)
	}
	return text
}

func decryptText(file []byte, i int, key string) string {
	text, temp := "", ""
	j := 0
	for i < len(file) {
		ind, _ := strconv.Atoi(string(rune(file[i])))
		for ; ind > 0; ind-- {
			i++
			temp += string(rune(file[i]))
		}
		n := base95ToDecimal(reverse(temp))
		key_value := int(key[j])
		text += string(rune(n / key_value))
		j = (j + 1) % len(key)
		temp = ""
		i++
	}
	return text
}

func main() {
	fmt.Print("Enter 1 for encryption or 2 for decryption : ")
	var num int
	fmt.Scanln(&num)
	if num == 1 {

		// Encryption
		password := false
		saveSpace := false
		fmt.Println("Enter the path of the file to be encrypted - ")
		var fileName string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fileName = scanner.Text()
		file, err := os.ReadFile(fileName)
		if err != nil {
			exitMessage("File not found")
		}

		fmt.Print("Do you wish to save space or increase security? (1/2) : ")
		scanner.Scan()
		securityRes := scanner.Text()
		var spaceOrSecurity = stringToNumber(securityRes)
		if securityRes == "1" {
			saveSpace = true
		} else {
			saveSpace = false
		}

		// Insert Key
		var encrypted = ""
		if saveSpace {
			encrypted += string(generateRandomOddChar())
		} else {
			encrypted += string(generateRandomEvenChar())
		}
		fmt.Print("Do you wish to have your own password? (y/n) : ")
		scanner.Scan()
		response := scanner.Text()
		var key string
		if response == "y" {
			encrypted += string(generateRandomEvenChar())
			for true {
				fmt.Print("Enter your password : ")
				scanner.Scan()
				key = scanner.Text()
				if key == "" {
					fmt.Println("Invalid Input")
					continue
				}
				if len(key) > 94 {
					fmt.Println("Password Length should be less than 95")
					continue
				}
				password = true
				break
			}
		} else if response == "n" {
			encrypted += string(generateRandomOddChar())
			key = generateRandomKey(spaceOrSecurity)
		} else {
			exitMessage("Invalid Input")
		}

		encrypted += encryptText(key, saveSpace, file)

		os.WriteFile(fileName, []byte(encrypted), 0644)
		if !password {
			exitMessage("File has been encrypted")
		} else {
			exitMessage("File has been encrypted using password")
		}
	} else if num == 2 {

		// Decryption

		fmt.Println("Enter the path of the file to be decrypted - ")
		var fileName string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fileName = scanner.Text()
		file, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("File not found")
		}

		saveSpace := int(file[0])%2 == 1
		password := int(file[1])%2 == 0

		var key, temp, passwordStr = "", "", ""
		var j, i, passwordLen = 0, 2, 10

		passwordLen = base95ToDecimal(string(file[i]))
		i++
		for ; j < passwordLen; j++ {
			ind, _ := strconv.Atoi(string(rune(file[i])))
			if err != nil {
				exitMessage("Invalid Decryption Format")
			}
			temp = ""
			for ind > 0 {
				i++
				ind--
				temp += string(rune(file[i]))
			}
			n := base95ToDecimal(reverse(temp))
			if saveSpace {
				key += string(rune((n)))
			} else {
				key += string(rune((n / 17)))
			}
			i++
		}

		if password {
			fmt.Print("Enter password : ")
			scanner.Scan()
			passwordStr = scanner.Text()
		} else {
			passwordStr = key
		}

		if passwordStr != key {
			exitMessage("Incorrect Password")
		}

		// Decrypt Main Document
		text := decryptText(file, i, key)

		os.WriteFile(fileName, []byte(text), 0644)
		if password {
			exitMessage("File has been decrypted using password")
		} else {
			exitMessage("File has been decrypted")
		}

	} else {
		exitMessage("Invalid input")
	}
}
