package gostarter

import _ "embed"

// Go allows this because templates/ is a subdirectory of the root
//go:embed templates/go/basic.yaml
var DefaultGoBasicTemplate []byte