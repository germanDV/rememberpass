package main

import (
	"os"
	"path"

	"github.com/germandv/rememberpass"
	"github.com/germandv/rememberpass/internal/homedir"
)

func main() {
	w := os.Stdout
	r := os.Stdin

	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	fpath := path.Join(home, ".rememberpass_targets.txt")
	rp := rememberpass.New(w, r, fpath)
	rp.Start()
}
