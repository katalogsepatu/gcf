package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	sepatu "github.com/katalogsepatu/be_sepatu/module"
)

func init() {
	functions.HTTP("sepatu", kategori_sepatu)
}

func kategori_sepatu(w http.ResponseWriter, r *http.Request) {
	allowedOrigins := []string{"https://ksi-billboard.github.io", "http://127.0.0.1:5500", "http://127.0.0.1:5501"}
	origin := r.Header.Get("Origin")

	for _, allowedOrigin := range allowedOrigins {
		if allowedOrigin == origin {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			break
		}
	}

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, sepatu.TambahKategoriSepatu("PASETOPUBLICKEY", "MONGOSTRING", "sepatu_db", "kategorisepatu", r))
		return
	}
	if r.Method == http.MethodPut {
		fmt.Fprintf(w, sepatu.EditKategoriSepatu("PASETOPUBLICKEY", "MONGOSTRING", "sepatu_db", "kategorisepatu", r))
		return
	}
	if r.Method == http.MethodDelete {
		fmt.Fprintf(w, sepatu.DeleteKategoriSepatu("PASETOPUBLICKEY", "MONGOSTRING", "sepatu_db", "kategorisepatu", r))
		return
	}
	// Set CORS headers for the main request.
	fmt.Fprintf(w, sepatu.GetKategoriSepatu("MONGOSTRING", "sepatu_db", "kategorisepatu", r))

}
