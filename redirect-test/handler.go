package function

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//Handle handles the function call to function
func Handle(w http.ResponseWriter, r *http.Request) {
	redirect := r.FormValue("redirect") != ""
	if redirect {
		http.Redirect(w, r, r.FormValue("redirect"), http.StatusTemporaryRedirect)
		return
	}

	for k, v := range r.Header {
		fmt.Printf("k: %s, v: %s\n", k, v)
	}

	if isValidSignature(r, "ffde76180518da0a7b31b80993697412cf9b2cf8") {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK done"))
		return
	}

	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("forbidden"))
}

func isValidSignature(r *http.Request, key string) bool {
	// Assuming a non-empty header
	gotHash := strings.SplitN(r.Header.Get("X-Hub-Signature"), "=", 2)
	if gotHash[0] != "sha1" {
		return false
	}

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Cannot read the request body: %s\n", err)
		return false
	}

	hash := hmac.New(sha1.New, []byte(key))
	if _, err := hash.Write(b); err != nil {
		log.Printf("Cannot compute the HMAC for request: %s\n", err)
		return false
	}

	expectedHash := hex.EncodeToString(hash.Sum(nil))
	log.Println("EXPECTED HASH:", expectedHash)
	return gotHash[1] == expectedHash
}
