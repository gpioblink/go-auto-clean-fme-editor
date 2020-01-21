package fme

type Fme struct {
	HeaderDataPart
	InformationDataPart
	LyricDataPart
	TimingDataPart
}

func NewFmeFromBinary(fme []byte) (*Fme, error) {
	headerData, err := NewHeaderDataPartFromBinary(fme)
	if err != nil {
		return nil, err
	}

	infoOffset, lyricOffset, timingOffset := headerData.GetOffsets()

	infoData, err := NewInformationDataPartFromBinary(fme[infoOffset:lyricOffset])
	if err != nil {
		return nil, err
	}

	lyricData, err := NewLyricDataPartFromBinary(fme[lyricOffset:timingOffset])
	if err != nil {
		return nil, err
	}

	timingData, err := NewTimingDataPartFromBinary(fme[timingOffset:])
	if err != nil {
		return nil, err
	}

	return &Fme{*headerData, *infoData, *lyricData, *timingData}, nil
}

func (f Fme) ExportBinary() ([]byte, error) {

	info, err := f.InformationDataPart.ExportBinary()
	if err != nil {
		return nil, err
	}

	lyric, err := f.LyricDataPart.ExportBinary()
	if err != nil {
		return nil, err
	}

	timing, err := f.TimingDataPart.ExportBinary()
	if err != nil {
		return nil, err
	}

	// recalculate offsets for header
	infoOffset := uint32(0x12)
	lyricOffset := infoOffset + uint32(len(info))
	timingOffset := lyricOffset + uint32(len(lyric))
	err = f.HeaderDataPart.UpdateOffsets(infoOffset, lyricOffset, timingOffset)
	if err != nil {
		return nil, err
	}

	header, err := f.HeaderDataPart.ExportBinary()
	if err != nil {
		return nil, err
	}

	var fmeBinary []byte
	fmeBinary = append(fmeBinary, header...)
	fmeBinary = append(fmeBinary, info...)
	fmeBinary = append(fmeBinary, lyric...)
	fmeBinary = append(fmeBinary, timing...)

	return fmeBinary, nil
}
