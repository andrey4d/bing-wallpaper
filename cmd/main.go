/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func main() {

	resp, err := http.Get("https://www.bing.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	body, _ := io.ReadAll(resp.Body)

	re := regexp.MustCompile("th\\?id=(.*?_tmb.jpg)")
	urlGroup := re.FindStringSubmatch(string(body))
	url := urlGroup[len(urlGroup)-1]

	re = regexp.MustCompile("tmb")
	uhd := "UHD"
	url = re.ReplaceAllLiteralString(url, uhd)

	fmt.Printf("%s\n", url)

	// scanner := bufio.NewScanner(resp.Body)

	// for i := 0; scanner.Scan(); i++ {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	panic(err)
	// }

}
