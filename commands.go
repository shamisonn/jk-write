package main

import (
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
	Usage:       "",
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
	return nil
}
func doNew(c *cli.Context) error {
	return nil
}
func doUpdate(c *cli.Context) error {
	return nil
}

func doRoot(c *cli.Context) error {
	return nil
}

func getRoot() string {
	root := os.Getenv("GH_WRITE_ROOT")
	if root == "" {
		log.Fatal("set GH_WRITE_ROOT")
	}
	return root
}
