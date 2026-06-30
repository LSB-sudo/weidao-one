package server

import "net/http"

func (s *Server) handleViewerStatic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "web/viewer-static.html")
}
