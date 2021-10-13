package gcl

type Buffer []byte

func (b *Buffer) Reset() {
	*b = Buffer([]byte(*b)[:0])
}

func (b *Buffer) Append(data []byte) {
	*b = append(*b, data...)
}

func (b *Buffer) AppendByte(data byte) {
	*b = append(*b, data)
}
