// Package aghio contains extensions for io package's types and methods
package aghio

import (
	"fmt"
	"io"

	"github.com/AdguardTeam/AdGuardHome/internal/agherr"
)

// ErrLimitReached is returned if the limit of LimitedReader is reached.
const ErrLimitReached agherr.Error = "limit reached"

// limitedReadCloser is a wrapper for io.ReadCloser with limited reader and
// dealing with agherr package.
type limitedReadCloser struct {
	limit int64
	N     int64
	io.ReadCloser
}

// Read implements Reader interface.
func (lrc *limitedReadCloser) Read(p []byte) (n int, err error) {
	if lrc.N <= 0 {
		return 0, fmt.Errorf("read %d bytes: %w", lrc.limit, ErrLimitReached)
	}
	if int64(len(p)) > lrc.N {
		p = p[0:lrc.N]
	}
	n, err = lrc.ReadCloser.Read(p)
	lrc.N -= int64(n)
	return n, err
}

// Close implements Closer interface.
func (lrc *limitedReadCloser) Close() error {
	return lrc.ReadCloser.Close()
}

// LimitReadCloser returns a ReadCloser with original Closer and Reader that
// stops with ErrLimitReached after n bytes read.
func LimitReadCloser(rc io.ReadCloser, n int64) io.ReadCloser {
	return &limitedReadCloser{
		limit:      n,
		N:          n,
		ReadCloser: rc,
	}
}
