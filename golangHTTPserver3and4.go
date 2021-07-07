package main

import (
    "net/http"
    "os"
    "path"
)

func FileServerWithCustom404(fs http.FileSystem) http.Handler {
    fsh := http.FileServer(fs)
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        _, err := fs.Open(path.Clean(r.URL.Path))
        if os.IsNotExist(err) {
            http.NotFound(w, r)
            return
        }
        fsh.ServeHTTP(w, r)
    })
}

func main() {

    http.Handle("/", FileServerWithCustom404(http.Dir("./views")))
    http.ListenAndServe(":5000", nil)
}
