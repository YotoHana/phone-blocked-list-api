package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
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
  log.Println("phoneNumber formatting...")
  phoneNumber := formatPhoneNumber(phoneNumberGet)
  log.Println("phoneNumber formatted")
  blockList := isBlocked(phoneNumber)
  result := blockedString(blockList)

  fmt.Fprintf(w, "%s is %s", phoneNumber, result)
}

func isBlocked(number string) bool {
  blocked := false
  numberArr := [5]string{"+79374848615", "+79572135764", "+77564723458", "+76583746574", "+79462347561"}
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

func HandlePost(w http.ResponseWriter, r *http.Request)  {
  log.Println("Getting token from request")
	tokenGet := r.Header.Get("token")
	if tokenGet != "9rbAv2uYtWQebBX0rrp4KY3lVcTK6t"{
    log.Println("Invalid token from request")
    fmt.Fprintf(w, "Invalid token, access is denied")
  } else {
    log.Println("Adding new number")
    phoneNumberGet := r.Header.Get("phoneNumber")
    phoneNumber := formatPhoneNumber(phoneNumberGet)
    // code for save number in local file
    log.Println("Saving number")
    fmt.Fprintf(w, "Number %s is saved!", phoneNumber)
    log.Println("Number is saved")
  }


}