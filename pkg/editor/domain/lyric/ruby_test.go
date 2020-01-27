package lyric_test

import (
	"github.com/gpioblink/go-auto-clean-fme-editor/pkg/editor/domain/lyric"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRuby(t *testing.T) {
	testCases := []struct {
		TestName    string
		ExpectedErr bool
		FedX        int
		RubyString  string
	}{
		{
			TestName:    "基本",
			ExpectedErr: false,
			RubyString:  "あいうえお",
			FedX:        12,
		},
		{
			TestName:    "基本",
			ExpectedErr: false,
			RubyString:  "ちよやちよ",
			FedX:        75,
		},
		{
			TestName:    "基本",
			ExpectedErr: true,
			RubyString:  "아아아아",
			FedX:        0,
		},
	}

	for _, c := range testCases {
		t.Run(c.TestName, func(t *testing.T) {
			ruby, err := lyric.NewRuby(c.FedX, c.RubyString)
			if c.ExpectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, c.FedX, ruby.FedX())
				assert.EqualValues(t, c.RubyString, ruby.RubyString())
			}
		})
	}
}
