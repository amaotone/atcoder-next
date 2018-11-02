package main

import (
	"fmt"

	"github.com/levigross/grequests"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

func download(c *cli.Context) {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Cannot find home: " + err.Error())
		return
	}
	baseDir := home + "/.atcoder_next"
	endpointBase := "https://kenkoooo.com/atcoder/atcoder-api"

	fmt.Println("Download problems.json")
	downloadJson(endpointBase+"/info/merged-problems", baseDir+"/problems.json")

	fmt.Println("Download contests.json")
	downloadJson(endpointBase+"/info/contest", baseDir+"/contests.json")
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
