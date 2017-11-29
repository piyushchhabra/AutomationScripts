// paytm
package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"github.com/antchfx/xpath"
	"github.com/antchfx/xquery/html"

	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 15000)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			test1(`https://paytm.com/movies/hyderabad/pvr-inorbit-cyberabad-c/220?fromdate=2017-11-24`,
				"//div[@class='_2tt5']",
				"Justice")
		}
	}()
	time.Sleep(time.Millisecond * 60000 * 60 * 10)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func test1(url string, xp string, keyword string) {
	// Load HTML file.
	f, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// Parse HTML document.
	doc, err := htmlquery.Parse(f.Body)
	if err != nil {
		panic(err)
	}

	expr := xpath.MustCompile(xp)
	//	fmt.Printf("%f \n", expr.Evaluate(htmlquery.CreateXPathNavigator(doc)).(float64))
	iter := expr.Evaluate(htmlquery.CreateXPathNavigator(doc)).(*xpath.NodeIterator)
	for iter.MoveNext() {
		str := iter.Current().Value()
		fmt.Print(str + ", ")

		if strings.Contains(str, keyword) {
			fmt.Print("\nFound Movie")
			//openBrowser(url)
		}
	}

	fmt.Println("\n----------------------------------------------------")

}

func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
