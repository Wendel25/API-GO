package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"api/models"
)

func ReadFile(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("dadosAPI.txt")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var pessoas []models.Pessoa
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ", ")
		if len(parts) == 3 {
			pessoa := models.Pessoa{}
			for _, part := range parts {
				keyValue := strings.Split(part, ": ")
				if len(keyValue) == 2 {
					switch keyValue[0] {
					case "Nome":
						pessoa.Nome = keyValue[1]
					case "Idade":
						fmt.Sscanf(keyValue[1], "%d", &pessoa.Idade)
					case "Profissão":
						pessoa.Profissao = keyValue[1]
					}
				}
			}
			pessoas = append(pessoas, pessoa)
		}
	}

	if err := scanner.Err(); err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pessoas)
}
