package main

import (
	"filestore-server/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/rename", handler.FileUpdateMetaHandler)
	http.HandleFunc("/file/del", handler.FileDelHandler)
	err := http.ListenAndServe(":9100", nil)

	if err != nil {
		fmt.Printf("failed to start server, err:%s", err.Error())
	}
}
