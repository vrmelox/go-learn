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
	fmt.Println("The velocity of the ball is :")
	fmt.Printf("(%.2f, %.2f, %.2f)\n", velo.x, velo.y, velo.z)
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
			return []float64{}, fmt.Errorf("Impossible de parser le %d argument en float64", i + 1)
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
	fmt.Printf("At time + %.0f, ball coordinates will be:\n", f[6])
	fmt.Printf("(%.2f, %.2f, %.2f)\n", result.x, result.y, result.z)
		velo := vectorDif(Vector{f[0], f[1], f[2]}, Vector{f[3], f[4], f[5]})
	if result.z == 0.0 || f[6] < 0.0 && result.z > 0 || f[6] > 0 && result.z < 0 {
		angle, err := hitAngle(velo)
		if err == nil {
			fmt.Printf("The incidence angle is : \n%.2f\n", angle)
			return
		}
	}
	fmt.Println("The ball won't reach the paddle")
}