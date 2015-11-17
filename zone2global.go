package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/scottdware/go-junos"
	"os"
	"strings"
)

var (
	srx    string
	u      string
	p      string
	commit bool
)

func init() {
	flag.Usage = func() {
		fmt.Println("zone2global - Convert an SRX from a zone-based address book to a global one.")
		fmt.Println("\nUsage: zone2global [OPTIONS]")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.StringVar(&srx, "srx", "", "SRX to run the conversion against. If specifying multiple, enclose in quotes, i.e. \"srx240-1 srx1400-2\"")
	flag.StringVar(&u, "u", "", "Username")
	flag.StringVar(&p, "p", "", "Password")
	flag.BoolVar(&commit, "commit", false, "Choose to apply the configuration directly instead of creating a file.")
	flag.Parse()
}

func main() {
	for _, s := range strings.Split(srx, " ") {
		jnpr, err := junos.NewSession(s, u, p)
		if err != nil {
			fmt.Println(err)
		}
		defer jnpr.Close()

		config := jnpr.ConvertAddressBook()

		if commit {
			jnpr.Config(config, "set", true)
		}

		if !commit {
			f, err := os.Create(fmt.Sprintf("%s_globaladdrbook.txt", s))
			if err != nil {
				fmt.Println(err)
			}
			defer f.Close()

			b := bufio.NewWriter(f)

			for _, cmd := range config {
				b.Write([]byte(cmd))
			}

			b.Flush()
		}
	}
}
