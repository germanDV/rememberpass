package rememberpass

import (
	"fmt"

	"github.com/germandv/rememberpass/internal/creds"
)

func (r *RememberPass) practice() {
	if len(r.credentials) == 0 {
		fmt.Fprintln(r.w, "You have no secrets to remember")
	}

	for _, credential := range r.credentials {
		r.ask(credential)
	}
}

func (r *RememberPass) ask(credential creds.Credential) {
	fmt.Fprintf(r.w, "Enter your secret for %q\n", credential.ID)

	ok := false
	for !ok {
		fmt.Fprint(r.w, r.prompt)
		candidate := ""
		fmt.Fscanln(r.r, &candidate)

		if credential.Compare(candidate) {
			ok = true
			fmt.Fprintf(r.w, "Good, that's correct.\n\n")
		} else {
			fmt.Fprintf(r.w, "Sorry, that's wrong. Try again\n")
		}
	}
}
