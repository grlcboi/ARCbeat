package main

import "bufio"
import "strings"
import "os"
import "strconv"

// turn arcstat file into native data types
func collectData(path string) {
  inFile, err := os.Open(path)
  _ = err // for now ignore error

  defer inFile.Close() // close file when task complete
  scanner := bufio.NewScanner(inFile) // create a scanner object
	scanner.Split(bufio.ScanLines) // split the file by lines
  
	arcMap := make(map[string]map[string]int) // create our master dictionary

  // handle each line
  for scanner.Scan() {
  	line := scanner.Text()
    fields := strings.Fields(line)

    dataName := fields[0]  // save our "name" column

    dataType, err := strconv.Atoi(fields[1]) // cast "type" column to int and save
    _ = err // for now ignore error

    dataValue, err := strconv.Atoi(fields[2]) // cast "data" column to int and save
    _ = err // for now ignore error

    dataMap := make(map[string]int)

    dataMap["type"] = dataType
    dataMap["value"] = dataValue

    arcMap[dataName] = dataMap


  }
}

func main() {
    
    collectData( "./arcstats")

    os.Exit(0)
}
