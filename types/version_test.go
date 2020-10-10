// Copyright 2013-2016 lessgo Author, All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"testing"
)

type valid_entry struct {
	v  string
	ok bool
}

func TestVersion(t *testing.T) {

	v1 := Version("2.0.12.1.aa")
	v2 := Version("2.0.20.2,bb")
	v3 := Version("2.0.20.2,bb")

	if v1.Compare(v2) != -1 || v2.Compare(v3) != 0 || v2.Compare(v1) != 1 {
		t.Fatal("Failed on Compare")
	}

	if v1.String() != "2.0.12.1.aa" {
		t.Fatal("Failed on String")
	}

	vs := []valid_entry{
		{"1.0.0", true},
		{"2.0.0.aa", true},
		{"1", true},
		{"100", true},
		{"1.dev-beta", true},
		{"1.dev_beta", true},
		{"v1", true},
		{"", false},
		{"-", false},
		{".", false},
		{" ", false},
		{"!", false},
		{".0", false},
		{".0.", false},
	}
	for _, v := range vs {
		vv := Version(v.v)
		if vv.Valid() != v.ok {
			t.Fatal("Failed on Valid " + v.v)
		}
	}
}

func Benchmark_Version_Valid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1 := Version("10.10.10.dev")
		v1.Valid()
	}
}

func Benchmark_VersionCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VersionCompare("10.10.10", "10.10.10")
	}
}
