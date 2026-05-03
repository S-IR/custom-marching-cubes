package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// pointsToUint8 takes a slice of ints [0..7] and returns a uint8 with those bits set

func TestMarchingCubes(t *testing.T) {
	indices := custom_marching_cubes()
	var FRONT_FACE = [6]int{0, 2, 6, 0, 6, 4}
	var TOP_FACE = [6]int{4, 6, 7, 4, 7, 5}
	var BACK_FACE = [6]int{3, 1, 5, 3, 5, 7}

	var LEFT_FACE = [6]int{
		1, 0, 4, 1, 4, 5,
	}
	var RIGHT_FACE = [6]int{
		2, 3, 7, 2, 7, 6,
	}

	var BOTTOM_FACE = [6]int{
		0, 2, 3, 0, 3, 1,
	}
	var DIAGONAL_BL_TR = [6]int{
		3, 0, 4, 3, 4, 7,
	}
	type TestStruct struct {
		index           uint8
		indicesExpected []int
	}

	structs := []TestStruct{
		TestStruct{index: pointsToUint8(), indicesExpected: []int{}},
		TestStruct{index: pointsToUint8(0), indicesExpected: []int{}},
		TestStruct{index: pointsToUint8(7), indicesExpected: []int{}},

		TestStruct{index: pointsToUint8(0, 2, 4, 6), indicesExpected: FRONT_FACE[:]},
		TestStruct{index: pointsToUint8(4, 5, 6, 7), indicesExpected: TOP_FACE[:]},
		TestStruct{index: pointsToUint8(2, 3, 4, 5), indicesExpected: []int{4, 5, 3, 4, 3, 2}},

		TestStruct{index: pointsToUint8(0, 1, 2, 3, 5, 7), indicesExpected: concat(BOTTOM_FACE[:], BACK_FACE[:], []int{1, 0, 5, 2, 3, 7, 5, 7, 2, 5, 2, 0})},

		TestStruct{index: pointsToUint8(7, 6, 5, 4), indicesExpected: TOP_FACE[:]},
		// TestStruct{index: pointsToUint8(4, 6, 5, 7), indicesExpected: []int{4, 6, 7, 4, 7, 5}},

		TestStruct{index: pointsToUint8(1, 3, 5, 7), indicesExpected: BACK_FACE[:]},

		TestStruct{index: pointsToUint8(0, 4, 2, 6, 3, 7), indicesExpected: concat(
			TOP_FACE[:], RIGHT_FACE[:], DIAGONAL_BL_TR[:], []int{4, 6, 7, 3, 0, 2},
		)},

		TestStruct{index: pointsToUint8(0, 4, 2, 6, 3, 7, 1), indicesExpected: concat(
			TOP_FACE[:], RIGHT_FACE[:], []int{4, 6, 7, 3, 0, 2, 3, 1, 0, 1, 0, 4, 3, 1, 7, 1, 0, 4, 7, 4, 6},
		)},

		TestStruct{index: pointsToUint8(0, 2, 5, 7), indicesExpected: []int{0, 2, 7, 0, 5, 7}},

		TestStruct{
			index: pointsToUint8(0, 1, 2, 3, 4, 5, 6, 7),
			indicesExpected: concat(FRONT_FACE[:],
				BACK_FACE[:],
				RIGHT_FACE[:],
				LEFT_FACE[:],
				TOP_FACE[:],
				BOTTOM_FACE[:]),
		},

		// TestStruct{index: pointsToUint8(1, 3, 4, 6), indicesExpected: []int{}},

		// TestStruct{index: pointsToUint8(0,1,4,5,3,7,2), indicesExpected: []int{}},

	}
	// fmt.Println("pointsToUint8(0, 1, 2, 3, 5, 7),", pointsToUint8(0, 1, 2, 3, 5, 7))
	for i, myStruct := range structs {
		gottenIndices := indices[myStruct.index]

		debugMsg := fmt.Sprintf("gotten indices: %v vs myStruct at index %d %v", gottenIndices, i, myStruct)
		require.True(t, len(gottenIndices) == len(myStruct.indicesExpected), fmt.Sprintf("lengths do not match: gotten indices len %v vs myStruct len %v ", len(gottenIndices), len(myStruct.indicesExpected))+debugMsg)
		require.True(t, len(gottenIndices)%3 == 0)

		if i == 6 {
			fmt.Print("here")
		}
		for iFromIndicesExpected := 0; iFromIndicesExpected < len(myStruct.indicesExpected); iFromIndicesExpected += 3 {

			idx0 := myStruct.indicesExpected[iFromIndicesExpected]
			idx1 := myStruct.indicesExpected[iFromIndicesExpected+1]
			idx2 := myStruct.indicesExpected[iFromIndicesExpected+2]
			found := false
			for j := 0; j < len(gottenIndices); j += 3 {
				if gottenIndices[j] == idx0 && gottenIndices[j+1] == idx1 && gottenIndices[j+2] == idx2 {
					found = true
					break
				}
			}

			require.True(t, found, fmt.Sprintf("did not find %d %d %d in gotten indices ", idx0, idx1, idx2)+debugMsg)
		}

	}
}
func concat(faces ...[]int) []int {
	var result []int
	for _, f := range faces {
		result = append(result, f...)
	}
	return result
}
