package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"pixur.org/pixur/tasks"
)

type PurgePicHandler struct {
	// embeds
	http.Handler

	// deps
	PixPath string
	DB      *sql.DB
}

// TODO: add tests
// TODO: Add csrf protection
func (h *PurgePicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestedRawPicID := r.FormValue("pic_id")
	var requestedPicId int64
	if requestedRawPicID != "" {
		if picId, err := strconv.Atoi(requestedRawPicID); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else {
			requestedPicId = int64(picId)
		}
	}

	var task = &tasks.PurgePicTask{
		DB:      h.DB,
		PixPath: h.PixPath,
		PicId:   requestedPicId,
	}
	runner := new(tasks.TaskRunner)
	if err := runner.Run(task); err != nil {
		returnTaskError(w, err)
		return
	}

	returnJSON(w, true)
}

func init() {
	register(func(mux *http.ServeMux, c *ServerConfig) {
		mux.Handle("/api/purgePic", &PurgePicHandler{
			DB:      c.DB,
			PixPath: c.PixPath,
		})
	})
}