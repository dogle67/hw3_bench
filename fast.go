package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//FastSearch вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	seenBrowsers := make(map[string]struct{})
	uniqueBrowsers := 0
	var isAndroid, isMSIE bool
	var user User

	fmt.Fprintln(out, "found users:")

	for i := 0; scanner.Scan(); i++ {

		line := scanner.Bytes()

		if !(bytes.Contains(line, []byte(`Android`)) || bytes.Contains(line, []byte(`MSIE`))) {
			continue
		}

		user = User{}
		// fmt.Printf("%v %v\n", err, line)
		err := user.UnmarshalJSON(line)
		if err != nil {
			panic(err)
		}

		isAndroid = false
		isMSIE = false

		for _, browser := range user.Browsers {

			switch {
			case strings.Contains(browser, "Android"):
				isAndroid = true
			case strings.Contains(browser, "MSIE"):
				isMSIE = true
			default:
				continue
			}

			if _, ok := seenBrowsers[browser]; !ok {
				seenBrowsers[browser] = struct{}{}
				uniqueBrowsers++
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		email := strings.ReplaceAll(user.Email, "@", " [at] ")
		fmt.Fprint(out, "[")
		fmt.Fprint(out, strconv.Itoa(i))
		fmt.Fprint(out, "] ")
		fmt.Fprint(out, user.Name)
		fmt.Fprint(out, " <")
		fmt.Fprint(out, email)
		fmt.Fprint(out, ">\n")
	}

	fmt.Fprintln(out, "\nTotal unique browsers", uniqueBrowsers)
}
