package packagelib

import (
	"flag"
	"fmt"
)

func TestFlag() {
	username := flag.String("username", "Abisena", "Database Username")
	password := flag.String("password", "*****", "Database password")
	host := flag.String("host", "localhost", "Database host")
	port := flag.Int("port", 3306, "Database port")

	flag.Parse()

	fmt.Println(*username, *password, *host, *port)
}
