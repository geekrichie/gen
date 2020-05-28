package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/chunkbody", HelloHandler)
	http.HandleFunc("/bigfile", BigFileHandler)
	http.ListenAndServe(":9999", nil)
}
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	var sArr [3]string = [3]string{"a", "123", "abcdef"}
	w.Header().Set("Transfer-Encoding", "chunked")
	for _, str := range sArr {
		resp := strconv.Itoa(len(str)) + "\r\n" + str + "\r\n"
		fmt.Printf(string(len(str)))
		w.Write([]byte(resp))
	}
	w.Write([]byte("0\r\n\r\n"))
}

func BigFileHandler(w http.ResponseWriter, req *http.Request) {
	dir := "D:\\迅雷下载"
	filename := "极速车王.HD1280高清中英双字版.mp4"
	data, err := ioutil.ReadFile(dir + "\\" + filename)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(data)
}
