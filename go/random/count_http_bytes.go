// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "net/http"
import "time"
import "io"
import "io/ioutil"

//import "os"

type Lang struct {
	Name string
	Year int
	URL  string
}

func count(name, url string, c chan<- string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("%s, %s\n", name, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	dt := time.Since(start).Seconds()
	c <- fmt.Sprintf("%s, %d [%.2fs]\n\n", name, n, dt)
}

func main() {
	langs := []Lang{
		Lang{"go", 2009, "http://golang.org"},
		Lang{"python", 2009, "http://www.python.org/"},
		Lang{"scala", 2009, "http://www.scala-lang.org/"},
		Lang{"erlang", 2009, "http://www.erlang.org/"},
	}
	start := time.Now()
	c := make(chan string)
	n := 0
	timeout := time.After(1000 * time.Millisecond)
	for _, lang := range langs {
		n++
		fmt.Printf("[%v]\n", lang)
		go count(lang.Name, lang.URL, c)
	}

	for i := 0; i < n; i++ {
		select {
		case result := <-c:
			fmt.Print(result)
		case <-timeout:
			fmt.Print("Timed out \n")
      return
		}
	}

	fmt.Printf("Total Time: [%.2fs]\n", time.Since(start).Seconds())
}
