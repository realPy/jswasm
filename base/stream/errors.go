package stream

import "errors"

var (
	ErrNotImplementedReadableStream         = errors.New("Browser not implemented ReadableStream")
	ErrNotImplementedWritableStream         = errors.New("Browser not implemented WritableStream")
	ErrNotImplementedTransformStream        = errors.New("Browser not implemented TransformStream")
	ErrNotAReadableStream                   = errors.New("Object is not a ReadableStream")
	ErrNotAWritableStream                   = errors.New("Object is not a WritableStream")
	ErrNotATransformStream                  = errors.New("Object is not a TransformStream")
	ErrNotAReadableStreamDefaultReader      = errors.New("Object is not a ReadableStreamDefaultReader")
	ErrNotATransformStreamDefaultController = errors.New("Object is not a TransformStreamDefaultController")
)
