package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path) // permet d'ouvrir le fichier
	if err != nil {
		return nil, err
	}
	defer file.Close() // permet de fermer le fichier

	var lines []string
	scanner := bufio.NewScanner(file) // permet de scanner le fichier
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // permet de lire le fichier ligne par ligne
	}
	return lines, scanner.Err() // permet de retourner les lignes et les erreurs
}

func main() {

	lines, err := readLines("c:/Users/roden/OneDrive/Documents/cours b1/hangman/File.txt") // permet de lire le fichier et verifier si il y a une erreur
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Println("")
	fmt.Println(lines[0])            // affiche la première ligne du fichier
	fmt.Println(lines[len(lines)-1]) // affiche la dernière ligne du fichier

	for index, a := range lines { // permet de parcourir ma liste de mot et afficher le mot suivant before et d'aller a l'index 23
		if a == "before" {
			result, _ := strconv.Atoi(lines[index+1]) // permet de convertir le string en int
			fmt.Println(lines[result-1])
		}
		if a == "now " {
			are := lines[index-1]
			elem := int(are[1]) / len(lines) // permet de diviser le deuxieme caractère (en rune) du mot par le nombre de mot dans le fichier
			fmt.Println(lines[elem-1])
		}
	}
	fmt.Println(rand.Intn(200)) // permet de générer un nombre aléatoire entre 0 et 200
}
