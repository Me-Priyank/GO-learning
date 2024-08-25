package main

import "fmt"

const Logintoken string = "ghsgdyuvicyedv48743bsuxcberru8493bryv93by9v" // Note :- To make anything public or , kisi variable ya constant ko other files me use karne ke liye first letter capital rakha jata hai .

func main() {
	var name string = "Priyank"
	fmt.Println(name)
	fmt.Printf("Variable is of type : %T \n",name)

	var isLogges bool = true
	fmt.Println(isLogges)
	fmt.Printf("Variable is of type : %T \n",isLogges)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n",smallVal)

	var smallFloat float32 = 4422.434231212343
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n",smallFloat)


	// default values and aliases

	var try1 int 
	fmt.Println(try1)
	fmt.Printf("Variable is of type : %T \n",try1)

	var try2 string 
	fmt.Println(try2)
	fmt.Printf("Variable is of type : %T \n",try2)

	var try3 uint64 
	fmt.Println(try3)
	fmt.Printf("Variable is of type : %T \n",try3)

	// implicit way of defining variables :-

	var userName = "Priyank Verma"
	fmt.Println(userName)

	// no var method (volorus operator method) :-

	userCount := 3000 // can be used inside fn/method only , cannot be used in global 
	fmt.Println(userCount)

	// Printing public value :-

	fmt.Println(Logintoken)
	fmt.Printf("Variable is of type : %T \n",Logintoken)
}