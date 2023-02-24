package main
import ("fmt"
//"strings"
)

type myStruct1 struct{
 i1 int
 f1 float32
	mm string
 str string


}

func main(){
 ms := new(myStruct1)
 ms.i1 = 10
 ms.f1 = 10.50
ms.str = "Abiodun"
 ///ms.mm :=strings.ToUpper(str)
 //ms.mm = "name"
fmt.Printf("The int is %d \n", ms.i1)
fmt.Printf("The float is %f \n", ms.f1)
fmt.Printf("The string is %s \n", ms.str)
fmt.Println(ms)
// fmt.Printf("The value of i1 is %d int \n, and the value of f1 is %f float, \n also the value of str is %s string ", ms.i1, ms.f1, ms.str)





}