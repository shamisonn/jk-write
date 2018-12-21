package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
)

// Commands has all commands vars.
// A command has cmd's help text and behavior.
// Behavior function named "doFunction".
// Example: cmdList -> doList
var Commands = []cli.Command{
	cmdList,
	cmdNew,
	cmdUpdate,
	cmdRoot,
}

var cmdList = cli.Command{
	Name:        "list",
	Usage:       "show post list",
	Description: `Not yet`,
	Action:      doList,
	//	Flags: []cli.Flag{},
}
var cmdNew = cli.Command{
	Name:        "new",
	Usage:       "make a post",
	Description: `Not yet`,
	Action:      doNew,
	//	Flags: []cli.Flag{},
}
var cmdUpdate = cli.Command{
	Name:        "update",
	Usage:       "reopen old file",
	Description: `Not yet`,
	Action:      doUpdate,
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
	for _, file := range files {
		if getLastChar(file.Name()) != '~' {
			fmt.Println(file.Name())
		}
	}
	return nil
}

// this can't use for multibyte string
func getLastChar(s string) byte {
	return s[len(s)-1]
}

func doNew(c *cli.Context) error {
	return nil
}
func doUpdate(c *cli.Context) error {
	return nil
}

func doRoot(c *cli.Context) error {
	fmt.Println(getRoot())
	return nil
}

func getRoot() string {
	root := os.Getenv("GH_WRITE_ROOT")
	if root == "" {
		log.Fatal("You should set fullpath to GH_WRITE_ROOT")
	}
	return root
}
