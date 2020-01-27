package http

import (
	"encoding/base64"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	commonHttp "github.com/gpioblink/go-auto-clean-fme-editor/pkg/common/http"
	"github.com/gpioblink/go-auto-clean-fme-editor/pkg/fme/application"
	"github.com/gpioblink/go-auto-clean-fme-editor/pkg/fme/converterDomain/fme"
	"net/http"
)

func AddRoutes(router *chi.Mux, service application.FmeService, repository fme.Repository) {
	resource := fmeResource{service, repository}
	router.Post("/fme/import", resource.PostImport)
	router.Get("/fme/export", resource.GetExport)
}

type fmeResource struct {
	service    application.FmeService
	repository fme.Repository
}

type PostImportRequest struct {
	FmeBase64 string `json:"fme"`
}

func (o fmeResource) PostImport(w http.ResponseWriter, r *http.Request) {
	req := PostImportRequest{}
	if err := render.Decode(r, &req); err != nil {
		_ = render.Render(w, r, commonHttp.ErrInternal(err))
		return
	}
	// TODO: こういうusecaseの入力っての1つでもcmd化してまとめると分かりやすいかも
	fmeBytes, err := base64.StdEncoding.DecodeString(req.FmeBase64)
	if err != nil {
		_ = render.Render(w, r, commonHttp.ErrBadRequest(err))
		return
	}

	err = o.service.ImportFme(fmeBytes)
	if err != nil {
		_ = render.Render(w, r, commonHttp.ErrInternal(err))
		return
	}
}

type GetExportView struct {
	FmeBase64 string `json:"fme"`
}

func (o fmeResource) GetExport(w http.ResponseWriter, r *http.Request) {
	lyrics, err := o.service.ExportFme()
	if err != nil {
		_ = render.Render(w, r, commonHttp.ErrInternal(err))
		return
	}
	lyricString := base64.StdEncoding.EncodeToString(lyrics)
	view := GetExportView{lyricString}
	render.Respond(w, r, view)
}
