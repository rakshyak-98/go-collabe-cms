// Variable is a momeory space to store a value of a specific type

package main

import "fmt"

var c, java, python bool //declared at Package level

var a, b int = 1, 2

func main() {
	var i int                       //declared at Function level
	fmt.Println(c, java, python, i) // false false false 0

	var c, python, java = true, false, "no!"
	fmt.Println(a, b, c, python, java) //1 2 true false no!

	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

	k := 10
	isJava := true

	fmt.Println(k, isJava) //10 true

	var u int8 = 25

	//Type
	fmt.Printf("Type: %v, Value: %T\n ", k, k)           //Type: 10, Value: int
	fmt.Printf("Type: %v, Value: %T \n", isJava, isJava) //Type: true, Value: bool
	fmt.Printf("Type: %v, Value: %T \n", u, u)           //Type: 25, Value: int8

	//Type Conversion
	//The expression T(v) converts the value v to the type T.

	var f float32 = float32(u)

	int1 := 3
	int2 := float32(int1)

	fmt.Printf("Type: %v, Value: %T \n", f, f)       //Type: 25, Value: float32
	fmt.Printf("Type: %v, Value: %T \n", int2, int2) //Type: 3, Value: float32

	// Constants are declared like variables, but with the const keyword. Constants can be character, string, boolean, or numeric values. Constants cannot be declared using the := syntax.
	const me = "Akash"
	fmt.Println(me)
}

/* Go's basic types are

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
*/
