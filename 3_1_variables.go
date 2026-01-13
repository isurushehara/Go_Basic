package main
import ("fmt")

func main() {

  // Declaring (Creating) Variables

  // 1. With the var keyword
  // var variablename type = value
  var age int = 25

  //2. With the := sign:
  // variablename := value
  name := "Alice"

  //Example usages
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)

  // Variable Declaration With Initial Value
  var student3 string = "John" //type is string
  var student4 = "Jane" //type is inferred
  y := 2 //type is inferred

  fmt.Println(student3)
  fmt.Println(student4)
  fmt.Println(y)


  // Variable Declaration Without Initial Value
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)

  // Value Assignment After Declaration
  var student1 string
  student1 = "John"
  fmt.Println(student1)
}