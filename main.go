package main

import (
  "fmt"
  "net/http"
  "strings"
)

func main() {
  http.HandleFunc("/", Handler)
  http.ListenAndServe(":8080", nil)

}

func Handler(w http.ResponseWriter, r *http.Request) {
  phoneNumberGet := r.Header.Get("phoneNumber")
  phoneNumber := formatPhoneNumber(phoneNumberGet)
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