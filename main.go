package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	var outputFileName string
	flag.StringVar(&outputFileName, "o", "", "Output file name")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: go run main.go -o <output_file> <directory>")
		return
	}

	dir := flag.Args()[0]

	var outputFile *os.File
	var err error
	if outputFileName != "" {
		outputFile, err = os.Create(outputFileName)
		if err != nil {
			fmt.Printf("Error creating output file: %v\n", err)
			return
		}
		defer outputFile.Close()
	}

	// Durchsuche das Ausgangsverzeichnis selbst
	err = processDirectory(dir, outputFile)
	if err != nil {
		fmt.Printf("Error processing directory %s: %v\n", dir, err)
		return
	}

	// Durchsuche alle Unterverzeichnisse rekursiv
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Wenn der Fehler auf "Zugriff verweigert" zurückzuführen ist, ignoriere ihn
			if os.IsPermission(err) {
				fmt.Printf("Skipping directory %s due to permission denied\n", path)
				return nil
			}
			return err
		}

		// Wenn die Datei ein Verzeichnis ist, ignorieren wir sie
		if info.IsDir() {
			return nil
		}

		return processFile(path, outputFile)
	})

	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", dir, err)
		return
	}
}

func processDirectory(dir string, outputFile *os.File) error {
	// Liste der Dateien im Verzeichnis lesen
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	// Verarbeite jede Datei im Verzeichnis
	for _, file := range files {
		if !file.IsDir() {
			err := processFile(filepath.Join(dir, file.Name()), outputFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func processFile(path string, outputFile *os.File) error {
	// Dateiinhalt lesen
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// MD5-Hash berechnen
	hash := md5.Sum(data)
	hashStr := hex.EncodeToString(hash[:])

	// Ausgabe des Hashes und des Dateipfades
	result := fmt.Sprintf("File: %s, MD5 Hash: %s\n", path, hashStr)
	fmt.Print(result)

	if outputFile != nil {
		if _, err := outputFile.WriteString(result); err != nil {
			fmt.Printf("Error writing to output file: %v\n", err)
			return err
		}
	}

	return nil
}
