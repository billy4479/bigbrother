package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const updateUrl = "localhost:4479/get-client"

func main() {

	workdir := filepath.Join(os.TempDir(), "bigbrother")
	err := os.MkdirAll(workdir, 0700)
	if err != nil {
		panic(err)
	}

	logfile, err := os.Create(filepath.Join(workdir, "launcher.log"))
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	checkErr := func(err error) {
		if err != nil {
			logfile.WriteString(err.Error())
			panic(err)
		}
	}

	logfile.WriteString("Making request")
	res, err := http.Get(updateUrl)
	checkErr(err)

	clientPath := filepath.Join(workdir, "client")
	if runtime.GOOS == "windows" {
		clientPath += ".exe"
	}

	logfile.WriteString("Creating client file")
	client, err := os.Create(clientPath)
	checkErr(err)
	defer client.Close()

	logfile.WriteString("Download client")
	_, err = io.Copy(client, res.Body)
	checkErr(err)

	res.Body.Close()
	client.Close()

	cmd := exec.Command(clientPath)
	err = cmd.Run()
	checkErr(err)

	os.Remove(clientPath)
}
