package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"math"
	"strconv"
	"sync/atomic"
	"time"
)

type Config struct {
	ArraySize		int
	MaxGoroutines	int
	Port			int
}
var config Config

var activeGoRoutines int32 = 0

// Cerinta individuala 1. 
// Intoarce un array de cuvinte unde cuvantul i este format din caracterele de pe pozitia i din fiecare cuvant din sirul initial.
// [mama,tata] => [mt, aa, mt, aa]
func processRequest1(data []string) []string {
	result := make([]string, len(data[0]))
    
	for i := range data[0] {
		for _, word := range data {
			if i < len(word) {
				result[i] = result[i] + string(word[i])
			}
		}
	}
	
	return result
}

// Cerinta individuala 2.
// Intoarce numarul de patrate perfecte gasite in string-urile din sir
// [1o6na,9,io2n4] => 2 patrate perfecte
func processRequest2(data []string) int {
	count := 0

	for _, word := range data {
		num_string := ""

		for _, char := range word {
			if char >= '0' && char <= '9' {
				num_string += string(char)
			}
		}

		if num_string != "" {
			num, _ := strconv.Atoi(num_string)
			sqrt := int(math.Sqrt(float64(num)))
			if sqrt * sqrt == num {
				count++
			}
		}
	}

	return count
}

// Cerinta individuala 3.
// Intoarce suma numerelor inversate din array.
// [11,12,13] => 11 + 21 + 31 = 63
func processRequest3(data []string) int {
	sum := 0

	for _, num_string := range data {
		reversed_num_string := ""
		for i := len(num_string) - 1; i >= 0; i-- {
			reversed_num_string = reversed_num_string + string(num_string[i])
		}

		reversed_num, _ := strconv.Atoi(reversed_num_string)
		sum = sum + reversed_num
	}

	return sum
}

// Cerinta individuala 4.
// Intoarce media aritmetica a numerelor citite, pentru care suma cifrelor aparține intervalului [a,b].
// a = 2, b = 10, n = 5 și [11,39,32,80,84] => 41
func sumOfDigits(n int) int {
	sum := 0
	for n != 0 {
		sum = sum + n % 10
		n = n / 10
	}
	return sum
}

func processRequest4(data []string) float64 {
	if len(data) < 3 {
		return 0
	}

	a, _ := strconv.Atoi(data[0])
	b, _ := strconv.Atoi(data[1])
	n, _ := strconv.Atoi(data[2])

	total := 0
	count := 0
	i := 3

	for n > 0 {
		num, _ := strconv.Atoi(data[i])
		digit_sum := sumOfDigits(num)

		if digit_sum >= a && digit_sum <= b {
			total = total + num
			count++
		}
		
		i++
		n--
	}

	if count == 0 {
		return 0
	}
	result := float64(total) / float64(count)
	return result
}

// Cerinta individuala 5.
// Intoarce conversia in baza 10 a numerelor binare din array-ul de string-uri dat.
// [2dasdas,12,dasdas,1010,101] => 10, 3 
func isBinary(s string) bool {
	for _, char := range s {
		if char != '0' && char != '1' {
			return false
		}
	}

	return true
}

func binaryToDecimal(binary_string string) int {
	result := 0
	for _, char := range binary_string {
		if char == '0' {
			result = result * 2
		} else if char == '1' {
			result = result * 2 + 1
		}
	}
	return result
}

func processRequest5(data []string) []int {
	result := []int{}

	for _, str := range data {
		if isBinary(str) {
			decimal := binaryToDecimal(str)
			result = append(result, decimal)
		}
	}

	return result
}

// Cerinta individuala 6.
// Intoarce un nou string format prin shiftarea fiecarui caracter din string-ul initial la DREAPTA sau la STANGA cu k pozitii in alfabet.
// [3,LEFT,abcdef] => xyzabc
func processRequest6(data []string) []string {
    k, _ := strconv.Atoi(data[0])
    direction := data[1] // "LEFT" or "RIGHT"
    strings := data[2:]

    shift := k % 26
    if direction == "LEFT" {
        shift = -shift
    }

    result := make([]string, len(strings))

    for i, str := range strings {
        var transformed_string string
        for _, char := range str {
            if char >= 'a' && char <= 'z' {
                new_char := ((int(char-'a') + shift + 26) % 26) + 'a'
                transformed_string += string(new_char)
            } else if char >= 'A' && char <= 'Z' {
                new_char := ((int(char-'A') + shift + 26) % 26) + 'A'
                transformed_string += string(new_char)
            } else {
                transformed_string += string(char)
            }
        }
        result[i] = transformed_string
    }

    return result
}

