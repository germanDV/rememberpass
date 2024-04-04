package rememberpass

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/germandv/rememberpass/internal/creds"
	"github.com/germandv/rememberpass/internal/store"
)

type RememberPass struct {
	w           io.Writer
	r           io.Reader
	prompt      string
	storer      store.Storer
	credentials []creds.Credential
}

func New(w io.Writer, r io.Reader, fpath string) *RememberPass {
	return &RememberPass{
		w:      w,
		r:      r,
		storer: store.New(fpath),
		prompt: "> ",
	}
}

func (r *RememberPass) Start() {
	r.load()
	r.welcome()
	r.options()
}

func (r *RememberPass) load() {
	lines, err := r.storer.Read()
	if err != nil {
		fmt.Println(err)
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
	}
	r.credentials = creds.Parse(lines)
}

func (r *RememberPass) welcome() {
	fmt.Fprintf(r.w, "Welcome to RememberPass\n\n")
}

func (r *RememberPass) goodbye() {
	fmt.Fprintf(r.w, "Gracias! Vuelva pronto.\n")
	os.Exit(0)
}

func (r *RememberPass) options() {
	fmt.Fprintf(r.w, "What would you like to to do?\n")
	fmt.Fprintln(r.w, "\ttype '(a)dd' to add a new secret")
	fmt.Fprintln(r.w, "\ttype '(p)ractice' to practice your secrets")
	fmt.Fprintln(r.w, "\ttype '(l)ist' to list the ID of your secrets")
	fmt.Fprintln(r.w, "\ttype '(q)uit' to exit")
	fmt.Fprintln(r.w, "")

	answer := ""
	fmt.Fprint(r.w, r.prompt)
	fmt.Fscanln(r.r, &answer)

	switch answer {
	case "a", "add":
		r.add()
		r.options()
	case "l", "list":
		r.list()
		r.options()
	case "p", "practice":
		r.practice()
		r.goodbye()
	case "q", "quit":
		r.goodbye()
	default:
		fmt.Fprintln(r.w, "I don't understand")
		r.options()
	}
}
