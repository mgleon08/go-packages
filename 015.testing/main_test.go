package main

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("失敗")
	} else {
		t.Log("成功")
	}
}

func Test_Division_2(t *testing.T) {
	t.Error("失敗")
}

func Test_Division_table(t *testing.T) {
	tables := []struct {
		x float64
		y float64
	}{
		{3, 1},
		{6, 2},
		{9, 3},
		{8, 2},
	}

	for _, table := range tables {
		if i, e := Division(table.x, table.y); i != 3 || e != nil {
			t.Error("失敗")
		} else {
			t.Log("成功")
		}
	}
}
