package main

import (
	"fmt"
	"os"

	"github.com/levigross/grequests"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

func download(c *cli.Context) {
	baseDir, _ := homedir.Expand("~/.atcoder_next")
	os.Mkdir(baseDir, os.ModePerm)

	endpointBase := "https://kenkoooo.com/atcoder/atcoder-api"

	fmt.Println("Download problems.json")
	downloadJson(endpointBase+"/info/merged-problems", baseDir+"/problems.json")

	fmt.Println("Download contests.json")
	downloadJson(endpointBase+"/info/contests", baseDir+"/contests.json")
}

func downloadJson(url string, filename string) error {
	resp, err := grequests.Get(url, nil)
	if err != nil {
		fmt.Println("Unable to access: ", url)
		return err
	}
	if err := resp.DownloadToFile(filename); err != nil {
		fmt.Println("Failed to download: ", err.Error())
		return err
	}
	return nil
}
