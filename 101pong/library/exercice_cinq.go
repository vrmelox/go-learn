package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

func positionAtn(x0, y0, z0, x1, y1, z1, n float64) Vector {
	velo := vectorDif(Vector{x0, y0, z0}, Vector{x1, y1, z1})
	vecoef := coefVector(velo, n)
	return vectorSum(vecoef, Vector{x1, y1, z1})
}

func checkRobustness() ([]float64, error) {
	if len(os.Args) != 8 {
		return []float64{}, errors.New("Nombre d'arguments différent")
	}
	args := make([]float64, 0, 7)
	for i, arg := range os.Args[1:] {
		f, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return []float64{}, errors.New(fmt.Sprintf("Impossible de parser le %d argument en float64", i + 1))
		} else {
			args = append(args, f);
		}
	}
	return args, nil
}
func main() {

	f, err := checkRobustness()

	if err != nil {
		fmt.Println(err)
		return
	}
	result := positionAtn(f[0], f[1], f[2], f[3], f[4], f[5], f[6])
	fmt.Println(result)
}