# RememberPass

A CLI tool to help you remember the passwords you should not forget.
Like the master password to unlock your password manager.

Passwords are hashed using [Argon2](https://pkg.go.dev/golang.org/x/crypto/argon2).
Which means the password is not stored in plain text and there's no way of retrieving it, so this will not work as a password manager.
