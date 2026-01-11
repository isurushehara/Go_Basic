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
}