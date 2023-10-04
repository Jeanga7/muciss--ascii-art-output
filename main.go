package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Vérifie si le nombre d'arguments est entre 2 et 3 inclus.
	var nom_fichier string
	if len(os.Args) == 4 {
		if strings.HasPrefix(os.Args[1],"--output=") && strings.HasSuffix(os.Args[1],".txt"){
			if len(os.Args[1]) > 13{
				for i:=9; i<len(os.Args[1]);i++{
					nom_fichier = nom_fichier + string(os.Args[1][i])
				}
			}
		}
		// Vérifie si l'argument passé en ligne de commande n'est pas imprimable.
		if !IsPrintable(os.Args[2]) {
			fmt.Println("Non-Printable Element Detected?")
			return
		}

		// Ouvre le fichier "standard" par défaut.
		lines := Open_File(os.Args[3])

		// if len(lines) != 855 {
		// 	return
		// }

		// Vérifie si la chaîne passée en argument n'est pas vide.
		if len(os.Args[2]) > 0 {
			t := Finder(os.Args[2])

			// Affiche le contenu du fichier en utilisant les indices générés par Finder.
			Redirect(Display(lines, t),nom_fichier)
		}
	} else if len(os.Args) >= 2 && len(os.Args) <= 3 {
		// Vérifie si l'argument passé en ligne de commande n'est pas imprimable.
		if !strings.HasPrefix(os.Args[1], "--output="){
			if !IsPrintable(os.Args[1]) {
				fmt.Println("Non-Printable Element Detected?")
				return
			}
		// Ouvre le fichier "standard" par défaut.
		lines := Open_File("standard")
		// Si un deuxième argument est fourni en ligne de commande, ouvre le fichier correspondant.
		if len(os.Args) == 3 {
			lines = Open_File(os.Args[2])
		}
		// Vérifie si la longueur du fichier est de 855 lignes.
		if len(lines) != 855 {
			return
		}
		// Vérifie si la chaîne passée en argument n'est pas vide.
		if len(os.Args[1]) > 0 {
			t := Finder(os.Args[1])
			// Affiche le contenu du fichier en utilisant les indices générés par Finder.
			fmt.Println(Display(lines, t))
		}
	} else if strings.HasPrefix(os.Args[1],"--output=") && strings.HasSuffix(os.Args[1],".txt"){
		if len(os.Args[1]) > 13{
			for i:=9; i<len(os.Args[1]);i++{
				nom_fichier = nom_fichier + string(os.Args[1][i])
			}
		}
		if !IsPrintable(os.Args[1]) {
			fmt.Println("Non-Printable Element Detected?")
			return
		}

		lines := Open_File("standard")

		if len(os.Args[2]) > 0 {
			t := Finder(os.Args[2])
			// Affiche le contenu du fichier en utilisant les indices générés par Finder.
			Redirect(Display(lines, t),nom_fichier)
		}
	}
	}
}

// Usage: go run . [STRING] 
// Usage: go run . [OPTION] [STRING]
// Usage: go run . [STRING] [BANNER] 
// Usage: go run . [OPTION] [STRING] [BANNER] 

