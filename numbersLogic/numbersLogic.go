package numlogic

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func IsBlocked(number string) bool {
	blocked := false
	numberArr := showNumbers()
	for _, s := range numberArr {
		if number == s {
			blocked := true
			return blocked
		}
	}
	return blocked
}

func BlockedString(blockList bool) string {
	if blockList == true {
		return "blocked"
	}
	return "non blocked"
}

func FormatPhoneNumber(phoneNumber string) string {
	phoneNumber = strings.ReplaceAll(phoneNumber, "(", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, ")", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "+", "")
	phoneNumber = "+7" + phoneNumber[len(phoneNumber)-10:]
	return phoneNumber
}

func showNumbers() []string {
	file, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var result = make([]string, 0)
	for {
		line, err := reader.ReadString(' ')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return nil
			}
		}
		formLine := FormatPhoneNumber(line)
		result = append(result, formLine)
		log.Printf("Showing numbers from file: %s", formLine)
	}
	log.Println(result)
	return result
}

func WriteFile(data []string) (string, error){
	file, err := os.OpenFile("numbers.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		return "error crating file", err
	}

	dataWriter := bufio.NewWriter(file)

	for _, value := range data {
		_, _ = dataWriter.WriteString(value + " ")
	}

	dataWriter.Flush()
	file.Close()
	return "all good", nil
}