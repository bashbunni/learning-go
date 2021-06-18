package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Yo(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Yo, %s", name)
}

func MyYoHandler(w http.ResponseWriter, r *http.Request) {

}

// io.Writer is preferred because it's a "parent" to os.Stdout and bytes.Buffer (both implement io.Writer), so we can do tests with both methods

func main() {
	Yo(os.Stdout, "Deadstar")
}
