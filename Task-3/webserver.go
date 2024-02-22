package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

// Joke represents the structure for storing joke data
type Joke struct {
	ID        string `json:"id"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

var (
	// jokes holds the current list of jokes in memory
	jokes = []Joke{
		// Initial jokes data can be placed here
	}

	// jokesMutex protects write operations to the jokes slice
	jokesMutex sync.Mutex

	// highestID keeps track of the last ID assigned to a joke
	highestID int
)

// setDefaultHeaders sets common headers for JSON responses
func setDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// sendErrorResponse sends a standardized error response in JSON format
func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// createJoke handles POST requests to create a new joke
func createJoke(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(w)
	var newJoke Joke
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(reqBody, &newJoke); err != nil {
		sendErrorResponse(w, "Error parsing joke data", http.StatusBadRequest)
		return
	}

	// Locking for write operation, including ID generation
	jokesMutex.Lock()
	highestID++                          // Increment the highest ID for each new joke
	newJoke.ID = strconv.Itoa(highestID) // Assign the incremented ID to the new joke
	jokes = append(jokes, newJoke)
	jokesMutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newJoke)
}

// getJokeByID handles GET requests to retrieve a joke by its ID
func getJokeByID(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(w)
	jokeID := mux.Vars(r)["id"]
	for _, joke := range jokes {
		if joke.ID == jokeID {
			json.NewEncoder(w).Encode(joke)
			return
		}
	}
	sendErrorResponse(w, "Joke not found", http.StatusNotFound)
}

// updateJoke handles PATCH requests to update an existing joke
func updateJoke(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(w)
	jokeID := mux.Vars(r)["id"] // Get the ID from the URL path
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var updatedJoke Joke
	if err := json.Unmarshal(reqBody, &updatedJoke); err != nil {
		sendErrorResponse(w, "Invalid joke data", http.StatusBadRequest)
		return
	}

	// Since IDs are immutable and managed by the server, ignore the ID in the request body.
	// Instead, use the jokeID from the URL path to update the joke.
	updatedJoke.ID = jokeID // Ensure the updated joke has the correct ID

	jokesMutex.Lock()         // Lock for write operation
	defer jokesMutex.Unlock() // Ensure unlocking even if the function returns early

	found := false
	for i, joke := range jokes {
		if joke.ID == jokeID {
			jokes[i] = updatedJoke // Update the joke
			found = true
			break
		}
	}

	if !found {
		sendErrorResponse(w, "Joke not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedJoke) // Send back the updated joke
}

// deleteJokeByID handles DELETE requests to remove a joke by its ID
func deleteJokeByID(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(w)
	jokeID := mux.Vars(r)["id"]

	jokesMutex.Lock()         // Lock for write operation
	defer jokesMutex.Unlock() // Ensure unlocking

	for i, joke := range jokes {
		if joke.ID == jokeID {
			// Remove the joke from the slice
			jokes = append(jokes[:i], jokes[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("The joke with ID %s has been deleted successfully", jokeID)})
			return
		}
	}

	sendErrorResponse(w, "Joke not found", http.StatusNotFound)
}

// getAllJokes handles GET requests to list all jokes
func getAllJokes(w http.ResponseWriter, r *http.Request) {
	setDefaultHeaders(w)
	json.NewEncoder(w).Encode(jokes)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Using a default port if none is specified
		log.Printf("PORT environment variable not set, defaulting to %s", port)
	}

	router := mux.NewRouter().StrictSlash(true) // Redirect trailing slash
	// Setup route handlers
	router.HandleFunc("/jokes", createJoke).Methods("POST")
	router.HandleFunc("/jokes", getAllJokes).Methods("GET")
	router.HandleFunc("/jokes/{id}", getJokeByID).Methods("GET")
	router.HandleFunc("/jokes/{id}", updateJoke).Methods("PATCH")
	router.HandleFunc("/jokes/{id}", deleteJokeByID).Methods("DELETE")

	// Listen on all interfaces
	address := "0.0.0.0:" + port
	log.Printf("Server is listening on %s", address)
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
