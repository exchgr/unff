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
  credentials *oauth.Credentials
)

func main() {
  flag.Parse()
  getCredentials()
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

  open.Run(authURL)

  http.HandleFunc("/oauthCallback", func(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
      http.Error(w, fmt.Sprintf("Error! %s\n", err), http.StatusInternalServerError)
      return
    }

    verifier := r.Form["oauth_verifier"][0]
    credentials, vals, err := anaconda.GetCredentials(tempCred, verifier)
    _ = vals
    if err != nil {
      http.Error(w, fmt.Sprintf("Error! %s\n", err), http.StatusInternalServerError)
      return
    }

    if credentials == nil {
      http.Error(w, "Credentials are nil!", http.StatusInternalServerError)
      return
    }

    fmt.Fprintf(w, "Success! You can close this and go back to the terminal.")
  })

  http.ListenAndServe(":9000", nil)

  return false // no error
}
