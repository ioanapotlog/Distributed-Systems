package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func readArrayFromFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var array_content string
	if scanner.Scan() {
		array_content = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	array_content = strings.TrimSpace(array_content)
	array_elements := strings.Split(array_content, ",")

	for i := range array_elements {
		array_elements[i] = strings.TrimSpace(array_elements[i])
	}

	formatted_array := strings.Join(array_elements, ",")
	return formatted_array, nil
}

func main() {
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		fmt.Println("Eroare la conectarea la server:", err)
		return
	}
	defer conn.Close()

	input := bufio.NewReader(os.Stdin)

	fmt.Print("Client name: ")
	client_name, _ := input.ReadString('\n')
	client_name = strings.TrimSpace(client_name)
	fmt.Fprintln(conn, client_name)

	fmt.Print("Task (TASK1/TASK2/...): ")
	task, _ := input.ReadString('\n')
	task = strings.TrimSpace(task)

	file_name := "ArrayInputFiles/" + client_name + "_input.txt"

	array_content, err := readArrayFromFile(file_name)
	if err != nil {
		fmt.Println("Eroare la citirea fișierului de intrare:", err)
		return
	}

	fmt.Fprintln(conn, task + ":" + array_content)

	connection_confirmation, _ := bufio.NewReader(conn).ReadString('\n')
	connection_confirmation = strings.TrimSpace(connection_confirmation)
	fmt.Println(connection_confirmation)

	request_confirmation, _ := bufio.NewReader(conn).ReadString('\n')
	request_confirmation = strings.TrimSpace(request_confirmation)
	fmt.Println(request_confirmation)

	if strings.Contains(request_confirmation, "Eroare") {
		fmt.Println("Client a oprit execuția din cauza unei erori de la server.")
		return
	}

	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Client a primit raspunsul:", strings.TrimSpace(response))
}