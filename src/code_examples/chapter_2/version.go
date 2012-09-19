// list runtime.Version()
package main

import (
	"fmt"
  "runtime"
)

func main() {
	fmt.Printf("%s\n", runtime.Version())
}
// Output:
// 7053 release.2011-01-06 release
// release.r57.1 8330
