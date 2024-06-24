package main

import (
    "fmt"           // Pacote para formatação de entrada e saída
    "html/template" // Pacote para manipulação de templates HTML
    "net/http"      // Pacote para construção de servidores HTTP
)

func main() {
    // Define o roteamento para a função serveTemplate quando o caminho raiz (/) for acessado
    http.HandleFunc("/", serveTemplate)
    
    // Define o roteamento para a função forgotPasswordHandler quando /esqueceu-senha for acessado
    http.HandleFunc("/esqueceu-senha", forgotPasswordHandler)
    
    // Serve arquivos estáticos a partir do diretório "static"
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    
    // Imprime uma mensagem no console indicando que o servidor foi iniciado
    fmt.Println("Server started at :8080")
    
    // Inicia o servidor HTTP na porta 8080
    http.ListenAndServe(":8080", nil)
}

// Função para servir o template HTML
func serveTemplate(w http.ResponseWriter, r *http.Request) {
    // Faz o parsing do arquivo index.html
    tmpl, err := template.ParseFiles("index.html")
    if err != nil {
        // Em caso de erro, retorna um erro 500 (Internal Server Error)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // Renderiza o template e envia para o cliente
    tmpl.Execute(w, nil)
}

// Função para lidar com a recuperação de senha
func forgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
    // Verifica se o método da requisição é POST
    if r.Method == http.MethodPost {
        // Obtém o valor do campo "email" do formulário
        email := r.FormValue("email")
        // Aqui, você lidaria com o processo de redefinição de senha (por exemplo, enviando um e-mail)
        fmt.Fprintf(w, "Password recovery email sent to %s", email)
        return
    }
    // Se o método não for POST, retorna um erro 405 (Method Not Allowed)
    http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}
