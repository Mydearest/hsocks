package cli

import "net/http"

func dispatcher(mux *http.ServeMux){
	mux.HandleFunc("/" ,handle)
}

func handle(writer http.ResponseWriter, res *http.Request){

}
