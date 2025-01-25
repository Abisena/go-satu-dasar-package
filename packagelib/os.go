package packagelib

import (
	"fmt"
	"os"
)

func TestOs() {
	args := os.Args
	for _, argas := range args {
		fmt.Println(argas)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		panic(err)
	}

}
