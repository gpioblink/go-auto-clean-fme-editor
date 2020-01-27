package fme

type LyricBlock struct {
	LyricHeader
	LyricBody
}

func (lb LyricBlock) CalcByteSize() int {
	return LyricHeaderSize + lb.LyricBody.CalcByteSize()
}

func NewLyricBlock(header LyricHeader, body LyricBody) (*LyricBlock, error) {
	return &LyricBlock{header, body}, nil
}
