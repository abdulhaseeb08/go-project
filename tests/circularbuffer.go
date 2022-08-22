package test

import "fmt"

// Define the Buffer type here.

type Buffer struct {
	readHead, writeHead, curentItems int
	buffer                           []byte
}

func NewBuffer(size int) *Buffer {
	return &Buffer{buffer: make([]byte, size)}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.curentItems == 0 {
		return 0, fmt.Errorf("empty buffer")
	}
	x := b.buffer[b.readHead]
	b.readHead = (b.readHead + 1) % len(b.buffer)
	b.curentItems--
	return x, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.curentItems == len(b.buffer) {
		return fmt.Errorf("Buffer full")
	}
	b.buffer[b.writeHead] = c
	b.writeHead = (b.writeHead + 1) % len(b.buffer)
	b.curentItems++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.curentItems < len(b.buffer) {
		b.WriteByte(c)
	} else {
		b.buffer[b.readHead] = c
		b.readHead = (b.readHead + 1) % len(b.buffer)
	}
}

func (b *Buffer) Reset() {
	b.readHead = 0
	b.writeHead = 0
	b.curentItems = 0
}
