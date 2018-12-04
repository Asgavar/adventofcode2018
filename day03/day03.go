package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Claim struct {
	Id         int
	LeftOffset int
	TopOffset  int
	Width      int
	Height     int
}

type FabricPoint struct {
	Column int
	Row    int
}

func parseClaimFromString(claimAsString string) Claim {
	splitted := strings.Split(claimAsString, " ")
	id, _ := strconv.Atoi(splitted[0][1:])

	leftAndTopOffsets := strings.Split(splitted[2], ",")
	leftOffset, _ := strconv.Atoi(leftAndTopOffsets[0])
	topOffsetAsString := leftAndTopOffsets[1]
	topOffset, _ := strconv.Atoi(topOffsetAsString[:len(topOffsetAsString)-1])

	widthAndHeight := strings.Split(splitted[3], "x")
	width, _ := strconv.Atoi(widthAndHeight[0])
	height, _ := strconv.Atoi(widthAndHeight[1])

	return Claim{id, leftOffset, topOffset, width, height}
}

func pointsWithinClaim(claim Claim) []FabricPoint {
	var fabricPoints []FabricPoint

	for row_idx := claim.LeftOffset; row_idx < claim.LeftOffset+claim.Width; row_idx++ {
		for col_idx := claim.TopOffset; col_idx < claim.TopOffset+claim.Height; col_idx++ {
			fabricPoints = append(fabricPoints, FabricPoint{Column: col_idx, Row: row_idx})
		}
	}

	return fabricPoints
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	var claims []Claim

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		claims = append(claims, parseClaimFromString(scanner.Text()))
	}

	fabric := make(map[FabricPoint][]Claim)

	for _, claim := range claims {
		pointsTakenByThisClaim := pointsWithinClaim(claim)
		for _, point := range pointsTakenByThisClaim {
			fabric[point] = append(fabric[point], claim)
		}
	}

	occupiedByMoreThanOneCount := 0
	for _, claimsOccupying := range fabric {
		if len(claimsOccupying) > 1 {
			occupiedByMoreThanOneCount += 1
		}
	}

	nonOverlappingClaim := Claim{}
	for _, claim := range claims {
		pointsWhichAreOccupiedOnlyByThisClaim := 0

		for _, claimsOccupying := range fabric {
			if len(claimsOccupying) == 1 {
				if claimsOccupying[0].Id == claim.Id {
					pointsWhichAreOccupiedOnlyByThisClaim += 1
				}
			}
		}

		if pointsWhichAreOccupiedOnlyByThisClaim == claim.Height*claim.Width {
			nonOverlappingClaim = claim
			break
		}
	}

	fmt.Println("Fabric points within more than one claim:", occupiedByMoreThanOneCount)
	fmt.Println("Non overlaping claim ID:", nonOverlappingClaim.Id)
}
