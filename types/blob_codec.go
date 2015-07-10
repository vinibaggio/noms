package types

import (
	"bytes"
	"io"

	"github.com/attic-labs/noms/chunks"
	. "github.com/attic-labs/noms/dbg"
	"github.com/attic-labs/noms/ref"
)

var (
	blobTag = []byte("b ")
)

func blobEncode(b Blob, s chunks.ChunkSink) (r ref.Ref, err error) {
	w := s.Put()
	if _, err = w.Write(blobTag); err != nil {
		return
	}
	if _, err = io.Copy(w, b.Reader()); err != nil {
		return
	}
	return w.Ref()
}

func blobDecode(r io.Reader, s chunks.ChunkSource) (Value, error) {
	buf := &bytes.Buffer{}
	_, err := io.CopyN(buf, r, int64(len(blobTag)))
	if err != nil {
		return nil, err
	}
	Chk.True(bytes.Equal(buf.Bytes(), blobTag))

	buf.Truncate(0)
	_, err = io.Copy(buf, r)
	if err != nil {
		return nil, err
	}
	return NewBlob(buf.Bytes()), nil
}