package main
 
import (
 "flag"
 "fmt"
 "os"
 "strings"
 lab2 "github.com/Ararni/postfix-to-infix"
)
 
var (
 inputExpression = flag.String("e", "", "Expression to compute")
 fileInput  = flag.String("f", "", "File containing expression to compute")
 fileOutput = flag.String("o", "", "Write expression to file")
 // TODO: Add other flags support for input and output configuration.
)
 
func main() {
 flag.Parse()
  handler := &lab2.ComputeHandler{}
 if len(*fileInput) == 0{
   handler.Input = strings.NewReader(*inputExpression)
 } else {
   file, err := os.Open(*fileInput)
   if err != nil {
     fmt.Fprintln(os.Stderr, "error")
     return
   }
   handler.Input = file
 }
 if len(*fileOutput) == 0{
  handler.Output = os.Stdout
 } else {
   file, err := os.OpenFile(*fileOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
   if err != nil {
     fmt.Fprintf(os.Stderr, "error")
     return
   }
   handler.Output = file
 }
 
 err := handler.Compute()
 if err != nil {
   fmt.Fprintf(os.Stderr, err.Error())
   return
 }
}
 
