package main

import (
	"fmt"
	"strings"
	"os"
	"bytes"
)

// Finder divise la chaîne de caractères `str` en sous-tableaux de valeurs.
// Chaque sous-tableau représente une ligne et chaque élément de la ligne est calculé en fonction des caractères de la chaîne.
func Finder(str string) (nn [][]int) {
	// Sépare la chaîne en lignes en utilisant "\\n" comme délimiteur.
	t := strings.Split(str, "\\n")

	start := 0

	// Vérifie s'il y a des lignes vides au début de la chaîne.
	if t[0] == "" && t[1] == "" {
		start = 1
	}

	var n []int

	// Itération à travers les lignes de la chaîne.
	for i := start; i < len(t); i++ {
		if t[i] != "" {
			// Itération à travers les caractères de la ligne.
			for _, v := range t[i] {
				// Calcule une valeur en fonction du caractère.
				n = append(n, (int(v-32) * 9))
			}

		} else {
			// Si la ligne est vide, ajoute -1 à la liste.
			n = append(n, -1)
		}

		// Ajoute la liste `n` à la liste de résultats `nn`.
		nn = append(nn, n)
		// Réinitialise `n` pour la prochaine itération.
		n = nil
	}

	return
}


// Display affiche les caractères correspondants aux indices donnés par `t`.
// Chaque sous-tableau dans `t` représente une ligne.
func Display(a []string, t [][]int) string {
	var output bytes.Buffer 
	for _, b := range t {
		if b[0] == -1 {
			output.WriteString("\n")
		} else {
			// Itération à travers les indices dans le sous-tableau.
			for i := 1; i < 9; i++ {
				// Itération à travers les indices du sous-tableau.
				for _, j := range b {
					// Écrivez le caractère correspondant à l'indice dans `a` dans le tampon.
					output.WriteString(a[j+i])
				}
				output.WriteString("\n")
			}
		}
	}
	return output.String() // Retourne la sortie sous forme de chaîne de caractères
}

func Redirect(s string, nom string) {
	file, errs := os.Create(nom)
	if errs != nil {
		fmt.Println("not a valid file name :", errs)
		return
	}
	defer file.Close()
	_, errs = file.WriteString(s)
	if errs != nil {
		fmt.Println("Failed to write to file :", errs)
		return
	}
}