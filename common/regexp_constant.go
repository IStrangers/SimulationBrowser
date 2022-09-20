package common

import "regexp"

var REGEXP_CONSTANT_SPACES = regexp.MustCompile(`/^[ \n\r\t\f]+/`)
