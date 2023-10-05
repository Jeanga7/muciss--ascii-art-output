package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var nom_fichier string
	if len(os.Args) == 4 {
		nom_fichier = Filename(nom_fichier)
		// Vérifie si l'argument passé en ligne de commande n'est pas imprimable.
		if !IsPrintable(os.Args[2]) {
			fmt.Println("Non-Printable Element Detected?")
			return
		}
		lines := Open_File(os.Args[3])
		if len(lines) != 855 {
			return
		}

		// Vérifie si la chaîne passée en argument n'est pas vide.
		if len(os.Args[2]) > 0 {
			t := Finder(os.Args[2])

			// Affiche le contenu du fichier en utilisant les indices générés par Finder.
			Redirect(Display(lines, t), nom_fichier)
		}
	} else if len(os.Args) == 3 {
		// Verify if the first argument doesn't begin with --output=
		//we are testing
		if !strings.HasPrefix(os.Args[1], "--output=") {
			if !IsPrintable(os.Args[1]) {
				fmt.Println("Non-Printable Element Detected?")
				return
			}

			lines := Open_File(os.Args[2])

			if len(lines) != 855 {
				return
			}
			// Vérifie si la chaîne passée en argument n'est pas vide.
			if len(os.Args[1]) > 0 {
				t := Finder(os.Args[1])
				// Affiche le contenu du fichier en utilisant les indices générés par Finder.
				fmt.Println(Display(lines, t))
			}
		}

		if ShouldReturn(nom_fichier) {
			return
		}
	} else if len(os.Args) == 2 {
		lines := Open_File("standard")

		if !IsPrintable(os.Args[1]) {
			fmt.Println("Non-Printable Element Detected?")
			return
		}

		if len(os.Args[1]) > 0 {
			t := Finder(os.Args[1])
			// Affiche le contenu du fichier en utilisant les indices générés par Finder.
			fmt.Println(Display(lines, t))
		}
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println()
		fmt.Println("Example: go run . --output=<fileName.txt> something standard")
	}

}
