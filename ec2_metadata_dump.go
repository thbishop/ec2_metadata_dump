package main

import "bufio"
import "fmt"
import "encoding/json"
import "net/http"
import "os"
import "strings"

func getData(url string) (data []string) {
	resp, err := http.Get(url)
	if err != nil {
		os.Stderr.Write([]byte("Error: " + err.Error()))
		os.Exit(1)
	}

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		data = append(data, strings.TrimRight(scanner.Text(), "\n"))
		if err != nil {
			break
		}
	}
	return
}

func crawlData(url string) map[string]interface{} {
	data := make(map[string]interface{})
	urlData := getData(url)

	for _, line := range urlData {
		switch {
		default:
			data[line] = getData(url + line)[0]
		case line == "dynamic":
			break
		case line == "meta-data":
			data[line] = crawlData(url + line + "/")
		case line == "user-data":
			data[line] = strings.Join(getData(url+line+"/"), "")
		case strings.HasSuffix(line, "/"):
			data[line[:len(line)-1]] = crawlData(url + line)
		case strings.HasSuffix(url, "public-keys/"):
			keyId := strings.SplitN(line, "=", 2)[0]
			data[line] = crawlData(url + keyId + "/")
		}
	}
	return data
}

func jsonData(url string) string {
	data, err := json.MarshalIndent(crawlData(url), "", "    ")
	if err != nil {
		os.Stderr.Write([]byte("Error: " + err.Error()))
		os.Exit(1)
	}
	return string(data)
}

func main() {
	url := "http://169.254.169.254/latest/"

	os.Stdout.WriteString(jsonData(url))
	fmt.Println("")
}
