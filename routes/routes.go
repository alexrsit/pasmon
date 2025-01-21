package routes

import (
	"net/http"
	"pasmon/api"
)

func InitRoutes(client *http.Client) {
	http.HandleFunc("/types", func(w http.ResponseWriter, r *http.Request) {
		err := api.GetTypes(client, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.HandleFunc("/oidc_rp", func(w http.ResponseWriter, r *http.Request) {
		err := api.GetOidcRp(client, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.HandleFunc("/authenticators", func(w http.ResponseWriter, r *http.Request) {
		err := api.GetAuthenticators(client, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.HandleFunc("/resources", func(w http.ResponseWriter, r *http.Request) {
		err := api.GetMetaData(client, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.HandleFunc("/samlsp", func(w http.ResponseWriter, r *http.Request) {
		err := api.GetSp(client, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":3000", nil)
}
