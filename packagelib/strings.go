package packagelib

import (
	"fmt"
	"strings"
)

func TestStrings() {
	text := "    Hello World    "
	fmt.Println(strings.ReplaceAll(text, "World", "Abisena"))
	fmt.Println(strings.Trim(text, " "))
	fmt.Println(strings.Split(text, " "))
	fmt.Println(strings.Contains(text, "@"))
}
