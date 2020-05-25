package main

import "fmt"

type Person struct {
  FirstName, MiddleName, LastName string
}

func (p *Person) Names() []string {
  return []string{p.FirstName, p.MiddleName, p.LastName}
}

func (p *Person) NamesGenerator() <-chan string {
  out := make(chan string)
  go func ()  {
    defer close(out)
    out <- p.FirstName
    if len(p.MiddleName) > 0 {
      out <- p.MiddleName
    } 
    out <- p.LastName
  }()
  return out
}

func main() {
  p := Person{"Alexander", "", "Bell"}
  //for _, name := range p.Names() {
  //  fmt.Println(name)
  //}
  for name := range p.NamesGenerator() {
    fmt.Println(name)
  }
}
