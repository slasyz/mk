package complete

import (
	"fmt"
	"os"
)

func Complete() {
	// TODO: autocompletion
	if _, ok := os.LookupEnv("COMP_LINE"); !ok {
		return
	}
	fmt.Println("+++++")
	fmt.Println("-----")
	os.Exit(0)
}
