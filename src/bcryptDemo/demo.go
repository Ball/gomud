
package main
import (
  "code.google.com/p/go.crypto/bcrypt"
  "fmt"
)

func main(){
  fmt.Printf("Let's try encrypting 'Toofer'\n")
  hpass, err := bcrypt.GenerateFromPassword([]byte("Toofer"), bcrypt.DefaultCost)
  if err != nil {
    panic(err)
  }
  fmt.Printf("We got <<%s>>\n", hpass)
  fmt.Printf("And when we authenticate...\n")
  err = bcrypt.CompareHashAndPassword(hpass, []byte("Toofer"))
  if err != nil {
    panic(err)
  }
  fmt.Printf("Success!\n")
  fmt.Printf("Just to be safe, let's make sure the negative works.\n")
  err = bcrypt.CompareHashAndPassword(hpass, []byte("Toofr"))
  if err == nil {
    panic(err)
  }
  fmt.Printf("Success!  Well, an expected failure of authentication.")
}