// Cerinta individuala 7.
// Intoarce un string codificat, dupa urmatoarea regula: in fata fiecarui caracter este scris un nr ce reorezinta numarul de aparitii consecutive al acestuia.  
// [1G11o1L] => [GoooooooooooL]
func processRequest7(data []string) []string {
	result := []string{}
    for _, encoded_text := range data {
        i := 0
		decoded_text := ""
        for i < len(encoded_text) {
            count := 0
            for i < len(encoded_text) && encoded_text[i] >= '0' && encoded_text[i] <= '9' {
                count = count*10 + int(encoded_text[i]-'0')
                i++
            }

            if i < len(encoded_text) {
                char := encoded_text[i]
                decoded_text += strings.Repeat(string(char), count)
                i++
            }
        }
		result = append(result, decoded_text)
    }

    return result
}

// Cerinta individuala 8.
// Intoarce numarul total de cifre al numerelor prime din array
// [23,17,15,3,18] => 5 cifre (nr prime 23, 17, 3) 
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for d := 2; d * d <= num; d++ {
		if num % d == 0 {
			return false
		}
	}
	return true
}

func processRequest8(data []string) int {
	total_digits := 0

	for _, num_string := range data {
		num, _ := strconv.Atoi(num_string)

		if isPrime(num) {
			digits := 0
			copy_num := num
			for copy_num != 0 {
				digits++
				copy_num = copy_num / 10
			}
			total_digits = total_digits + digits
		}
	}

	return total_digits
}

// Cerinta individuala 9.
// Intoarce numerul de cuvinte care au un numar par de vocale aflate pe pozitii pare în cuvânt.  
// [mama,iris,bunica,ala] => 2 cuvinte [iris, ala] 
func processRequest9(data []string) int {
	result := 0

	for _, word := range data {
		vowel_count := 0
		vowel_positions := []int{}
		
		for i, char := range word {
			if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || 
			   char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
				vowel_count++
				vowel_positions = append(vowel_positions, i)
			}
		}
		
		if vowel_count % 2 == 0 {
			all_vowels_on_even_pos := true
			for _, pos := range vowel_positions {
				if pos % 2 != 0 {
					all_vowels_on_even_pos = false
					break
				}
			}
			
			if all_vowels_on_even_pos {
				result++
			}
		}
	}

	return result
}


// Cerinta individuala 10.
// Intoarce cmmdc pentru toate numere (se calculeaza listele divizorilor pentru fiecare numar si se calculeaza intersectia listelor).
// [24,16,8,aaa,bbb] => 8 
func extractNumbersFromString(s string) int {
	num_string := ""
	for _, char := range s {
		if char >= '0' && char <= '9' {
			num_string = num_string + string(char)
		}
	}

	if num_string != "" {
		num, _ := strconv.Atoi(num_string)
		return num
	}
	return -1
}

func getDivisors(num int) []int {
	divisors := []int{}
	for d := 1; d <= num; d++ {
		if num % d == 0 {
			divisors = append(divisors, d)
		}
	}
	return divisors
}

func intersectLists(list1, list2 []int) []int {
	intersection := []int{}
	for _, val1 := range list1 {
		for _, val2 := range list2 {
			if val1 == val2 {
				intersection = append(intersection, val1)
				break
			}
		}
	}
	return intersection
}

func processRequest10(data []string) int {
	divisors_list := [][]int{}

	for _, str := range data {
		num := extractNumbersFromString(str)
		if num != -1 {
			divisors := getDivisors(num)
			divisors_list = append(divisors_list, divisors)
		}
	}

	if len(divisors_list) == 0 {
		return 0
	}

	result := divisors_list[0]
	for _, divisors := range divisors_list[1:] {
		result = intersectLists(result, divisors)
	}

	if len(result) > 0 {
		return result[len(result)-1]
	}

	return 0
}

// Cerinta individuala 11.
// Intoarce suma nr din array formatate prin permutarea la dreapta a k cifre ale fiecaruia
// [2,1234,3456,4567] => 3412 + 5634 + 6745 = 15791 (k = 2)
func rotateRight(num_string string, k int) int {
	n := len(num_string)

	k = k % n

	rotated_string := num_string[n-k:] + num_string[:n-k]

	rotated_num, _ := strconv.Atoi(rotated_string)
	return rotated_num
}

func processRequest11(data []string) int {
	sum := 0
	k, _ := strconv.Atoi(data[0])

	for _, num := range data[1:] {
		rotated_num := rotateRight(num, k)
		sum = sum + rotated_num
	}

	return sum
}

