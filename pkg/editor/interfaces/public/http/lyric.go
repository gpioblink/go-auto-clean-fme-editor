package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	commonHttp "github.com/gpioblink/go-auto-clean-fme-editor/pkg/common/http"
	"github.com/gpioblink/go-auto-clean-fme-editor/pkg/editor/application"
	"github.com/gpioblink/go-auto-clean-fme-editor/pkg/editor/domain/lyric"
	"net/http"
)

func AddRoutes(router *chi.Mux, service application.LyricService, repository lyric.Repository) {
	resource := lyricResource{service, repository}
	router.Post("/lyric/edit", resource.PostEdit)
	router.Get("/lyric", resource.GetAll)
}

type lyricResource struct {
	service    application.LyricService
	repository lyric.Repository
}

func (o lyricResource) GetAll(w http.ResponseWriter, r *http.Request) {
	lyrics, err := o.service.ListLyrics()
	if err != nil {
		_ = render.Render(w, r, commonHttp.ErrInternal(err))
		return
	}

	var view []lyricView
	for _, l := range lyrics {
		var lyrics lyricViewLyricString
		for _, lst := range l.Lyric() {
			lyrics = append(lyrics, lyricViewLyricChar{
				lst.Length(),
				lst.Char(),
			})
		}
		var rubyString lyricViewRubyString
		for _, r := range l.Ruby() {
			rubyString = append(rubyString, lyricViewRuby{
				r.FedX(),
				r.RubyString(),
			})
		}
		view = append(view, lyricView{
			Point: lyricViewPoint{l.Point().X(), l.Point().Y()},
			Colors: lyricViewColorPicker{
				lyricViewColorPickerColor{
					l.Colors().BeforeCharColor().Red(),
					l.Colors().BeforeCharColor().Green(),
					l.Colors().BeforeCharColor().Blue(),
				},
				lyricViewColorPickerColor{
					l.Colors().AfterCharColor().Red(),
					l.Colors().AfterCharColor().Green(),
					l.Colors().AfterCharColor().Blue(),
				},
				lyricViewColorPickerColor{
					l.Colors().BeforeOutlineColor().Red(),
					l.Colors().BeforeOutlineColor().Green(),
					l.Colors().BeforeOutlineColor().Blue(),
				},
				lyricViewColorPickerColor{
					l.Colors().AfterOutlineColor().Red(),
					l.Colors().AfterOutlineColor().Green(),
					l.Colors().AfterOutlineColor().Blue(),
				},
			},
			Lyric: lyrics,
			Ruby:  rubyString,
		})
	}

	render.Respond(w, r, view)
}

type lyricView struct {
	Point  lyricViewPoint       `json:"point"`
	Colors lyricViewColorPicker `json:"colors"`
	Lyric  lyricViewLyricString `json:"lyric"`
	Ruby   lyricViewRubyString  `json:"ruby"`
}

type lyricViewPoint struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type lyricViewColorPicker struct {
	BeforeCharColor    lyricViewColorPickerColor `json:"beforeCharColor"`
	AfterCharColor     lyricViewColorPickerColor `json:"afterCharColor"`
	BeforeOutlineColor lyricViewColorPickerColor `json:"beforeOutlineColor"`
	AfterOutlineColor  lyricViewColorPickerColor `json:"afterOutlineColor"`
}

type lyricViewColorPickerColor struct {
	Red   int `json:"red"`
	Green int `json:"green"`
	Blue  int `json:"blue"`
}

type lyricViewRubyString []lyricViewRuby

type lyricViewRuby struct {
	FedX       int    `json:"fedx"`
	RubyString string `json:"string"`
}

type lyricViewLyricString []lyricViewLyricChar

type lyricViewLyricChar struct {
	Length    int    `json:"length"`
	LyricChar string `json:"char"`
}

func (o lyricResource) PostEdit(w http.ResponseWriter, r *http.Request) {
	req := PostEditRequest{}
	if err := render.Decode(r, &req); err != nil {
		_ = render.Render(w, r, commonHttp.ErrInternal(err))
		return
	}

	var lyrics []application.EditLyricLyric
	for _, l := range req.LyricString {
		lyrics = append(lyrics, application.EditLyricLyric(l))
	}

	cmd := application.EditLyricCommand{
		Lyrics: lyrics,
		Index:  req.Index,
	}

	err := o.service.EditLyric(cmd)
	if err != nil {
		_ = render.Render(w, r, commonHttp.ErrBadRequest(err))
	}
}

type PostEditRequest struct {
	Index       int                 `json:"index"`
	LyricString []PostEditLyricChar `json:"lyric"`
	Ruby        []PostEditRuby      `json:"ruby"`
}

type PostEditRuby struct {
	FedX       int    `json:"fedx"`
	RubyString string `json:"string"`
}

type PostEditLyricChar struct {
	Length    int    `json:"len"`
	LyricChar string `json:"char"`
}
