package main;

import (
  "fmt"
  "time"
  "flag"
)

func main() {
  // CLI flag to get the time zone
  // default is local
  tzPtr := flag.String("z", "Local", "a string")
  flag.Parse()

  // Loads the timezone
  loc, err := time.LoadLocation(*tzPtr)
  if err != nil {
    fmt.Println("Error loading location: ", err)
    return
  }

  t := time.Now().In(loc)
  fmt.Println("Time: ", t.Format(time.RFC850))
}
