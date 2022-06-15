//go:build !windows

package assets

import _ "embed"

//go:embed logo.png
var logo []byte
