package main

import (
	"log"
	"strconv"

	"github.com/heathcliff26/containers/apps/simple-fileserver/pkg/fileserver"
)

var (
	webroot       string
	port          int
	sslCert       string
	sslKey        string
	withoutIndex  bool
	enableLogging bool
)

func main() {
	parseFlags()

	fs := fileserver.NewFileserver(webroot, !withoutIndex)
	log.Printf("Serving content from %s", webroot)

	fs.Handle("/")

	if sslCert != "" && sslKey != "" {
		log.Print("Enabling ssl")
		fs.UseSSL(sslCert, sslKey)
	}

	fs.Log = enableLogging

	log.Printf("Listening on :%d", port)
	err := fs.ListenAndServe(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}
}
