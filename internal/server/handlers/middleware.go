package handlers

import (
	"fmt"
	"net/http"

	"github.com/ChipArtem/Metric/internal/server/validator"

	"github.com/gorilla/mux"
)

func (m *metricHandler) MiddlewareCheckHost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/" {
			vars := mux.Vars(r)
			mType := vars["mtype"]
			fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n")
			if !validator.IsMType(mType) {
				http.Error(w, ``, http.StatusNotImplemented)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