// Cerinta individuala 12.
// Intoarce suma elementelor array-ului format din dublarea primei cifre a fiecarui nr din array-ul initial
// [23,43,26,74] => 223 + 443 + 226 + 774 = 1666
func processRequest12(data []string) int {
	sum := 0

	for _, num_string := range data {
		num, _ := strconv.Atoi(num_string)
		copy_num := num

		first_digit := num
		for first_digit >= 10 {
			first_digit /= 10
		}

		processed_num := 0
		p := 1
		for copy_num > 9 {
			processed_num = processed_num + (copy_num % 10) * p
			copy_num = copy_num / 10
			p = p * 10
		}
		processed_num = processed_num + first_digit * p
		p = p * 10
		processed_num = processed_num + first_digit * p

		sum = sum + processed_num
	}

	return sum
}


func handleTask(task string, data []string) string {
	time.Sleep(8 * time.Second)
	switch task {
		case "TASK1":
			return strings.Join(processRequest1(data), ", ")
		case "TASK2":
			result := processRequest2(data)
			return fmt.Sprintf("%d pătrate perfecte", result)
		case "TASK3":
			result := processRequest3(data)
			return fmt.Sprintln(result)
		case "TASK4":
			result := processRequest4(data)
			return fmt.Sprintln(result)
		case "TASK5":
			result := processRequest5(data)
			return fmt.Sprintln(result)
		case "TASK6":
			result := processRequest6(data)
			return fmt.Sprintln(result)
		case "TASK7":
			result := processRequest7(data)
			return fmt.Sprintln(result)
		case "TASK8":
			result := processRequest8(data)
			return fmt.Sprintf("%d cifre", result)
		case "TASK9":
			result := processRequest9(data)
			return fmt.Sprintf("%d cuvinte", result)
		case "TASK10":
			result := processRequest10(data)
			return fmt.Sprintln(result)
		case "TASK11":
			result := processRequest11(data)
			return fmt.Sprintln(result)
		case "TASK12":
			result := processRequest12(data)
			return fmt.Sprintln(result)
		default:
			return "Task necunoscut"
	}
}

func handleClient(conn net.Conn) {
    defer conn.Close()

    for atomic.LoadInt32(&activeGoRoutines) >= int32(config.MaxGoroutines) {
        time.Sleep(3 * time.Second) // asteapta 3 secunde inainte de a verifica din nou
        fmt.Println("Un client asteapta")
    }

    atomic.AddInt32(&activeGoRoutines, 1)
    fmt.Println("Un client a fost acceptat, numar goroutines active:", atomic.LoadInt32(&activeGoRoutines))

    reader := bufio.NewReader(conn)
    client_name, _ := reader.ReadString('\n')
    client_name = strings.TrimSpace(client_name)

    fmt.Printf("Client %s conectat.\n", client_name)
    conn.Write([]byte(fmt.Sprintf("Client %s conectat.\n", client_name)))

    request, _ := reader.ReadString('\n')
    request = strings.TrimSpace(request)
    parts := strings.SplitN(request, ":", 2)

    if len(parts) != 2 {
        conn.Write([]byte("Formatul mesajului este invalid.\n"))
        fmt.Printf("Clientul %s a trimis un format de mesaj invalid.\n", client_name)
        atomic.AddInt32(&activeGoRoutines, -1)
        return
    }

    task, data_string := parts[0], parts[1]
    data := strings.Split(data_string, ",")

    if len(data) > config.ArraySize {
        errorMessage := "Eroare: Dimensiunea mesajului depaseste limita maxima permisa.\n"
        conn.Write([]byte(errorMessage))
        fmt.Printf("Mesajul trimis de client %s depaseste dimensiunea maxima permisa pentru array-ul de input.\n", client_name)
        atomic.AddInt32(&activeGoRoutines, -1)
        return
    }

    fmt.Printf("Client %s a făcut request cu task %s și datele: %v.\n", client_name, task, data)
    conn.Write([]byte("Server a primit requestul.\n"))

    fmt.Println("Server proceseaza datele.")
    response := handleTask(task, data)
    conn.Write([]byte(response + "\n"))

    atomic.AddInt32(&activeGoRoutines, -1)
}


func main() {
	config_file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Eroare la deschiderea fișierului de configurare:", err)
		return
	}
	defer config_file.Close()

	decoder := json.NewDecoder(config_file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("Eroare la decodificarea fișierului de configurare:", err)
		return
	}

	address := fmt.Sprintf(":%d", config.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Eroare la pornirea serverului:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server pornit pe portul %d...\n", config.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Eroare la acceptarea conexiunii:", err)
			continue
		}

		go handleClient(conn)
	}
}
