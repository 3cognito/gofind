# gofind

A simple, efficient tool to search for specific text occurrences within files in a directory. This tool can be used both as a **Go package** and a **CLI tool**.

---

## Features
- Traverse directories and search for text in files.
- Count occurrences of the specified text in all files.
- Exclude file extensions
- Simple, lightweight, and easy to integrate.

---

## Installation

### **As a Go Package**
To use this tool in your Go project, install it via:

```bash
go get github.com/<3cognito>/<gofind>
```

Import the package into your project:

```go
import "github.com/<3cognito>/<gofind>"
```

### **As a CLI Tool**
If you’ve built the CLI executable and published it, users can install it via:

```bash
go install github.com/<3cognito>/<gofind>/cmd/gofind@latest
```

Run the tool directly from the terminal:

```bash
gofind -text="search_term" -dir="/path/to/directory"
```

---

## Usage

### **Package Example**
Use the tool in your Go code to search for text:

```go
package main

import (
	"fmt"
	"log"
	"github.com/3cognito/gofind"
)

func main() {
	count, err := gofind.CountOccurrences("example.txt", "search_term", ".gitignore,.git")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Occurrences found: %d
", count)
}
```

### **CLI Example**
From the command line, you can run:

```bash
mycli -text="search_term" -dir="/path/to/directory"
```

This will output the total occurrences of `search_term` across all files in the specified directory.

---

## Future Features
- Parallel processing optimizations.
- File type filtering (`.txt`, `.log`, etc.).

---

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

---

## License
This project is licensed under the MIT License.
