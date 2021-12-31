// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		//练习1.8
		hasHttp := strings.HasPrefix(url, "http://")
		if !hasHttp {
			url = "http://" + url
			fmt.Printf("url: %v\n", url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		// _, err = io.Copy(os.Stdout, resp.Body) ##直接copy，不用使用b作为缓存区
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)                        //直接将字节数据按照字符串的格式打印
		fmt.Printf("resp.Status: %v\n", resp.Status) //打印响应状态码
	}
}

//!-
