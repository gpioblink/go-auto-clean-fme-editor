package lyric

import (
	intraproces "github.com/gpioblink/go-auto-clean-fme-editor/editor/interfaces/private/intraprocess"
	"github.com/gpioblink/go-auto-clean-fme-editor/fme/converterDomain/fme"
)

type IntraprocessService struct {
	intraprocessInterface intraproces.LyricInterface
}

func NewIntraprocessService(intraprocessInterface intraproces.LyricInterface) IntraprocessService {
	return IntraprocessService{intraprocessInterface}
}

func (i IntraprocessService) AddLyric(block fme.LyricBlock, colorPicker fme.LyricColorPicker) error {

}

func (i IntraprocessService) ListLyrics() (blocks []fme.LyricBlock, colorPicker fme.LyricColorPicker, err error) {
	gotLyric, err := i.intraprocessInterface.ListLyrics()
	if err != nil {
		return nil, fme.LyricColorPicker{}, err
	}

	blocks, colorPicker, err = fmeBlockAndColorPickerFromIntraprocess(gotLyric)
	return blocks, colorPicker, err
}

func fmeBlockAndColorPickerFromIntraprocess(intraLyric []intraproces.LyricView) (blocks []fme.LyricBlock, colorPicker fme.LyricColorPicker, err error) {
	// TODO: 標準色に含まれない色があった場合の処理

	blocks = []fme.LyricBlock{}
	for _, il := range intraLyric {

		// lyric
		var lyricString []fme.LyricChar
		for _, l := range il.Lyric {
			lyricChar, err := fme.NewLyricChar(l.LyricChar, l.Length)
			if err != nil {
				return nil, fme.LyricColorPicker{}, err
			}
			lyricString = append(lyricString, *lyricChar)
		}

		// ruby
		var rubyString []fme.LyricRuby
		for _, r := range il.Ruby {
			ruby, err := fme.NewLyricRuby(r.RubyString, r.FedX)
			if err != nil {
				return nil, fme.LyricColorPicker{}, err
			}
			rubyString = append(rubyString, *ruby)
		}

		// body
		body, err := fme.NewLyricBody(lyricString, rubyString)
		if err != nil {
			return nil, fme.LyricColorPicker{}, err
		}

		// header
		xPoint := il.Point.X
		yPoint := il.Point.Y

		bc := fme.NewColorFromRGB888(il.Colors.BeforeCharColor.Red,
			il.Colors.BeforeCharColor.Green, il.Colors.BeforeCharColor.Blue)
		ac := fme.NewColorFromRGB888(il.Colors.BeforeCharColor.Red,
			il.Colors.BeforeCharColor.Green, il.Colors.BeforeCharColor.Blue)
		bo := fme.NewColorFromRGB888(il.Colors.BeforeOutlineColor.Red,
			il.Colors.BeforeOutlineColor.Green, il.Colors.BeforeOutlineColor.Blue)
		ao := fme.NewColorFromRGB888(il.Colors.AfterOutlineColor.Red,
			il.Colors.AfterOutlineColor.Green, il.Colors.AfterOutlineColor.Blue)
		header, err := fme.NewLyricHeaderWithStandardColorPicker(body.CalcByteSize(), xPoint, yPoint, *bc, *ac, *bo, *ao)
		if err != nil {
			return nil, fme.LyricColorPicker{}, err
		}

		block, err := fme.NewLyricBlock(*header, *body)
		if err != nil {
			return nil, fme.LyricColorPicker{}, err
		}

		blocks = append(blocks, *block)
	}

	return blocks, fme.StandardColorPicker, err
}