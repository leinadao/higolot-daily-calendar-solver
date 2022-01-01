package main

import (
	"errors"
	"fmt"
)

type piece struct {
	shape [][]bool
}

var pieces []piece = []piece{
	{shape: [][]bool{
		{true, true, true},
		{true, false, false},
		{true, false, false},
	}},
	{shape: [][]bool{
		{true, false},
		{true, false},
		{true, false},
		{true, true},
	}},
	{shape: [][]bool{
		{true, true, false},
		{false, true, true},
	}},
	{shape: [][]bool{
		{true, false, false},
		{true, true, true},
		{true, false, false},
	}},
	{shape: [][]bool{
		{false, true, true, true},
		{true, true, false, false},
	}},
	{shape: [][]bool{
		{true, true, true},
		{true, true, false},
	}},
	{shape: [][]bool{
		{true, true, true},
		{true, false, true},
	}},
	{shape: [][]bool{
		{false, true},
		{false, true},
		{true, true},
	}},
	{shape: [][]bool{
		{true, false, false},
		{true, true, true},
		{false, false, true},
	}},
	{shape: [][]bool{
		{true},
		{true},
		{true},
		{true},
	}},
}

func (p *piece) width() int {
	return len(p.shape[0])
}

func (p *piece) height() int {
	return len(p.shape)
}

func (p *piece) getEmptyRotatedArray() *[][]bool {
	newArr := make([][]bool, p.width())
	for i := range newArr {
		newArr[i] = make([]bool, p.height())
	}
	return &newArr
}

func (p *piece) transpose() {
	newArr := *p.getEmptyRotatedArray()
	for iy := 0; iy < p.height(); iy++ {
		for ix := 0; ix < p.width(); ix++ {
			newArr[ix][iy] = p.shape[iy][ix]
		}
	}
	p.shape = newArr
}

var grid [][]string = [][]string{
	{"Jan", "Feb", "Mar", "Apr", "May", "Jun", ""},
	{"Jul", "Aug", "Sep", "Oct", "Nov", "Dec", ""},
	{"1", "2", "3", "4", "5", "6", "7"},
	{"8", "9", "10", "11", "12", "13", "14"},
	{"15", "16", "17", "18", "19", "20", "21"},
	{"22", "23", "24", "25", "26", "27", "28"},
	{"29", "30", "31", "Sun", "Mon", "Tues", "Wed"},
	{"", "", "", "", "Thur", "Fri", "Sat"},
}

func setTarget(month, dayNum, day string) error {
	setNum := 0
	for iy, r := range grid {
		for ix, x := range r {
			if x == month || x == dayNum || x == day {
				grid[iy][ix] = ""
				setNum++
			}
		}
	}
	if setNum != 3 {
		return errors.New("failed to set target")
	}
	return nil
}

func solve(month, dayNum, day string) {
	setTarget(month, dayNum, day)
	// TODO..
}

func main() {
	var month, dayNum, day string
	fmt.Println("Enter the month to solve for e.g. 'Jan' or 'Feb': ")
	fmt.Scanln(&month)
	fmt.Println("Enter the day number solve for e.g. '1' or '10': ")
	fmt.Scanln(&dayNum)
	fmt.Println("Enter the day to solve for e.g. 'Mon' or 'Tues': ")
	fmt.Scanln(&day)
	solve(month, dayNum, day)
}
