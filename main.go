package main

import (
	"fmt"

	aws0 "github.com/aaronfriel/go-multimod/friel/aws"
	aws1 "github.com/aaronfriel/go-multimod/stack72/aws"
)

func main() {
	fmt.Println(aws0.Provider())
	fmt.Println(aws1.Provider())
}
