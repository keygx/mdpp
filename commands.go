package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandSearch,
}

var commandSearch = cli.Command{
	Name:        "search",
	Usage:       "mdpp search [Provisioning Profiles Name]",
	Description: ``,
	Action:      doSearch,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doSearch(c *cli.Context) {

	if len(c.Args()) < 1 {
		return
	}

	keyword := c.Args()[0]

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	var path = usr.HomeDir + "/Library/MobileDevice/Provisioning Profiles/"

	list, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	var provisions []string
	for _, finfo := range list {
		if finfo.IsDir() || -1 == strings.Index(finfo.Name(), ".mobileprovision") {
			continue
		}
		provisions = append(provisions, finfo.Name())
	}

	for _, fileName := range provisions {
		cmd := exec.Command("egrep", "-a", "-A", "2", "<key>Name</key>", path+fileName)

		stdout, err := cmd.StdoutPipe()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cmd.Start()

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), "<string>"+keyword+"</string>") {
				fmt.Printf("%s%s\n", path, fileName)
				//fmt.Println(scanner.Text())
			}
		}

		cmd.Wait()
	}
}
