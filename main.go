package main

import "fmt"

type Point struct {
	p bool
	i int
}

func ccw_add_faces(left, right, topLeft, topRight Point) (res []int) {
	if left.p && right.p && topRight.p {
		res = append(res, left.i, right.i, topRight.i)
	}
	if left.p && topRight.p && topLeft.p {
		res = append(res, left.i, topRight.i, topLeft.i)
	}

	if !left.p && right.p && topLeft.p && topRight.p {
		res = append(res, topLeft.i, right.i, topRight.i)
	}

	if !topRight.p && left.p && right.p && topLeft.p {
		res = append(res, left.i, right.i, topLeft.i)
	}

	return res
}
func ccw_checker_pattern(checkerLeft, checkerRight, checkerTopLeft, checkerTopRight Point, badCheckerTopLeft, badCheckerTopRight, badCheckerLeft, badCheckerRight Point) (res []int) {

	checker1True := checkerLeft.p && checkerRight.p && checkerTopLeft.p && checkerTopRight.p
	checker2True := badCheckerLeft.p && badCheckerRight.p && badCheckerTopLeft.p && badCheckerTopRight.p

	checker1AllFalse := checkerLeft.p == false && checkerRight.p == false && checkerTopLeft.p == false && checkerTopRight.p == false
	checker2AllFalse := badCheckerLeft.p == false && badCheckerRight.p == false && badCheckerTopLeft.p == false && badCheckerTopRight.p == false

	if checker1True && checker2AllFalse {
		res = append(res, checkerLeft.i, checkerRight.i, checkerTopRight.i)
		res = append(res, checkerLeft.i, checkerTopRight.i, checkerTopLeft.i)
	}

	if checker2True && checker1AllFalse {
		res = append(res, badCheckerLeft.i, badCheckerRight.i, badCheckerTopRight.i)
		res = append(res, badCheckerLeft.i, badCheckerTopRight.i, badCheckerTopLeft.i)
	}
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
		if has0 && has2 && has5 && has7 {
			fmt.Println("here")
		}

		p0 := Point{p: has0, i: 0}
		p1 := Point{p: has1, i: 1}
		p2 := Point{p: has2, i: 2}
		p3 := Point{p: has3, i: 3}

		p4 := Point{p: has4, i: 4}
		p5 := Point{p: has5, i: 5}
		p6 := Point{p: has6, i: 6}
		p7 := Point{p: has7, i: 7}

		indices := []int{}

		indices = append(indices, ccw_add_faces(p0, p2, p4, p6)...)
		indices = append(indices, ccw_add_faces(p1, p0, p5, p4)...)
		indices = append(indices, ccw_add_faces(p2, p3, p6, p7)...)
		indices = append(indices, ccw_add_faces(p4, p6, p5, p7)...)
		indices = append(indices, ccw_add_faces(p3, p1, p7, p5)...)
		indices = append(indices, ccw_add_faces(p3, p1, p2, p0)...)

		indices = append(indices, ccw_add_edges(p2, p6, p0, p3, p4, p7, p1, p5)...)
		indices = append(indices, ccw_add_edges(p0, p4, p1, p2, p5, p6, p3, p7)...)
		indices = append(indices, ccw_add_edges(p3, p7, p2, p1, p6, p5, p0, p4)...)
		indices = append(indices, ccw_add_edges(p1, p5, p3, p0, p7, p4, p2, p6)...)

		indices = append(indices, ccw_checker_pattern(p0, p2, p5, p7, p4, p6, p1, p3)...)
		indices = append(indices, ccw_checker_pattern(p1, p0, p7, p7, p5, p4, p3, p2)...)

		// indices = append(indices, ccw_checker_pattern(p2, p3, p4, p5, p6, p7, p0, p1)...)
		// indices = append(indices, ccw_checker_pattern(p3, p1, p6, p4, p7, p5, p2, p0)...)

		finalIndices[i] = indices

	}
	return finalIndices
}
func main() {
	custom_marching_cubes()
}
