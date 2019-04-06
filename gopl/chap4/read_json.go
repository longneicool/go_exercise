package main

// 本文描述了一个典型的读取网站的ｊｓｏｎ文件并且格式化打印出来的示例

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://xkcd.com/571/info.0.json"

func main() {
	resp, _ := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Failed to get from %s!status %s", url, resp.Status)
		os.Exit(1)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to Readall from resp!%s", err)
		os.Exit(1)
	}

	var out bytes.Buffer
	json.Indent(&out, data, "", "    ")
	out.WriteTo(os.Stdout)
}
