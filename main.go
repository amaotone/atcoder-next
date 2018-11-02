package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

type contestInfo struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Timestamp string `json:"start_epoch_second"`
	Rated     string `json:"rate_change"`
}

type problemInfo struct {
	Id          string `json:"id"`
	Contest     string `json:"contest_id"`
	Title       string `json:"title"`
	SolverCount int    `json:"solver_count"`
	Point       int    `json:"point"`
}

func _fetch(c *cli.Context) {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Cannot find homedir:" + err.Error())
		return
	}
	configDir := home + "/.atcoder_next"

	var data []byte

	data, _ = ioutil.ReadFile(configDir + "/contests.json")
	var contests []contestInfo
	json.Unmarshal(data, &contests)
	fmt.Println(contests)

	data, _ = ioutil.ReadFile(configDir + "/problems.json")
	var problems []problemInfo
	json.Unmarshal(data, &problems)
	for _, c := range problems {
		fmt.Println(c)
	}
}

func show(c *cli.Context) {
	fmt.Println("show command")
}

func main() {
	app := cli.NewApp()
	app.Name = "AtCoder Next"
	app.Usage = "Tell you the next problem to solve."
	app.Version = "0.1"
	app.Commands = []cli.Command{
		{
			Name:   "download",
			Usage:  "Download the information of contests and problems.",
			Action: download,
		},
		{
			Name:   "show",
			Usage:  "Show the next problem to solve.",
			Action: show,
		},
	}
	app.Run(os.Args)
}
