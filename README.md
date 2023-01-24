# Password Manager
This is a simple and safe cli password manager written in go.

# Installation
```bash
git clone https://github.com/NotVasil/Password-Manager
cd Password-Manager
go build -o pwdm cmd/main/main.go
```

# Usage
```bash
./pwdm -new <website> <your-password> : Creates random password for a website and encrypts it using the entered passwords hash.
./pwdm -show <your-password>: Shows passwords with their respective id and website.
./pwdm -remove <password-id>: Deletes password by ID.
```

## License
[MIT](https://choosealicense.com/licenses/mit/)
