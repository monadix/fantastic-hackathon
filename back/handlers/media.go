package handlers

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func (s *HandlersServer) HandleMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	path, has := vars["path"]
	if !has {
		ErrorMap(w, http.StatusBadRequest, map[string]interface{}{
			"type":    "data",
			"reason":  "path",
			"explain": ErrExplainInvalidPhotoURL,
		})
		return
	}

	pic, err := os.ReadFile("../media/" + path)
	if err != nil {
		ErrorMap(w, http.StatusBadRequest, map[string]interface{}{
			"type":    "media",
			"reason":  "path",
			"explain": ErrExplainInvalidMedia,
		})
		return
	}

	w.Write(pic)
}
