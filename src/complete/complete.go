package complete

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/shlex"

	"github.com/slasyz/mk/src/schema"
)

func debug() {
	f, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "COMP_LINE=***%s***\n", os.Getenv("COMP_LINE"))
	fmt.Fprintf(f, "COMP_WORDS=***%s***\n", os.Getenv("COMP_WORDS"))
	fmt.Fprintf(f, "COMP_CWORD=%s\n", os.Getenv("COMP_CWORD"))
	fmt.Fprintf(f, "COMP_POINT=%s\n", os.Getenv("COMP_POINT"))
	fmt.Fprintln(f)
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "COMP") {
			fmt.Fprintln(f, "->", env)
		} else {
			fmt.Fprintln(f, env)
		}
	}
	fmt.Fprintln(f)

	fmt.Fprintf(f, "args=%v\n", os.Args)
	fmt.Fprintln(f)
	shlexRes, err := shlex.Split(os.Getenv("COMP_LINE"))
	if err != nil {
		fmt.Fprintf(f, "shlex err: %s\n", err)
	} else {
		for i, el := range shlexRes {
			fmt.Fprintf(f, "shlex#%d: %v\n", i, el)
		}
	}
}

func Complete() {
	compLine, ok := os.LookupEnv("COMP_LINE")
	if !ok {
		return
	}
	debug()

	compPoint, err := strconv.Atoi(os.Getenv("COMP_POINT"))
	if err != nil {
		fmt.Println("Error parsing COMP_POINT:", err)
		return
	}
	compWord, err := strconv.Atoi(os.Getenv("COMP_CWORD"))
	if err != nil {
		fmt.Println("Error parsing COMP_CWORD:", err)
		return
	}

	root, err := schema.Load("mk.yml")
	if err != nil {
		return
	}

	res, err := Generate(root, compLine, compPoint, compWord)
	if err != nil {
		return
	}
	for _, el := range res {
		fmt.Println(el)
	}
	os.Exit(0)
}
