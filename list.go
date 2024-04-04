package rememberpass

import (
	"fmt"
)

func (r *RememberPass) list() {
	if len(r.credentials) == 0 {
		fmt.Fprintln(r.w, "You have no secrets to remember")
	}

	for i, credential := range r.credentials {
		fmt.Fprintf(r.w, "\t%d. %s\n", i+1, credential.ID)
	}

	fmt.Fprintln(r.w)
}
