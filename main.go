package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const SVG = "svg";
const PNG = "png";

func downloadFile(url string, output string) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

}

func main() {

	if (len(os.Args) < 3) {
		panic("URL and output are required")
	}

	url := os.Args[1];
	output := os.Args[2]+"."+PNG;

	useUrl := strings.Replace(url, SVG, PNG, 1);

	fmt.Printf("Downloading from %s, to %s", useUrl, output);
	downloadFile(useUrl, output)
}