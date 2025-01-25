// string conversion  (strconv)
package packagelib

import (
	"fmt"
	"strconv"
)

// func testStrconv() {
// 	fmt.Println(strconv.Itoa(100))                        // int to string dan Atoi itu string to int
// 	fmt.Println(strconv.ParseInt("100", 2, 64))           // string to int
// 	fmt.Println(strconv.ParseFloat("100.123", 64))        // string to float
// 	fmt.Println(strconv.ParseBool("true"))                // string to bool
// 	fmt.Println(strconv.FormatBool(true))                 // bool to string
// 	fmt.Println(strconv.FormatFloat(100.123, 'f', 2, 64)) // float to string
// 	fmt.Println(strconv.FormatInt(100, 16))               // int to string
// }

func TestStrconv() {
	boolValue, err := strconv.ParseBool("salah")
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	fmt.Println(boolValue)
}
