package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

// Commands has all commands vars.
// A command has cmd's help text and behavior.
// Behavior function named "doFunction".
// Example: cmdList -> doList
var Commands = []cli.Command{
	cmdList,
	cmdNew,
	cmdRoot,
}

var cmdList = cli.Command{
	Name:        "list",
	Usage:       "show post list",
	Description: `Not yet`,
	Action:      doList,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "reverse, r",
			Usage: "reverse list",
		},
	},
}
var cmdNew = cli.Command{
	Name:        "new",
	Usage:       "make a post",
	Description: `Not yet`,
	ArgsUsage:   "<title>",
	Action:      doNew,
	//	Flags: []cli.Flag{},
}
var cmdRoot = cli.Command{
	Name:        "root",
	Usage:       "show root directory",
	Description: `Not yet`,
	Action:      doRoot,
	//	Flags: []cli.Flag{},
}

func doList(c *cli.Context) error {
	files, err := ioutil.ReadDir(getRoot())
	if err != nil {
		log.Fatal("Can't get root directory files.")
	}
	if c.Bool("reverse") {
		files = reverseFiles(files)
	}
	for _, file := range files {
		if getLastChar(file.Name()) != '~' {
			fmt.Println(file.Name())
		}
	}
	return nil
}

func reverseFiles(files []os.FileInfo) []os.FileInfo {
	if len(files) == 0 {
		return files
	}
	return append(reverseFiles(files[1:]), files[0])
}

func doNew(c *cli.Context) error {
	title := c.Args().First()
	if title == "" {
		cli.ShowCommandHelp(c, "new")
		os.Exit(1)
	}
	today := time.Now().Format("2006-01-02")
	newFileName := fmt.Sprintf("%s-%s.md", today, title)
	if !confirm("you make new file?(" + newFileName + ")") {
		fmt.Println("cancel")
		return nil
	}
	makeNewFile(newFileName)
	fmt.Println("new file: " + newFileName)
	return nil
}

func makeNewFile(filename string) {
	var nf *os.File
	if !isExist("~/tmp.md") {
		nf, err := os.Create(getRoot() + "/" + filename)
		if err != nil {
			log.Fatal("Can't make new file!")
			os.Exit(1)
		}
		defer nf.Close()
		return
	}

	tmp, err := os.Open("~/tmp.md")
	if err != nil {
		log.Fatal("Can't open tmp file!")
		os.Exit(1)
	}
	defer tmp.Close()

	_, err = io.Copy(nf, tmp)
	if err != nil {
		log.Fatal("Can't copy to new file!")
		os.Exit(1)
	}
}

func doRoot(c *cli.Context) error {
	fmt.Println(getRoot())
	return nil
}

func getRoot() string {
	root := os.Getenv("JK_WRITE_ROOT")
	if root == "" {
		log.Fatal("You should set fullpath to JK_WRITE_ROOT")
	}
	if getLastChar(root) == '/' {
		os.Setenv("JK_WRITE_ROOT", root[:(len(root)-1)])
	}
	return root
}
