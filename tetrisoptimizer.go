package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func main() {

	launch()
}

// Recupere en argument un fichier .txt lors de l'excution du script
func launch() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")

	} else if len(os.Args) == 2 {
		tetrisoptimizer()
	} else {
		fmt.Println("Too many arguments")
	}
}

func tetrisoptimizer() {

	size := math.Sqrt(8 * 4)
	round := int(math.Ceil(size))
	matrix := make([][]string, round)
	for i := range matrix {
		matrix[i] = make([]string, round)
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = "."
		}
	}

	code := 65
	var a = recupForme(15, code)
	var tab = remplir()

	test(a, tab)

}

// Fonction TEST, remplis la matrice final 6x6
//
// Prend en parametre une forme tetris en matrice et la tableau 6x6
func test(forme [4][4]string, tab [6][6][1]string) {
	var c = 0
	fmt.Println()
	fmt.Println(forme)
	fmt.Println()
	for i, res := range tab {
		if i != 4 && i != 5 {
			for j, res2 := range forme {
				for k, res3 := range res2 {
					res = res
					res3 = res3
					j = j
					if i == 0 || i == 1 || i == 2 || i == 3 && tab[i][c+k][0] != "." && string(res2[j]) == "#" {
						tab[i][c+k][0] = forme[i][0]
					}
				}
			}
		}
		fmt.Println(tab[i])
	}
}

// Recupere une forme tetris et l'envoie sous forme de matrice
func recupForme(endroit int, code int) [4][4]string {
	dat, err := ioutil.ReadFile("tetris.txt")
	if err != nil {
		fmt.Println(err)
	}

	a := strings.Split(string(dat), "\r\n")
	b := a[endroit]
	c := a[endroit+1]
	d := a[endroit+2]
	e := a[endroit+3]

	var mat = make([]string, 0)
	var res [4][4]string
	mat = append(mat, b)
	mat = append(mat, c)
	mat = append(mat, d)
	mat = append(mat, e)

	res = res

	for i, res1 := range mat {
		i = i
		for j, res2 := range res1 {
			j = j
			if res2 == 35 {
				res[i][0] += string(code)
			}
		}
	}
	return res
}

// Créer un tableau 6x6
func remplir() [6][6][1]string {
	var tableau [6][6][1]string

	for i, res := range tableau {
		for j, res2 := range res {
			res2 = res2
			tableau[i][j][0] = "."
		}
	}
	return tableau
}
