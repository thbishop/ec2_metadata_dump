package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func testJsonData() string {
	return `{
		"meta-data": {
			"ami-id": "12345",
			"block-device-mapping": {
				"ami": "/dev/sda1",
				"root": "/dev/sda1"
			},
			"public-keys": {
				"0=foo_key": {
					"openssh-key": "my ssh key"
				}
			}
		},
		"user-data": "foo bar"
	}`
}

func crawlDataSetup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	mux.HandleFunc("/latest/meta-data/public-keys/0/openssh-key", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("my ssh key"))
	})

	mux.HandleFunc("/latest/meta-data/public-keys/0/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("openssh-key"))
	})

	mux.HandleFunc("/latest/meta-data/public-keys/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("0=foo_key"))
	})

	mux.HandleFunc("/latest/meta-data/block-device-mapping/ami", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/dev/sda1"))
	})

	mux.HandleFunc("/latest/meta-data/block-device-mapping/root", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/dev/sda1"))
	})

	mux.HandleFunc("/latest/meta-data/block-device-mapping/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root\nami"))
	})

	mux.HandleFunc("/latest/meta-data/ami-id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("12345"))
	})

	mux.HandleFunc("/latest/meta-data/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ami-id\nblock-device-mapping/\npublic-keys/"))
	})

	mux.HandleFunc("/latest/user-data/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("foo bar"))
	})

	mux.HandleFunc("/latest/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("dynamic\nmeta-data\nuser-data"))
	})
}

func crawlDataTeardown() {
	server.Close()
}

func minimizeJson(data string) string {
	re := regexp.MustCompile(`\s`)
	return re.ReplaceAllString(data, "")
}

func TestJsonData(t *testing.T) {
	crawlDataSetup()
	defer crawlDataTeardown()

	data := jsonData(server.URL + "/latest/")

	if !reflect.DeepEqual(minimizeJson(data), minimizeJson(testJsonData())) {
		t.Errorf("jsonData:\n\texpected: %+v\n\t     got: %+v", minimizeJson(testJsonData()), minimizeJson(data))
	}
}
