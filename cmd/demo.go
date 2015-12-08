package main

import "github.com/thriqon/fort"
import "time"
import "strconv"

func main() {
	f := fort.NewContainer()
	te := f.AddTextElement("id", "Hello, World!")

	for i := 0; i < 10; i++ {
		time.Sleep(400 * time.Millisecond)
		te.Message = strconv.Itoa(i) + " iterations"
		f.Refresh()
	}
}
