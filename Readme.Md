# gofind

A simple, efficient tool to search for specific text occurrences within files in a directory. This tool can be used both as a **Go package**.

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
go get github.com/3cognito/gofind
```

Import the package into your project:

```go
import "github.com/3cognito/gofind"
```

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
	count := gofind.CountOccurrences("example.txt", "search_term", ".gitignore,.git")
	fmt.Printf("Occurrences found: %d
", count)
}
```

This will output the total occurrences of `search_term` across all files in the specified directory except those with the extensions `.gitignore` and `.git`.

---

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

---

## License
This project is licensed under the MIT License.
