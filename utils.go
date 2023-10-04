package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// IsPrintable vérifie si tous les caractères de la chaîne `s` sont imprimables (c'est-à-dire ASCII entre 32 et 127).
func IsPrintable(s string) bool {
	// Itération à travers les runes de la chaîne.
	for _, v := range s {
		// Vérifie si le code ASCII de la rune est en dehors de la plage imprimable.
		if v < 32 || v > 127 {
			return false
		}
	}
	return true
}

// Open_File ouvre le fichier de bannière correspondant à la chaîne `s`.
// Il retourne un tableau de lignes de texte lues depuis le fichier.
func Open_File(s string) (t []string) {
	// Ouvre le fichier correspondant à la chaîne `s`.
	file, err := os.Open("./banner/" + s + ".txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	defer file.Close()

	// Crée un scanner pour lire le fichier ligne par ligne.
	scn := bufio.NewScanner(file)

	// Itère à travers les lignes du fichier.
	for scn.Scan() {
		// Ajoute chaque ligne au tableau `t`.
		t = append(t, scn.Text())
	}

	return
}

// this function returns True if the all the Characters are printable
//
//	redirect them to the name of our file that we put as a parameter
func ShouldReturn(nom_fichier string) bool {
	if strings.HasPrefix(os.Args[1], "--output=") && strings.HasSuffix(os.Args[1], ".txt") {
		if len(os.Args[1]) > 13 {
			for i := 9; i < len(os.Args[1]); i++ {
				nom_fichier = nom_fichier + string(os.Args[1][i])
			}
		}
		if !IsPrintable(os.Args[1]) {
			fmt.Println("Non-Printable Element Detected?")
			return true
		}

		lines := Open_File("standard")

		if len(os.Args[2]) > 0 {
			t := Finder(os.Args[2])

			Redirect(Display(lines, t), nom_fichier)
		}
	}
	return false
}

// This function returns the name of the file where we are going put the stdout
func Filename(nom_fichier string) string {
	if strings.HasPrefix(os.Args[1], "--output=") && strings.HasSuffix(os.Args[1], ".txt") {
		if len(os.Args[1]) > 13 {
			for i := 9; i < len(os.Args[1]); i++ {
				nom_fichier = nom_fichier + string(os.Args[1][i])
			}
		}
	}
	return nom_fichier
}
