package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"api/models"
)

func WriteFile(w http.ResponseWriter, r *http.Request) {
	var pessoa models.Pessoa

	err := json.NewDecoder(r.Body).Decode(&pessoa)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	file, err := os.OpenFile("dadosAPI.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	newLine := fmt.Sprintf("Nome: %s, Idade: %d, Profissão: %s\n", pessoa.Nome, pessoa.Idade, pessoa.Profissao)
	if _, err := file.WriteString(newLine); err != nil {
		http.Error(w, "Unable to write to file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Data saved successfully")
}
