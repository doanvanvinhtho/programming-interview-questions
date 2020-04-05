package util

import (
	"errors"
	"testing"
)

func TestArrayIntEqual(t *testing.T) {
	cases := []struct {
		inA []int
		inB []int
		out bool
	}{
		{
			inA: []int{},
			inB: []int{},
			out: true,
		},
		{
			inA: []int{1},
			inB: []int{},
			out: false,
		},
		{
			inA: []int{},
			inB: []int{1},
			out: false,
		},
		{
			inA: []int{2, 1},
			inB: []int{2, 1},
			out: true,
		},
		{
			inA: []int{2, 1},
			inB: []int{2, 2},
			out: false,
		},
	}
	for i, c := range cases {
		r := ArrayIntEqual(c.inA, c.inB)
		if c.out != r {
			t.Errorf("Wrong %v got %v", i, r)
		}
	}
}

func TestArrayStringEqual(t *testing.T) {
	cases := []struct {
		inA []string
		inB []string
		out bool
	}{
		{
			inA: []string{},
			inB: []string{},
			out: true,
		},
		{
			inA: []string{"1"},
			inB: []string{},
			out: false,
		},
		{
			inA: []string{},
			inB: []string{"1"},
			out: false,
		},
		{
			inA: []string{"2", "1"},
			inB: []string{"2", "1"},
			out: true,
		},
		{
			inA: []string{"2", "1"},
			inB: []string{"2", "2"},
			out: false,
		},
	}
	for i, c := range cases {
		r := ArrayStringEqual(c.inA, c.inB)
		if c.out != r {
			t.Errorf("Wrong %v got %v", i, r)
		}
	}
}

func TestArrayIntMax(t *testing.T) {
	cases := []struct {
		in     []int
		outMax int
		outErr error
	}{
		{
			in:     []int{},
			outMax: -1,
			outErr: errors.New("The array is empty"),
		},
		{
			in:     []int{1},
			outMax: 1,
			outErr: nil,
		},
		{
			in:     []int{2, 1},
			outMax: 2,
			outErr: nil,
		},
		{
			in:     []int{1, 2},
			outMax: 2,
			outErr: nil,
		},
		{
			in:     []int{1, 1},
			outMax: 1,
			outErr: nil,
		},
	}
	for i, c := range cases {
		max, e := ArrayIntMax(c.in)
		if (c.outErr != nil && e == nil) ||
			(c.outErr == nil && e != nil) {
			t.Errorf("Wrong %v got %v", i, e)
		}
		if c.outMax != max {
			t.Errorf("Wrong %v got %v", i, max)
		}
	}
}
