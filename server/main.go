package main

import "net/http"
import "io/ioutil"

func main() {
    http.ListenAndServe("0.0.0.0:8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        body, err := ioutil.ReadAll(r.Body)
        defer r.Body.Close()

        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        if len(body) == 0 {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        w.Write(append([]byte("hi "), body...))
    }))
}
