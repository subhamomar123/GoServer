package main

import (
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/api/cmd", handleCommand)
	log.Fatal(http.ListenAndServe(":8080", addCorsHeaders(http.DefaultServeMux)))
}

func handleCommand(w http.ResponseWriter, r *http.Request) {
	cmdStr := r.FormValue("command") // Get the command from the query parameter "command"
	if cmdStr == "" {
		// If the command is not provided in the query parameter, check the request body
		cmdStr = r.URL.Query().Get("command")
	}
	if cmdStr == "" {
		http.Error(w, "Command not provided", http.StatusBadRequest)
		return
	}

	cmdParts := strings.Fields(cmdStr) // Split the command string into parts
	cmd := exec.Command("cmd", "/C", cmdParts[0], cmdParts[1:]...)
	output, err := cmd.Output()

	if err != nil {
		// If the command execution failed, return an error response
		if exitErr, ok := err.(*exec.ExitError); ok {
			http.Error(w, exitErr.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(output)
}

func addCorsHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add the necessary CORS headers to allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight (OPTIONS) requests
		if r.Method == http.MethodOptions {
			return
		}

		handler.ServeHTTP(w, r)
	})
}
