package day3

import (
	"fmt"
	"strconv"
	"strings"
)

// Overlaps finds the claims that overlap
func Overlaps(claims []string) int {
	fabric := CreateMap(1000, 1000)
	for _, claim := range claims {
		fabric = AddToMap(fabric, ParseClaim(claim))
	}
	return CountThe(fabric, "X")
}

// CountThe number of squares with a particular claim
func CountThe(fabric [][]string, ID string) int {
	overlaps := 0
	for i := 0; i < len(fabric); i++ {
		for j := 0; j < len(fabric[0]); j++ {
			if fabric[i][j] == ID {
				overlaps++
			}
		}
	}
	return overlaps
}

// Claim that elves make
type Claim struct {
	ID     int
	X      int
	Y      int
	Width  int
	Height int
}

// ParseClaim finds out the specifics about the elves claim
func ParseClaim(claim string) Claim {
	parsed := strings.Split(claim, " ")
	ID, err := strconv.Atoi(strings.Replace(parsed[0], "#", "", 1))
	if err != nil {
		panic(err)
	}
	X, err := strconv.Atoi(strings.Split(parsed[2], ",")[0])
	if err != nil {
		panic(err)
	}
	Y, err := strconv.Atoi(strings.Replace(strings.Split(parsed[2], ",")[1], ":", "", 1))
	if err != nil {
		panic(err)
	}
	Width, err := strconv.Atoi(strings.Split(parsed[3], "x")[0])
	if err != nil {
		panic(err)
	}
	Height, err := strconv.Atoi(strings.Split(parsed[3], "x")[1])
	if err != nil {
		panic(err)
	}
	return Claim{ID, X, Y, Width, Height}
}

// AddToMap updates the map with a claim
func AddToMap(ogmap [][]string, claim Claim) [][]string {
	for i := claim.Y; i < claim.Y+claim.Height; i++ {
		for j := claim.X; j < claim.X+claim.Width; j++ {
			if ogmap[i][j] == "." {
				ogmap[i][j] = strconv.Itoa(claim.ID)
			} else {
				ogmap[i][j] = "X"
			}
		}
	}
	return ogmap
}

// CreateMap creates the map
func CreateMap(x int, y int) [][]string {
	themap := [][]string{}
	for i := 0; i < y; i++ {
		themap = append(themap, []string{})
		for j := 0; j < y; j++ {
			themap[i] = append(themap[i], ".")
		}
	}
	return themap
}

// FindPerfectClaim finds the claim that santa needs
func FindPerfectClaim(claims []string) int {
	fabric := CreateMap(1000, 1000)
	for _, claim := range claims {
		fabric = AddToMap(fabric, ParseClaim(claim))
	}
	for _, claim := range claims {
		c := ParseClaim(claim)
		area := c.Width * c.Height
		count := CountThe(fabric, strconv.Itoa(c.ID))
		fmt.Printf("Testing Claim %d: area=%d and count==%d \n", c.ID, area, count)
		if area == count {
			return c.ID
		}
	}
	return 0
}
