/*
* unff.go - Unfollow some friggin' Twitter accounts, you crazy person!
*/

package main

import (
  "fmt"
  "flag"
  "io/ioutil"
  "encoding/json"
  "github.com/chimeracoder/anaconda"
)

// Options
var (
  inactive = flag.Bool("inactive", false, "Unfollow inactive accounts")
  tooactive = flag.Bool("tooactive", false, "Unfollow too-active accounts")
  interactive = flag.Bool("interactive", false, "Select which accounts to unfollow")
)

func main() {
  flag.Parse()
  fmt.Printf("inactive: %t\ntooactive: %t\ninteractive: %t\n", *inactive, *tooactive, *interactive)

  getCredentials()
}

type Keys struct {
  TW_CONSUMER_KEY, TW_CONSUMER_SECRET string
}

func getCredentials() (bool) {
  // Read credentials from JSON file and set them in anaconda

  fileData, err := ioutil.ReadFile("credentials.json")

  if err != nil {
    return true
  }

  keys := Keys{}
  json.Unmarshal(fileData, &keys)
  anaconda.SetConsumerKey(keys.TW_CONSUMER_KEY)
  anaconda.SetConsumerSecret(keys.TW_CONSUMER_SECRET)

  // OAuth

  authURL, tempCred, err := anaconda.AuthorizationURL("")

  if err != nil {
    return true
  }

  return false // no error
}
