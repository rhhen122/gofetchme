package main

import "fmt"

// User CONFIG

var (
	// CONFIG
	// See "os.txt" for list of supported os's
	// See de.txt for list of supported desktop enviroments

	OS = "mac"
	DE = "gnome"
)

var (
	// Important Values (Dont Touch)
	reset string = "\033[0m"

	// Colors
	blue     string = "\033[36m"
	acsiicol string = "\033[96m"
	col1     string = "\033[31m" // Editable
	col2     string = "\033[35m" // Editable

	// Ascii Art
	asciiln3 string = "  **  "
	asciiln2 string = " **** "
	asciiln1 string = "**  **"

	// Extra Values
	dead string = "		"
)

// Script

func checkforsupport() {
	// Check for MAC
	if OS == "mac" {
		OS = "The Macintosh Operating System"
		DE = "Macintosh Desktop"

		// Ascii Art
		asciiln3 = "  ^^  "
		asciiln2 = " ^--^ "
		asciiln1 = "^    ^"
	}
	// Check for ARCH
	if OS == "arch" {
		OS = "The Arch Linux Operating System"
		if DE == "wayland" {
			DE = "The Wayland Desktop Enviroment"
		}
		if DE == "gnome" {
			DE = "The GNOME GUI / Desktop Enviroment"
		}
		if DE == "kde" {
			DE = "The KDE Plasma Desktop Enviroment"
		}
	}
	// Check for LINUX
	if OS == "linux" {
		OS = "GNU + Linux"
		if DE == "wayland" {
			OS = "The Arch Linux Operating System"
			DE = "The Wayland Desktop Enviroment"
		}
		if DE == "gnome" {
			DE = "The GNOME GUI / Desktop Enviroment"
		}
		if DE == "kde" {
			DE = "The KDE Plasma Desktop Enviroment"
		}
		asciiln3 = "&     "
		asciiln2 = "&     "
		asciiln1 = "&*&*&*"
	}
	// Check for FreeBSD
	if OS == "freebsd" {
		OS = "The FreeBSD Operating System"
		asciiln3 = " ^  ^ "
		asciiln2 = "(    )"
		asciiln1 = " {  } "
	}
	// Check for Fedora Linux
	if OS == "fedora" {
		OS = "The Fedora Distro(Bution) of GNU / Linux"
		if DE == "gnome" {
			DE = "The GNOME GUI / Desktop Enviroment"
		}
		if DE == "kde" {
			DE = "The KDE Plasma Desktop Enviroment"
		}
		asciiln3 = "  *.  "
		asciiln2 = " / |  "
		asciiln1 = "----- "
	}
}

func main() {
	checkforsupport()
	fmt.Print("")

	//
	fmt.Print(acsiicol + asciiln3 + reset + dead)
	fmt.Println(blue + "My OS is, " + reset + col1 + OS + reset)
	fmt.Println(acsiicol + asciiln2 + reset + dead)
	fmt.Print(acsiicol + asciiln1 + reset + dead)
	fmt.Println(blue + "My DE is, " + reset + col2 + DE + reset)
}
