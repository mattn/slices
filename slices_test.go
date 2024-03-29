// Copyright © The Slices Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slices_test

import (
	"github.com/vorduin/slices"
	"testing"
)

func TestWithLen(t *testing.T) {
	s := slices.WithLen[int](10)

	if len(s) != 10 {
		t.Failed()
	}
}

func TestWithCap(t *testing.T) {
	s := slices.WithCap[int](10)

	if len(s) == 10 {
		t.Fail()
	}

	if cap(s) != 10 {
		t.Failed()
	}
}

func TestCloneInt(t *testing.T) {
	orig := []int{1, 2, 3}
	cp := slices.Clone(orig)
	orig[0] = 0

	if cp[0] == 0 {
		t.Failed()
	}
}

func TestCloneString(t *testing.T) {
	orig := []string{"the", "slices", "author"}
	cp := slices.Clone(orig)
	orig[0] = "a"

	if cp[0] == "a" {
		t.Failed()
	}
}

func TestClone2DSlice(t *testing.T) {
	orig := [][]int{{1}, {2}, {3}}
	cp := slices.Clone(orig)
	orig[0] = []int{0}

	if cp[0][0] == 0 {
		t.Failed()
	}
}

func TestEqualDiffLen(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3, 4, 5}

	if slices.Equal(s1, s2) {
		t.Failed()
	}
}

func TestEqualInt(t *testing.T) {
	s1 := []int{1, 2, 3}
	if !slices.Equal(s1, s1) {
		t.Fail()
	}

	s2 := []int{0, 2, 3}
	if slices.Equal(s1, s2) {
		t.Fail()
	}
}

func TestEqualString(t *testing.T) {
	s1 := []string{"the", "slices", "author"}
	if !slices.Equal(s1, s1) {
		t.Fail()
	}

	s2 := []string{"a", "slices", "author"}
	if slices.Equal(s1, s2) {
		t.Fail()
	}
}
