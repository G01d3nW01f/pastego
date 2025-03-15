package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// XOR encryption
const xorKey = 0x55

// XOR decrypt
func xorDecrypt(data []byte) []byte {
	decrypted := make([]byte, len(data))
	for i, b := range data {
		decrypted[i] = b ^ xorKey
	}
	return decrypted
}

// RAW data get from pastebin
func fetchPastebinRaw(pastebinURL string) ([]byte, error) {
	// RAW formed into url （ex: https://pastebin.com/xxxxxxxx → https://pastebin.com/raw/xxxxxxxx）
	rawURL := strings.Replace(pastebinURL, "pastebin.com/", "pastebin.com/raw/", 1)
	if !strings.HasPrefix(rawURL, "https://pastebin.com/raw/") {
		return nil, fmt.Errorf("invalid Pastebin URL: %s", pastebinURL)
	}

	// HTTP GET
	resp, err := http.Get(rawURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Pastebin data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Pastebin returned status: %d", resp.StatusCode)
	}

	// data read
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Pastebin response: %v", err)
	}

	return data, nil
}

func main() {
	// argument check
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run decrypt.go <pastebin_url> <output_filename>")
		fmt.Println("Example: go run decrypt.go https://pastebin.com/xxxxxxxx restored.txt")
		os.Exit(1)
	}

	pastebinURL := os.Args[1]
	outputFilename := os.Args[2]

	// get encryption from Pastebin 
	encryptedData, err := fetchPastebinRaw(pastebinURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching Pastebin data: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Fetched %d bytes from Pastebin\n", len(encryptedData))

	// XOR decrypt
	decryptedData := xorDecrypt(encryptedData)
	fmt.Printf("Decrypted %d bytes of data\n", len(decryptedData))

	// save to file
	err = os.WriteFile(outputFilename, decryptedData, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file %s: %v\n", outputFilename, err)
		os.Exit(1)
	}

	// success message
	fmt.Printf("Data successfully decrypted and saved to %s\n", outputFilename)
}
