package main

import (
	"bufio"
	"fmt"
	"os"
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
