package main

import (
 "fmt"
 "net/http"
 "time"
)

func main() {
 startTime := time.Now()

 // Define a rota GET
 http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
   http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
   return
  }

  // Responde com um JSON simples
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  response := `{"message": "Bem-vindo à API em Go!"}`
  fmt.Fprintln(w, response)
 })

 // Inicia o servidor HTTP na porta 8090
 fmt.Printf("Servidor rodando em http://localhost:8090\n")
 endTime := time.Since(startTime)
 fmt.Printf("Tempo de início da aplicação: %s\n", endTime)
 err := http.ListenAndServe(":8090", nil)
 if err != nil {
  fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
 }
}