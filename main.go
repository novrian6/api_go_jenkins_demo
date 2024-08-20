package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD

<<<<<<< HEAD

	fmt.Fprintfa(w, "Hello1, World apa kabar arif")
	fmt.Fprintfa(w, "Baris ini ditulis oleh arif")

=======
	fmt.Fprintfb(w, "Hello2, World apa kabar nova")
	fmt.Fprintfb(w, "Baris, ini ditulis oleh nova")
        fmt.Fprintf(w, "Hello, World apa kabar sumar")
        fmt.Fprintf(w, "Baris ini ditulis oleh sumar")
=======
	fmt.Fprintfc(w, "Hello3, World apa kabar raki")
	fmt.Fprintfc(w, "Baris ini, ditulis oleh raki")
>>>>>>> bb4de33e813a5665ec59557ad07c13eca47ceefc
>>>>>>> d2cf0607c6a7c570b046500351bb275228d12e10

}

func main() {
	port := ":8070"

	// Print a message indicating the server is starting
	fmt.Printf("Starting server on port%s\n", port)
	http.HandleFunc("/", handler)

	http.ListenAndServe(port, nil)
	fmt.Printf("Appication stop\n")
}
