package main

import (
  "fmt"
  "strings"
)

type Expression interface {
  // nothing here!
}

type DoubleExpression struct {
  value float64
}

type AdditionExpression struct {
  left, right Expression
}

func Print(e Expression, sb *strings.Builder) {
  switch e := e.(type) {
	  case *DoubleExpression:
		sb.WriteString(fmt.Sprintf("%g", e.value))
	  case *AdditionExpression:
	    sb.WriteString("(")
	    Print(e.left, sb)
	    sb.WriteString("+")
	    Print(e.right, sb)
	    sb.WriteString(")")
	  default: 
		fmt.Println("I dont understand.")	  
  }
  // breaks OCP
  // will work incorrectly on missing case
}

func main() {
  // 1+(2+3)
  e := &AdditionExpression{
    &DoubleExpression{1},
    &AdditionExpression{
      left:  &DoubleExpression{2},
      right: &DoubleExpression{3},
    },
  }
  sb := strings.Builder{}
  Print(e, &sb)
  fmt.Println(sb.String())
}