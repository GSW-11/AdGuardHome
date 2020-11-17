package aghio

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimitedReadCloser_Read(t *testing.T) {
	testCases := []struct {
		name  string
		limit int64
		rStr  string
		want  int
		err   error
	}{{
		name:  "perfectly_match",
		limit: 3,
		rStr:  "abc",
		want:  3,
		err:   nil,
	}, {
		name:  "eof",
		limit: 3,
		rStr:  "",
		want:  0,
		err:   io.EOF,
	}, {
		name:  "limit_reached",
		limit: 0,
		rStr:  "abc",
		want:  0,
		err:   ErrLimitReached,
	}, {
		name:  "truncated",
		limit: 2,
		rStr:  "abc",
		want:  2,
		err:   nil,
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			readCloser := ioutil.NopCloser(strings.NewReader(tc.rStr))
			buf := make([]byte, tc.limit+1)

			lreader := LimitReadCloser(readCloser, tc.limit)
			n, err := lreader.Read(buf)

			assert.Equal(t, n, tc.want)
			assert.True(t, errors.Is(err, tc.err), buf)
		})
	}
}
