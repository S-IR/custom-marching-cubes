package main

import (
	"fmt"
)

type Point struct {
	p bool
	i int
}

func ccw_add_faces(left, right, topLeft, topRight Point) (res []int) {
	// if !left.p || !right.p || !topLeft.p || !topRight.p {
	// 	return
	// }
	if left.p && right.p && topRight.p {
		res = append(res, left.i, right.i, topRight.i)
	}
	if left.p && topRight.p && topLeft.p {
		res = append(res, left.i, topRight.i, topLeft.i)
	}

	// if !left.p && right.p && topLeft.p && topRight.p {
	// 	res = append(res, topLeft.i, right.i, topRight.i)
	// }

	// if !topRight.p && left.p && right.p && topLeft.p {
	// 	res = append(res, left.i, right.i, topLeft.i)
	// }

	return res
}

func ccw_checker_pattern(checkerLeft, checkerRight, checkerTopLeft, checkerTopRight Point, badCheckerBottomLeft, badCheckerBottomRight, badCheckerTopLeft, badCheckerTopRight Point) (res []int) {

	checkerTrue := checkerLeft.p && checkerRight.p && checkerTopLeft.p && checkerTopRight.p
	// checker2True := badCheckerLeft.p && badCheckerRight.p && badCheckerTopLeft.p && badCheckerTopRight.p
	badCheckerFalse := (badCheckerBottomLeft.p == false && badCheckerBottomRight.p == false)
	badCheckerFalse = badCheckerFalse || (badCheckerTopLeft.p == false && badCheckerTopRight.p == false)
	// checker1AllFalse := checkerLeft.p == false && checkerRight.p == false && checkerTopLeft.p == false && checkerTopRight.p == false
	// checker2AllFalse := badCheckerLeft.p == false && badCheckerRight.p == false && badCheckerTopLeft.p == false && badCheckerTopRight.p == false

	if checkerTrue && badCheckerFalse {
		res = append(res, ccw_add_faces(checkerLeft, checkerRight, checkerTopLeft, checkerTopRight)...)
		// res = append(res, checkerLeft.i, checkerTopRight.i, checkerTopLeft.i)
	}

	// if checker2True && checker1AllFalse {
	// 	res = append(res, ccw_add_faces(badCheckerLeft, badCheckerRight, badCheckerTopLeft, badCheckerTopRight)...)
	// 	// res = append(res, badCheckerLeft.i, badCheckerTopRight.i, badCheckerTopLeft.i)
	// }
	return res
}

func ccw_add_edges(sideBottom, sideTop Point, otherDiagonalLeft, otherDiagonalRight, otherDiagonalTopLeft, otherDiagonalTopRight, otherSideBottom, otherSideTop Point) (res []int) {
	if sideBottom.p && sideTop.p {
		return res
	}
	if !sideBottom.p && !sideTop.p {
		res = append(res, ccw_add_faces(otherDiagonalLeft, otherDiagonalRight, otherDiagonalTopLeft, otherDiagonalTopRight)...)
		return res
	}
	if !sideTop.p && sideBottom.p && otherDiagonalRight.p && otherDiagonalLeft.p {
		res = append(res, ccw_add_faces(otherDiagonalTopLeft, sideBottom, otherSideTop, otherDiagonalTopRight)...)
		return res
	}

	if sideTop.p && !sideBottom.p && otherDiagonalTopRight.p && otherDiagonalTopLeft.p {
		res = append(res, ccw_add_faces(otherDiagonalLeft, sideTop, otherSideBottom, otherDiagonalRight)...)
		return res
	}

	return res

}

var edges = [6][2]int{

	{0, 2},
	{1, 3},
	{0, 1},
	{2, 3},
	{0, 4},
	{1, 5},
}

var aboveEdges = [12][2]int{
	{4, 6},
	{5, 7},
	{4, 5},
	{6, 7},
	{3, 7},
	{2, 6},

	{1, 3},
	{0, 2},
	{2, 3},
	{0, 1},
	{2, 4},
	{1, 5},
}

var oppositeEdges = [6][2]int{
	{5, 7},
	{4, 6},
	{6, 7},
	{4, 5},
	{3, 7},
	{2, 6},
}

