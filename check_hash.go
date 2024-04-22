package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"regexp"
	"flag"
)

func loadMD5Hashes(filename string) ([]string, error) {
	// Array für MD5-Hashes initialisieren
	var md5Hashes []string

	// Textdatei öffnen
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Öffnen der Datei: %v", err)
	}
	defer file.Close()

	// Scanner zum Lesen der Datei initialisieren
	scanner := bufio.NewScanner(file)

	// Durch jede Zeile der Datei iterieren
	for scanner.Scan() {
		line := scanner.Text()

		// Überprüfen, ob die Zeile mit "#" beginnt (Kommentarzeile)
		if strings.HasPrefix(line, "#") {
			continue // Diese Zeile ignorieren
		}

		// Zeile anhand von Kommas aufteilen
		parts := strings.Split(line, ",")

		// Überprüfen, ob genügend Teile vorhanden sind und ob das dritte Element vorhanden ist
		if len(parts) >= 3 {
			// MD5-Hash aus dem dritten Element extrahieren (ohne Anführungszeichen)
			md5 := strings.TrimSpace(strings.ReplaceAll(parts[2], `"`, ""))
			md5Hashes = append(md5Hashes, md5)
		}
	}

	// Überprüfen, ob ein Fehler beim Lesen der Datei aufgetreten ist
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Fehler beim Lesen der Datei: %v", err)
	}

	return md5Hashes, nil
}

func extractMD5FromFile(filename string) ([]string, error) {
    // Öffnen der Datei
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Ein Array zum Speichern der extrahierten MD5-Hashes
    var md5Hashes []string

    // Regex-Pattern für MD5-Hashes
    md5Regex := regexp.MustCompile(`\b([a-fA-F0-9]{32})\b`)

    // Durchlaufen der Datei Zeile für Zeile
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        // Suche nach MD5-Hashes in der Zeile
        matches := md5Regex.FindAllString(line, -1)
        // Füge gefundene Hashes dem Array hinzu
        for _, match := range matches {
            md5Hashes = append(md5Hashes, match)
        }
    }

    // Überprüfe auf Fehler beim Scanner
    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return md5Hashes, nil
}




func main() {


	fmt.Println()
	
	fmt.Println("************* Hash-Check *************")
	
	
	
	var inputFileName string
	flag.StringVar(&inputFileName, "i", "", "Input file name")
	flag.Parse()

	if inputFileName == "" {
		fmt.Println("Usage: go run main.go -i <input_file>")
		return
	}


	// Laden der MD5-Hashes aus der CSV-Datei
	csvFilename := "full.csv"
	md5HashesFromCSV, err := loadMD5Hashes(csvFilename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Anzahl der geladenen MD5-Hashes ausgeben
	fmt.Printf("Anzahl der geladenen MD5-Hashes: %d\n", len(md5HashesFromCSV))

	// Meldung ausgeben, dass alle MD5-Hashes aus der CSV-Datei geladen wurden
	fmt.Println("Alle MD5-Hashes aus der CSV-Datei wurden erfolgreich geladen.")






	// Extrahiere MD5-Hashes aus der Textdatei
	textFilename := inputFileName
	md5HashesFromText, err := extractMD5FromFile(textFilename)
	if err != nil {
		fmt.Println("Fehler beim Extrahieren der MD5-Hashes:", err)
		return
	}
	
	
	
/*
	// Ausgabe der extrahierten MD5-Hashes
	fmt.Println("Extrahierte MD5-Hashes aus der Textdatei:")
	for _, hash := range md5HashesFromText {
		fmt.Println(hash)
	}
*/
	

	// Überprüfe, ob es Übereinstimmungen zwischen den Hashes gibt
	fmt.Println()
	fmt.Println("Malizioese Hashes gefunden:")
	for _, textHash := range md5HashesFromText {
		for _, csvHash := range md5HashesFromCSV {
			if textHash == csvHash {
				fmt.Println(textHash)
				break
			}
		}
	}



}

