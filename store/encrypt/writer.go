package encrypt

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/disorganizer/brig/util"
	"github.com/glycerine/rbuf"
)

// Writer encrypts the data stream before writing to Writer.
type Writer struct {
	// Common fields with Reader
	aeadCommon

	// Internal Writer we would write to.
	Writer io.Writer

	// A buffer that is MaxBlockSize big.
	// Used for caching blocks
	rbuf *rbuf.FixedSizeRingBuf

	// True after the first write.
	headerWritten bool
}

func (w *Writer) Write(p []byte) (int, error) {
	if !w.headerWritten {
		w.headerWritten = true

		_, err := w.Writer.Write(GenerateHeader())
		if err != nil {
			return 0, err
		}
	}

	for w.rbuf.Readable >= MaxBlockSize {
		_, err := w.flushPack(MaxBlockSize)
		if err != nil {
			return 0, err
		}
	}

	// Remember left-overs for next write:
	_, err := w.rbuf.Write(p)
	if err != nil {
		return 0, nil
	}

	// Fake the amount of data we've written:
	return len(p), nil
}

func (w *Writer) flushPack(chunkSize int) (int, error) {
	n, err := w.rbuf.Read(w.decBuf[:chunkSize])
	if err != nil {
		return 0, err
	}

	// Create a new Nonce for this block:
	blockNum := binary.BigEndian.Uint64(w.nonce)
	binary.BigEndian.PutUint64(w.nonce, blockNum+1)

	// Encrypt the text:
	w.encBuf = w.aead.Seal(w.encBuf[:0], w.nonce, w.decBuf[:n], nil)

	// Pass it to the underlying writer:
	written := 0
	if n, err := w.Writer.Write(w.nonce); err == nil {
		written += n
	}

	if n, err := w.Writer.Write(w.encBuf); err == nil {
		written += n
	}

	return written, nil
}

// Seek the write stream. This maps to a seek in the underlying datastream.
func (w *Writer) Seek(offset int64, whence int) (int64, error) {
	if seeker, ok := w.Writer.(io.Seeker); ok {
		return seeker.Seek(offset, whence)
	}

	return 0, fmt.Errorf("write: Seek is not supported by underlying datastream")
}

// Close the Writer and write any left-over blocks
// This does not close the underlying data stream.
func (w *Writer) Close() error {
	for w.rbuf.Readable > 0 {
		n := util.Min(MaxBlockSize, w.rbuf.Readable)
		_, err := w.flushPack(n)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewWriter returns a new Writer which encrypts data with a
// certain key.
func NewWriter(w io.Writer, key []byte) (*Writer, error) {
	writer := &Writer{
		Writer: w,
		rbuf:   rbuf.NewFixedSizeRingBuf(MaxBlockSize * 2),
	}

	if err := writer.initAeadCommon(key, defaultCipherType); err != nil {
		return nil, err
	}

	return writer, nil
}