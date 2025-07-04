// /home/${USER}/livestreampro/backend/cmd/api-gateway/main.go
package main

import (
        "encoding/json"
        "log"
        "net/http"
        "strings"
)

type Channel struct {
        ID   string `json:"id"`
        Name string `json:"name"`
}

func main() {
        http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
                w.WriteHeader(http.StatusOK)
                w.Write([]byte("ok"))
        })

        http.HandleFunc("/channels/", func(w http.ResponseWriter, r *http.Request) {
                id := strings.TrimPrefix(r.URL.Path, "/channels/")
                if id == "" {
                        http.Error(w, "missing id", http.StatusBadRequest)
                        return
                }

                ch := Channel{ID: id, Name: "Channel " + id}
                w.Header().Set("Content-Type", "application/json")
                json.NewEncoder(w).Encode(ch)
        })

	log.Println("api-gateway started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
