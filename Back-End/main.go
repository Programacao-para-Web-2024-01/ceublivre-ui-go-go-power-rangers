package main

import (
    "fmt"           
    "html/template" 
    "net/http"      
)

func main() {

    http.HandleFunc("/", serveTemplate)
    
    http.HandleFunc("/esqueceu-senha", forgotPasswordHandler)

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    
    fmt.Println("Server started at :8080")
    
    http.ListenAndServe(":8080", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

func forgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost 
        email := r.FormValue("email")
        fmt.Fprintf(w, "Password recovery email sent to %s", email)
        return
    }
    http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}
