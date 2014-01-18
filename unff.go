/*
* unff.go - Unfollow some friggin' Twitter accounts, you crazy person!
*/

package main

import (
  "fmt"
  "flag"
  // "github.com/chimeracoder/anaconda"
)

// Options
var inactive, tooactive, interactive *bool

func main() {
  getFlags()
  fmt.Printf("inactive: %t\ntooactive: %t\ninteractive: %t\n", *inactive, *tooactive, *interactive)
}

func getFlags() {
  inactive = flag.Bool("inactive", false, "Unfollow inactive accounts")
  tooactive = flag.Bool("tooactive", false, "Unfollow too-active accounts")
  interactive = flag.Bool("interactive", false, "Select which accounts to unfollow")
  flag.Parse()
}
