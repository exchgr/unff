/*
* unff.go - Unfollow some friggin' Twitter accounts, you crazy person!
*/

package main

import (
  "fmt"
  "flag"
  "io/ioutil"
  "encoding/json"
  "net/http"
  "html"
  "github.com/chimeracoder/anaconda"
  "github.com/skratchdot/open-golang/open"
  "github.com/garyburd/go-oauth/oauth"
)

// Options
var (
  inactive = flag.Bool("inactive", false, "Unfollow inactive accounts")
  tooactive = flag.Bool("tooactive", false, "Unfollow too-active accounts")
  interactive = flag.Bool("interactive", false, "Select which accounts to unfollow")
)

// Global variables
var (
  tempCred *oauth.Credentials
)

func main() {
  flag.Parse()
  fmt.Printf("inactive: %t\ntooactive: %t\ninteractive: %t\n", *inactive, *tooactive, *interactive)

  getCredentials()

  http.HandleFunc("/oauthCallback", func(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
      fmt.Printf("Error: %v", err)
      return true
    }

    fmt.Fprintf(w, "Hello, %v", html.EscapeString(string(body)))
  })

  http.ListenAndServe(":9000", nil)
}

// Twitter keys struct for reading from JSON
type Keys struct {
  TW_CONSUMER_KEY, TW_CONSUMER_SECRET string
}

func getCredentials() (bool) {
  // Read credentials from JSON file and set them in anaconda

  fileData, err := ioutil.ReadFile("credentials.json")
  if err != nil {
    fmt.Printf("Error: %v", err)
    return true
  }

  keys := Keys{}
  json.Unmarshal(fileData, &keys)
  anaconda.SetConsumerKey(keys.TW_CONSUMER_KEY)
  anaconda.SetConsumerSecret(keys.TW_CONSUMER_SECRET)

  // OAuth

  authURL, tempCred, err := anaconda.AuthorizationURL("")
  if err != nil {
    fmt.Printf("Error: %v", err)
    return true
  }

  _ = tempCred
  open.Run(authURL)

  return false // no error
}
