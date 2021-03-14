package server

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/edspc/go-libs/web/data"
)

func ServeOneRequest(location data.ServerLocation) data.RequestData {

	srv := &http.Server{Addr: location.Server.GetServer()}
	var data data.RequestData

	http.HandleFunc(location.Uri, func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Fatal(err)
		}

		data.Body = body
		data.Params = r.URL.Query()

		go func() {
			cxt := context.Background()
			if err := srv.Shutdown(cxt); err != http.ErrServerClosed {
				log.Fatalf("Shutdown(): %v", err)
			}
		}()
	})

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe(): %v", err)
	}

	return data
}