func custom_marching_cubes() [256][]int {
	finalIndices := [256][]int{}
	was0 := false
	for i := uint8(0); i <= 255; i++ {
		if i == 0 {
			if was0 {
				break
			}
			was0 = true
		}
		has0 := i&(1<<0) != 0
		has1 := i&(1<<1) != 0
		has2 := i&(1<<2) != 0
		has3 := i&(1<<3) != 0
		has4 := i&(1<<4) != 0
		has5 := i&(1<<5) != 0
		has6 := i&(1<<6) != 0
		has7 := i&(1<<7) != 0

		p0 := Point{p: has0, i: 0}
		p1 := Point{p: has1, i: 1}
		p2 := Point{p: has2, i: 2}
		p3 := Point{p: has3, i: 3}

		p4 := Point{p: has4, i: 4}
		p5 := Point{p: has5, i: 5}
		p6 := Point{p: has6, i: 6}
		p7 := Point{p: has7, i: 7}

		// ps := []Point{p0, p1, p2, p3, p4, p5, p6, p7}
		indices := []int{}
		if i == 175 {
			// fmt.Println("front:", ccw_add_faces(p0, p2, p4, p6))
			// fmt.Println("back:", ccw_add_faces(p3, p1, p7, p5))
			// fmt.Println("left:", ccw_add_faces(p1, p0, p5, p4))
			// fmt.Println("right:", ccw_add_faces(p2, p3, p6, p7))
			// fmt.Println("bottom:", ccw_add_faces(p1, p3, p0, p2))
			// fmt.Println("top:", ccw_add_faces(p4, p6, p5, p7))
			// fmt.Println("edge1:", ccw_add_edges(p2, p6, p0, p3, p4, p7, p1, p5))
			// fmt.Println("edge2:", ccw_add_edges(p0, p4, p1, p2, p5, p6, p3, p7))
			// fmt.Println("edge3:", ccw_add_edges(p3, p7, p2, p1, p6, p5, p0, p4))
			// fmt.Println("edge4:", ccw_add_edges(p1, p5, p3, p0, p7, p4, p2, p6))
			// fmt.Println("checker1:", ccw_checker_pattern(p0, p2, p5, p7, p4, p6, p1, p3))
			// fmt.Println("checker2:", ccw_checker_pattern(p2, p3, p4, p5, p6, p7, p0, p1))

			// fmt.Println("checker2:", ccw_checker_pattern(p1, p0, p7, p7, p5, p4, p3, p2))
		}
		indices = append(indices, ccw_add_faces(p0, p2, p4, p6)...)
		indices = append(indices, ccw_add_faces(p3, p1, p7, p5)...)

		indices = append(indices, ccw_add_faces(p1, p0, p5, p4)...)
		indices = append(indices, ccw_add_faces(p2, p3, p6, p7)...)

		indices = append(indices, ccw_add_faces(p0, p2, p1, p3)...)
		indices = append(indices, ccw_add_faces(p4, p6, p5, p7)...)

		finalIndices[i] = indices

	}
	// maxLen := 36

	return finalIndices
}
func main() {
	finalIndices := custom_marching_cubes()
	fmt.Println("POINTS_TO_TRIANGLES_CONVERTER:=[256][]i32{")

	for _, p := range finalIndices {
		fmt.Print("{")
		totalPrinted := 0
		for _, indices := range p {
			if totalPrinted == (len(p) - 1) {
				fmt.Printf("%d", indices)
			} else {
				fmt.Printf("%d,", indices)

			}
			totalPrinted += 1

		}
		// diff := maxLen - len(p)
		// for j := 0; j < diff; j += 1 {
		// 	if totalPrinted == (maxLen - 1) {
		// 		fmt.Printf("%d", -1)
		// 	} else {
		// 		fmt.Printf("%d,", -1)
		// 	}
		// 	totalPrinted += 1

		// }
		fmt.Println("},")

	}
	fmt.Println("};")

}
func intAbs(v int) int {
	if v < 0 {
		v *= -1
	}
	return v
}
func getP(val int, ps []Point) Point {
	for _, p := range ps {
		if p.i == val {
			return p
		}
	}
	panic(fmt.Sprintf("did not find p for val %v , ps: %v", val, ps))
}
func pointsToUint8(points ...int) uint8 {
	var result uint8
	for _, p := range points {
		if p >= 0 && p <= 7 {
			result |= 1 << p
		}
	}
	return result
}
