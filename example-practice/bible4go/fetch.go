package main

import (
	"fmt"
	_ "io"
	// "io/ioutil"
	_ "net/http"
	_ "os"
	"strings"
)

func MakeAddSuffix(prefix string) func(string) string {
	return func(name string) string {
		if !strings.HasPrefix(name, prefix) {
			return prefix + name
		}
		return name
	}
}

const boilingF = 212.0

func main() {
	/*
		prefix := MakeAddSuffix("http://")
		for _, url := range os.Args[1:] {
			url = prefix(url)

			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
				os.Exit(1)
			}
			// b, err := ioutil.ReadAll(resp.Body)
			_, err = io.Copy(os.Stdout, resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s : %v\n", url, err)
				os.Exit(1)
			}
			fmt.Println(url)
			fmt.Println(resp.Status)
		}
	*/
	var f = boilingF
	fmt.Printf("boiling point = %g°F or %g°C\n", f, FtoC(f))
}

func FtoC(f float64) float64 {
	return (f - 32) * 5 / 9
}
