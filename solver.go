package main

import (
	"errors"
	"fmt"
)

// piece is a higolot daily calendar puzzle piece.
type piece [][]bool

// width returns the number of grid columns required to accommodate the piece.
func (p *piece) width() int {
	return len((*p)[0])
}

// height returns the number of grid rows required to accommodate the piece.
func (p *piece) height() int {
	return len(*p)
}

// getEmptyRotatedArray returns an empty rotated copy of the piece's representation.
func (p *piece) getEmptyRotatedArray() piece {
	newArr := make([][]bool, p.width())
	for i := range newArr {
		newArr[i] = make([]bool, p.height())
	}
	return newArr
}

// transpose transposes the piece.
func (p *piece) transpose() {
	newArr := p.getEmptyRotatedArray()
	for iy := 0; iy < p.height(); iy++ {
		for ix := 0; ix < p.width(); ix++ {
			newArr[ix][iy] = (*p)[iy][ix]
		}
	}
	*p = newArr
}

// grid is a higolot daily calendar layout.
type grid [][]string

// width returns the grid width in piece spaces.
func (g *grid) width() int {
	return len((*g)[0])
}

// height returns the grid height in piece spaces.
func (g *grid) height() int {
	return len(*g)
}

// setTarget marks the grid with the given target month, day number and day name.
// i.e. the target spots of the grid are blanked out.
// An error is returned if it's unable to set all three targets.
func (g *grid) setTarget(month, dayNum, dayName string) error {
	setNum := 0
	for iy, r := range *g {
		for ix, x := range r {
			if x == month || x == dayNum || x == dayName {
				(*g)[iy][ix] = ""
				setNum++
			}
		}
	}
	if setNum != 3 {
		return errors.New("failed to set target")
	}
	return nil
}

// puzzle is a higolot daily calendar.
type puzzle struct {
	grid   grid
	pieces []piece
}

// solve tries to solve the puzzle for the given month and day number and name.
func (p *puzzle) solve(month, dayNum, dayName string) {
	p.grid.setTarget(month, dayNum, dayName)
	// TODO..
}

var p puzzle = puzzle{
	grid: grid{
		{"Jan", "Feb", "Mar", "Apr", "May", "Jun", ""},
		{"Jul", "Aug", "Sep", "Oct", "Nov", "Dec", ""},
		{"1", "2", "3", "4", "5", "6", "7"},
		{"8", "9", "10", "11", "12", "13", "14"},
		{"15", "16", "17", "18", "19", "20", "21"},
		{"22", "23", "24", "25", "26", "27", "28"},
		{"29", "30", "31", "Sun", "Mon", "Tues", "Wed"},
		{"", "", "", "", "Thur", "Fri", "Sat"},
	}, pieces: []piece{
		{
			{true, true, true},
			{true, false, false},
			{true, false, false},
		},
		{
			{true, false},
			{true, false},
			{true, false},
			{true, true},
		},
		{
			{true, true, false},
			{false, true, true},
		},
		{
			{true, false, false},
			{true, true, true},
			{true, false, false},
		},
		{
			{false, true, true, true},
			{true, true, false, false},
		},
		{
			{true, true, true},
			{true, true, false},
		},
		{
			{true, true, true},
			{true, false, true},
		},
		{
			{false, true},
			{false, true},
			{true, true},
		},
		{
			{true, false, false},
			{true, true, true},
			{false, false, true},
		},
		{
			{true},
			{true},
			{true},
			{true},
		},
	},
}

func main() {
	var month, dayNum, day string
	fmt.Println("Enter the month to solve for e.g. 'Jan' or 'Feb': ")
	fmt.Scanln(&month)
	fmt.Println("Enter the day number solve for e.g. '1' or '10': ")
	fmt.Scanln(&dayNum)
	fmt.Println("Enter the day to solve for e.g. 'Mon' or 'Tues': ")
	fmt.Scanln(&day)
	// TODO: Validation (just handle the error from setting target?)?
	p.solve(month, dayNum, day)
}
