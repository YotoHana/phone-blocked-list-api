package handlers

import (
	"fmt"
	"log"
	"net/http"
	numlogic "phone-blocked-list-api/numbersLogic"
)

func GetNumberHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Wrong HTTP method", http.StatusMethodNotAllowed)
		log.Fatalf("Wrong HTTP method: %v", r.Method)
	}

	log.Println("GET phoneNumber request")
 	phoneNumberGet := r.Header.Get("phoneNumber")
  	log.Printf("%s formatting...", phoneNumberGet)
  	phoneNumber := numlogic.FormatPhoneNumber(phoneNumberGet)
  	log.Printf("%s formatted", phoneNumber)
  	blockList := numlogic.IsBlocked(phoneNumber)
  	result := numlogic.BlockedString(blockList)

  	fmt.Fprintf(w, "%s is %s", phoneNumber, result)
}

func AddNumberHandler(w http.ResponseWriter, r *http.Request)  {
	log.Println("Getting token from request")
	  tokenGet := r.Header.Get("token")
	  if tokenGet != "9rbAv2uYtWQebBX0rrp4KY3lVcTK6t"{
	  log.Println("Invalid token from request")
	  fmt.Fprintf(w, "Invalid token, access is denied")
	} else {
	  numberGet := r.Header.Get("phoneNumber")
	  var number = make([]string, 0)
	  formattedNumber := numlogic.FormatPhoneNumber(numberGet)
	  number = append(number, formattedNumber)
	  writeNumber, err := numlogic.WriteFile(number)
	  if err != nil {
		fmt.Fprintf(w, "Something wrong, operation failed!")
	  }
	  fmt.Fprintf(w, "Operation completed! %s", writeNumber)
	}
  
  
  }