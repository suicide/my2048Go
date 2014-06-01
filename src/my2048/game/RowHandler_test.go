package game

import ("testing"

		"reflect"
)

func make0Row() []Cell {

	row := make([]Cell, 0, BoardSize)

	for i := 0; i < BoardSize; i++ {
		row = append(row, GetCell(0))
	}

	return row

}

func makeRow(values []int) []Cell {
	row := make([]Cell, 0, BoardSize)

	for i := 0; i < BoardSize; i++ {
		row = append(row, GetCell(values[i]))
	}

	return row
}

func TestFoldRowEmpty(t *testing.T) {

	row := make0Row()

	newRow, score, changed := foldRow(row)

	if !reflect.DeepEqual(newRow, row) {
		t.Fail()
	}
	if score != 0 {
		t.Fail()
	}
	if changed {
		t.Fail()
	}
}

func TestFoldRowSingle2(t *testing.T) {

	row := makeRow([]int{2,0,0,0})

	newRow, score, changed := foldRow(row)

	if !reflect.DeepEqual(newRow, makeRow([]int{2,0,0,0})) {
		t.Fail()
	}
	if score != 0 {
		t.Fail()
	}
	if changed {
		t.Fail()
	}
}

func TestFoldRowSingle2Middle(t *testing.T) {

	row := makeRow([]int{0,2,0,0})

	newRow, score, changed := foldRow(row)

	if !reflect.DeepEqual(newRow, makeRow([]int{2,0,0,0})) {
		t.Fail()
	}
	if score != 0 {
		t.Fail()
	}
	if !changed {
		t.Fail()
	}
}

func TestFoldRow0204(t *testing.T) {

	row := makeRow([]int{0,2,0,4})

	newRow, score, changed := foldRow(row)

	if !reflect.DeepEqual(newRow, makeRow([]int{2,4,0,0})) {
		t.Fail()
	}
	if score != 0 {
		t.Fail()
	}
	if !changed {
		t.Fail()
	}
}

func TestFoldRow2200(t *testing.T) {

	row := makeRow([]int{2,2,0,0})

	newRow, score, changed := foldRow(row)

	if !reflect.DeepEqual(newRow, makeRow([]int{4,0,0,0})) {
		t.Fail()
	}
	if score != 4 {
		t.Fail()
	}
	if !changed {
		t.Fail()
	}
}

func TestFoldRow0022(t *testing.T) {

	row := makeRow([]int{0,0,2,2})

	newRow, score, changed := foldRow(row)

	if !reflect.DeepEqual(newRow, makeRow([]int{4,0,0,0})) {
		t.Fail()
	}
	if score != 4 {
		t.Fail()
	}
	if !changed {
		t.Fail()
	}
}

func TestFoldRow0202(t *testing.T) {

	row := makeRow([]int{0,2,0,2})

	newRow, score, changed := foldRow(row)

	if !reflect.DeepEqual(newRow, makeRow([]int{4,0,0,0})) {
		t.Fail()
	}
	if score != 4 {
		t.Fail()
	}
	if !changed {
		t.Fail()
	}
}

func TestFoldRow0222(t *testing.T) {

	row := makeRow([]int{0,2,2,2})

	newRow, score, changed := foldRow(row)

	if !reflect.DeepEqual(newRow, makeRow([]int{4,2,0,0})) {
		t.Fail()
	}
	if score != 4 {
		t.Fail()
	}
	if !changed {
		t.Fail()
	}
}

func TestFoldRow2222(t *testing.T) {

	row := makeRow([]int{2,2,2,2})

	newRow, score, changed := foldRow(row)

	if !reflect.DeepEqual(newRow, makeRow([]int{4,4,0,0})) {
		t.Fail()
	}
	if score != 8 {
		t.Fail()
	}
	if !changed {
		t.Fail()
	}
}
