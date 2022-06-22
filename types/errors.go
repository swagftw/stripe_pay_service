package types

import "errors"

type CopyError error

var ErrCopyingData CopyError = errors.New("error copying data")
