package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
  "os"
  "bufio"
  "io"
)

func main() {
  http.HandleFunc("/", HandlerGet)
  http.HandleFunc("/add", HandlePost)
  log.Println("Server is running at localhost:8080")
  http.ListenAndServe(":8080", nil)
}

func HandlerGet(w http.ResponseWriter, r *http.Request) {
  log.Println("GET phoneNumber request")
  phoneNumberGet := r.Header.Get("phoneNumber")
  log.Printf("%s formatting...", phoneNumberGet)
  phoneNumber := formatPhoneNumber(phoneNumberGet)
  log.Printf("%s formatted", phoneNumber)
  blockList := isBlocked(phoneNumber)
  result := blockedString(blockList)

  fmt.Fprintf(w, "%s is %s", phoneNumber, result)
}

func isBlocked(number string) bool {
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

func blockedString(blockList bool) string {
  if blockList == true {
    return "blocked"
  }
  return "non blocked"
}

func formatPhoneNumber(phoneNumber string) string {
  //symbolArr := [5]string{"(", ")", "-", " ", "+"}
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
    formLine := formatPhoneNumber(line)
		result = append(result, formLine)
    log.Printf("Showing numbers from file: %s", formLine)
    }
    log.Println(result)
    return result
}

func HandlePost(w http.ResponseWriter, r *http.Request)  {
  log.Println("Getting token from request")
	tokenGet := r.Header.Get("token")
	if tokenGet != "9rbAv2uYtWQebBX0rrp4KY3lVcTK6t"{
    log.Println("Invalid token from request")
    fmt.Fprintf(w, "Invalid token, access is denied")
  } else {
    numberGet := r.Header.Get("phoneNumber")
    var number = make([]string, 0)
    formattedNumber := formatPhoneNumber(numberGet)
    number = append(number, formattedNumber)
    writeNumber, err := writeFile(number)
    if err != nil {
      fmt.Fprintf(w, "Something wrong, operation failed!")
    }
    fmt.Fprintf(w, "Operation completed! %s", writeNumber)
  }


}

func writeFile(data []string) (string, error){
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