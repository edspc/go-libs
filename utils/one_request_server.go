package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ServeOneRequest(location URI) RequestData {

	srv := &http.Server{Addr: location.Host.GetHostString()}
	var data RequestData

	http.HandleFunc(location.Uri, func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Fatal(err)
		}

		data.Body = body
		data.Params = r.URL.Query()

		fmt.Fprintf(w, "<script type=\"text/javascript\">window.close();</script>")

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
