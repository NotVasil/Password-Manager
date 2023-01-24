package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"log"

	db "github.com/vuk-v/Password-Manager/pkg/db"
	pwd "github.com/vuk-v/Password-Manager/pkg/password"
)

func main() {
	db.CreateTable()

	h := sha256.New()
	var key []byte

	remove := flag.Int("remove", 0, "Remove certain password by ID")
	show := flag.Bool("show", false, "Show current passwords, enter your password after the command.")
	new := flag.String("new", "", "Create new password and assign it to a website, enter your password after the website.\nExample of usage: pwdm -new website.com password")

	flag.Parse()
	if *new != "" {
		tail := flag.Args()

		h.Write([]byte(tail[0]))
		key = h.Sum(nil)

		websiteexists := fmt.Sprintf("SELECT * FROM passwords WHERE website='%v';", pwd.Encrypt(*new, string(key)))
		rows, err := db.GetDataBase().Query(websiteexists)
		if err != nil {
			log.Fatal(err)
		}

		if !rows.Next() {
			statement, err := db.GetDataBase().Prepare("INSERT INTO passwords (password, website) values(?, ?)")
			if err != nil {
				fmt.Println("There was an error creating your password, please rethink your input.")
				return
			}

			statement.Exec(pwd.Encrypt(pwd.RandomPassword(20), string(key)), pwd.Encrypt(*new, string(key)))
			fmt.Println("New password created and stored.")

			return
		}

		fmt.Println("There is already a password for that website. If you want to set a new one, use the remove command.")
		return
	}

	if *show {
		tail := flag.Args()

		if len(tail) != 0 {
			h.Write([]byte(tail[0]))
			key = h.Sum(nil)

			rows, err := db.GetDataBase().Query("SELECT * FROM passwords")
			if err != nil {
				fmt.Println("There was an error finding passwords, please rethink your input.")
				return
			}

			var pid int
			var password string
			var website string

			for rows.Next() {
				err = rows.Scan(&pid, &password, &website)
				if err != nil {
					fmt.Println("Wrong password! Try again.")
					return
				}

				fmt.Println(pid, pwd.Decrypt(password, string(key)), " - ", pwd.Decrypt(website, string(key)))
			}

			rows.Close()
			return
		}

		fmt.Println("Enter the password you generated with.")
		return
	}

	if *remove != 0 {
		statementtxt := fmt.Sprintf("DELETE FROM passwords WHERE pid = (?)")
		statement, err := db.GetDataBase().Prepare(statementtxt)

		if err != nil {
			fmt.Println("There was an error removing your password, please rethink your input.")
			return
		}

		statement.Exec(remove)
		fmt.Println("Password successfully deleted.")

		return
	}

	db.GetDataBase().Close()
}
