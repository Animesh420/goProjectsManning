package newgame

import "net/http"

func Handle(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Creating a new game\n"))

}
