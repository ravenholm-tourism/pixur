package handlers

import (
	"html/template"
	"net/http"

	"pixur.org/pixur/api"
)

type indexData struct {
	baseData

	Pic []*api.Pic

	NextID, PrevID string
}

var indexTpl = template.Must(template.Must(paneTpl.Clone()).ParseFiles("fe/tpl/index.html"))

type indexHandler struct {
	c  api.PixurServiceClient
	pt paths
}

func (h *indexHandler) static(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		httpError(w, &HTTPErr{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}
	id := r.FormValue(h.pt.pr.IndexPic())
	_, isPrev := r.Form[h.pt.pr.IndexPrev()]
	req := &api.FindIndexPicsRequest{
		StartPicId: id,
		Ascending:  isPrev,
	}

	res, err := h.c.FindIndexPics(r.Context(), req)
	if err != nil {
		httpError(w, err)
		return
	}
	var prevID string
	var nextID string
	if !isPrev {
		if len(res.Pic) >= 2 {
			nextID = res.Pic[len(res.Pic)-1].Id
		}
		if id != "" {
			prevID = id
		}
	} else {
		if len(res.Pic) >= 2 {
			prevID = res.Pic[len(res.Pic)-1].Id
		}
		if id != "" {
			nextID = id
		}
	}

	if isPrev {
		for i := 0; i < len(res.Pic)/2; i++ {
			res.Pic[i], res.Pic[len(res.Pic)-i-1] = res.Pic[len(res.Pic)-i-1], res.Pic[i]
		}
	}
	xsrfToken, _ := xsrfTokenFromContext(r.Context())
	data := indexData{
		baseData: baseData{
			Title:       "Index",
			Paths:       h.pt,
			Params:      h.pt.pr,
			XsrfToken:   xsrfToken,
			SubjectUser: subjectUserOrNilFromCtx(r.Context()),
		},
		Pic:    res.Pic,
		NextID: nextID,
		PrevID: prevID,
	}
	if err := indexTpl.Execute(w, data); err != nil {
		httpError(w, err)
		return
	}
}
