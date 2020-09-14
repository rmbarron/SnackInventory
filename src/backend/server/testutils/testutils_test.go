/*
Copyright 2020 Robert Barron

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package testutils

import (
	"testing"

	sipb "github.com/rmbarron/SnackInventory/src/proto/snackinventory"
)

func TestSnackKeyMapComparerFunc(t *testing.T) {
	for _, tc := range []struct {
		name string
		x    map[*sipb.Snack]int
		y    map[*sipb.Snack]int
		want bool
	}{
		{
			name: "MismatchedLength",
			x: map[*sipb.Snack]int{
				&sipb.Snack{Barcode: "1337"}: 1,
			},
			y: map[*sipb.Snack]int{
				&sipb.Snack{Barcode: "1337"}:  1,
				&sipb.Snack{Barcode: "31337"}: 1,
			},
			want: false,
		},
		{
			name: "SameLengthMismatchedKeys",
			x: map[*sipb.Snack]int{
				&sipb.Snack{Barcode: "1337"}:  1,
				&sipb.Snack{Barcode: "31337"}: 1,
			},
			y: map[*sipb.Snack]int{
				&sipb.Snack{Barcode: "1337", Name: "leetTreat"}:   1,
				&sipb.Snack{Barcode: "31337", Name: "eleetTreat"}: 1,
			},
			want: false,
		},
		{
			name: "MismatchedValues",
			x: map[*sipb.Snack]int{
				&sipb.Snack{Barcode: "1337"}: 1,
			},
			y: map[*sipb.Snack]int{
				&sipb.Snack{Barcode: "1337"}: 2,
			},
			want: false,
		},
		{
			name: "Match",
			x: map[*sipb.Snack]int{
				&sipb.Snack{Barcode: "1337"}: 1,
			},
			y: map[*sipb.Snack]int{
				&sipb.Snack{Barcode: "1337"}: 1,
			},
			want: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := snackKeyMapComparerFunc(tc.x, tc.y)
			if got != tc.want {
				t.Errorf("snackKeyMapComparerFunc(%v, %v) = got %v, want %v", tc.x, tc.y, got, tc.want)
			}
		})
	}
}
