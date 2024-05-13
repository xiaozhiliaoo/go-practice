package unit_test

import (
	"fmt"
	sm "github.com/cch123/supermonkey"
	"testing"
)

func TestPatchByFullSymbolName(t *testing.T) {
	fmt.Println("original function output:")
	heyHeyHey()

	patchGuard := sm.PatchByFullSymbolName("unit_test.heyHeyHey", func() {
		fmt.Println("please be polite")
	})
	fmt.Println("after patch, function output:")
	heyHeyHey()

	patchGuard.Unpatch()
	fmt.Println("unpatch, then output:")
	heyHeyHey()
}

//go:noinline
func heyHeyHey() {
	fmt.Println("fake")
}
