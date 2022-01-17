package merkle_test

import (
	"fmt"
	"reflect"
	"simcap/srtest/merkle"
	"testing"
)

func TestLevel(t *testing.T) {
	assert := func(actual []string, exp ...string) {
		if got, want := actual, exp; !reflect.DeepEqual(got, want) {
			t.Fatalf("\ngot\n%v\nwant\n%v\n", got, want)
		}
	}

	tree := merkle.BuildTree("data_1", "data_2", "data_3", "data_4")
	assert(tree.Level(0), "bfdc3a56f1351828c94fe5bcf68a091abe468f6edf85bf3edbc0d083fbc1f037")
	assert(tree.Level(1), "88c0d7257704a0e77a3725c3161eac2653fc5ee10fca9f63c9d0bd62c31ce410", "117859a7e70956f9b77a6511a6170155efb17cf630d4bcee82c65896bd6edfa3")
	assert(tree.Level(2), "28cebd73182082c77a65a23f4b30fdf81bc7f6af1a9d504ce7c8d1e22f847de9", "7b1dda0e7b4524eae7432b56994fff094891c48d660a1f314b7ec14f611acaf4", "77805cd60afed8a6f4b5f6e404ecdade9395d0e5e16ef815e3e9584329edb422", "faba64f502ac5bd7644b8a2601cc7d2d4078cc6acd7035b25a66066787de3373")
}

func TestBuildTree(t *testing.T) {

	t.Run("with valid root hash and height", func(t *testing.T) {
		cases := []struct {
			input  []string
			root   string
			height int
		}{
			{
				input:  []string{"data_1"},
				height: 1, root: "28cebd73182082c77a65a23f4b30fdf81bc7f6af1a9d504ce7c8d1e22f847de9",
			},
			{
				input:  []string{"data_1", "data_2"},
				height: 2, root: "88c0d7257704a0e77a3725c3161eac2653fc5ee10fca9f63c9d0bd62c31ce410",
			},
			{
				input:  []string{"data_1", "data_2", "data_3"},
				height: 3, root: "bfbfbd4fed7063dac7e6e8e1bf4250794d0a1bdaae58a9c56c19fc20cf889c7a",
			},
			{
				input:  []string{"data_1", "data_2", "data_3", "data_4"},
				height: 3, root: "bfdc3a56f1351828c94fe5bcf68a091abe468f6edf85bf3edbc0d083fbc1f037",
			},
		}

		for _, c := range cases {
			t.Run(fmt.Sprintf("data_length_of_%d", len(c.input)), func(t *testing.T) {
				tree := merkle.BuildTree(c.input...)
				if got, want := tree.Height(), c.height; got != want {
					t.Fatalf("got %v, want %v", got, want)
				}
				if got, want := tree.Root(), c.root; got != want {
					t.Fatalf("got %v, want %v", got, want)
				}
			})
		}
	})
}
