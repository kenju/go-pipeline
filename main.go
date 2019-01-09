/*
 Concurrent stream-processing library in Go.

 Repository: https://github.com/kenju/go-pipeline
 */
package pipeline

import (
	"github.com/cheekybits/genny/generic"
)

// GenType is used for `go generate`, which is offered by github.com/cheekybits/genny.
//
// Add specific comments (which starts with "//go:generate ..." into the source code
// and run `go generate`.
//
// Read https://github.com/cheekybits/genny for more details.
type GenType generic.Type
