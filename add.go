package rememberpass

import (
	"fmt"

	"github.com/germandv/rememberpass/internal/creds"
)

func (r *RememberPass) add() {
	fmt.Fprintln(r.w, "What's the ID/Name of the secret?")
	id := ""
	fmt.Fprint(r.w, r.prompt)
	fmt.Fscanln(r.r, &id)

	fmt.Fprintln(r.w, "What's the secret? (no worries, we'll save it as a crytographic hash)")
	secret := ""
	fmt.Fprint(r.w, r.prompt)
	fmt.Fscanln(r.r, &secret)

	credential, err := creds.New(id, secret)
	if err != nil {
		panic(err)
	}

	err = r.storer.Write(credential.String())
	if err != nil {
		panic(err)
	}

	r.credentials = append(r.credentials, credential)
	fmt.Fprintln(r.w, "Secret saved!")
	fmt.Fprintln(r.w)
}
