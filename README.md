# Ascii Art Web

This small project takes the main [Ascii Art Project](https://zone01normandie.org/git/faoudia/ascii-art) and runs it on the web. The server is fully written in go and uses HTML templates to display content on the front-end.

## Authors

- Fares Aoudia
- Maxime TREBERT

## Code

Two main handlers are used: the homehandler and the posthandler. The home handler displays the home page which contains an HTML form to submit the input and the output of the Ascii Art.

The post handler takes the input typed on the home page and turns it into Ascii Art before sending it back to the front-end.

The server handles error codes 404, 400 and 500 with a specific template for each case.

## Usage: How to Run

1. **Clone the repository:**

    ```bash
    git clone https://zone01normandie.org/git/faoudia/ascii-art-web.git
    cd ascii-art-web
    ```

2. **Run the application:**

    ```bash
    go run .
    ```
    - if the server runs correctly, you should see this:
    ```bash
    ascii-web-server:00:58:12  Server started at http://localhost:8080                                                          
    ```

3. **Open your browser and navigate to:**

    ```
    http://localhost:8080
    ```

4. **Use the application:**

    - Enter text in the input field.
    - Select a banner style from the dropdown menu.
    - Click "Get ASCII Art" to generate and display the ASCII art.

    To use the existing test file `main_test.go` to test the provided Go program, you can follow these steps:

## Test : How to Test 

1. **Navigate to the Test File**: Locate the `main_test.go` file in the repository.

2. **Review Test Cases**:

    - **TestHomeHandler**: Verifies that the home handler returns a status code of 200 (OK).
    - **TestNotFoundHandler**: Verifies that a non-existent route returns a status code of 404 (Not Found).
    - **TestPostHandler**: Verifies that a valid POST request to the `/ascii` endpoint returns a status code of 302 (Found).
    - **TestBadRequestHandler**: Verifies that a GET request to the `/ascii` endpoint returns a status code of 400 (Bad Request).
    - **TestInternalServerErrorHandler**: Verifies that a malformed POST request to the `/ascii` endpoint returns a status code of 500 (Internal Server Error).

3. **Run Tests**: Ensure that you are in the directory containing the `main_test.go` file and run:

    ```
    go test -v
    ```

    The `-v` flag stands for verbose output, which provides more detailed information about the tests being run.

## Implementation Details: Algorithm

### Arborescence

- `main.go`
- `main_test.go`
- `core/`
  - `fileopen.go`
  - `page.go`
- `ascii-art/`
  - `ascii-art.go`
  - `fileopen.go`
  - `getbanner.go`
  - `getletter.go`
  - `getword.go`
- `templates/`
  - `index.html`
  - `500.html`
  - `404.html`
  - `400.html`
- `banners/`
  - `standard.txt`
  - `shadow.txt`
  - `thinkertoy.txt`

### main.go

The `main.go` file sets up the HTTP server and handles routing for the home page and the ASCII art generation. It includes functions for handling different HTTP requests and error responses.

#### Constants
```go
const RED = "\033[31;1m"
const GREEN = "\033[32;1m"
const YELLOW = "\033[33;1m"
const NONE = "\033[0m"
```
These constants define ANSI color codes for colored terminal output.

#### Struct
```go
type PageData struct {
    Lines []string
}
```
`PageData` is used to pass data (lines of text) to HTML templates.

#### Main Function
```go
func main() {
    log.SetFlags(log.Ltime)
    log.SetPrefix("ascii-web-server:")

    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/ascii", postHandler)

    log.Println(GREEN, "Server started at http://localhost:8080", NONE)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```
- Configures logging to include timestamps and a custom prefix.
- Registers handlers for the root ("/") and "/ascii" routes.
- Starts the server on port 8080.

#### Handlers

##### Home Handler
```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        log.Printf("%v Bad request %v on %v page%v\n", RED, r.Method, r.URL.Path, NONE)
        badRequestHandler(w)
        return
    }

    if r.URL.Path != "/" {
        log.Printf("%v Tried to access unexistant route %v%v\n", YELLOW, r.URL.Path, NONE)
        notFoundHandler(w)
        return
    }

    t, err := template.ParseFiles("templates/index.html")
    if err != nil {
        log.Printf("%v Error parsing home template: %v%v", RED, err, NONE)
        internalServerErrorHandler(w)
        return
    }

    f, err := os.ReadFile("content.txt")
    if err != nil {
        log.Printf("%v Error reading content file: %v%v", RED, err, NONE)
    }

    fileContent := string(f)
    data := PageData{
        Lines: strings.Split(fileContent, "\n"),
    }

    err = t.Execute(w, data)
    if err != nil {
        log.Printf("%v Error executing home template: %v%v", RED, err, NONE)
        internalServerErrorHandler(w)
        return
    }
}
```
- Verifies the request method and URL path.
- Parses and executes the home page template, passing file content as data.
- Handles errors by calling appropriate error handlers.

##### Post Handler
```go
func postHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        log.Printf("%v Bad request %v on %v page%v\n", RED, r.Method, r.URL.Path, NONE)
        badRequestHandler(w)
        return
    }

    err := r.ParseForm()
    if err != nil {
        log.Printf("%v Error parsing data form: %v%v", RED, err, NONE)
        internalServerErrorHandler(w)
        return
    }

    input := r.FormValue("input")
    style := r.FormValue("banner")

    input = strings.Replace(input, "\r\n", "\n", -1)

    if style == "" {
        log.Printf("%v No banner provided: style: %s%v\n", RED, style, NONE)
        internalServerErrorHandler(w)
        return
    }

    output := asciiart.GetAscii(input, style)

    err = core.Save(output)
    if err != nil {
        log.Printf("%v Error saving output: %v%v", RED, err, NONE)
        internalServerErrorHandler(w)
        return
    }

    log.Printf("%v POST request on /ascii successful %v", GREEN, NONE)

    http.Redirect(w, r, "/", http.StatusFound)
}
```
- Verifies the request method.
- Parses form data for the ASCII input and style.
- Generates ASCII art and saves it using the core package.
- Redirects to the home page upon success.

##### Error Handlers
These handlers render specific error pages when errors occur.

- **Not Found (404)**
    ```go
    func notFoundHandler(w http.ResponseWriter) {
        w.WriteHeader(http.StatusNotFound)
        t, err := template.ParseFiles("templates/404.html")
        if err != nil {
            log.Printf("%v Error executing 404 template: %v%v", RED, err, NONE)
            internalServerErrorHandler(w)
            return
        }
        err = t.Execute(w, nil)
        if err != nil {
            log.Printf("%v Error executing 404 template: %v%v", RED, err, NONE)
            internalServerErrorHandler(w)
            return
        }
    }
    ```
- **Bad Request (400)**
    ```go
    func badRequestHandler(w http.ResponseWriter) {
        w.WriteHeader(http.StatusBadRequest)
        t, err := template.ParseFiles("templates/400.html")
        if err != nil {
            log.Printf("%v Error executing 400 template: %v%v", RED, err, NONE)
            internalServerErrorHandler(w)
            return
        }
        err = t.Execute(w, nil)
        if err != nil {
            log.Printf("%v Error executing 400 template: %v%v", RED, err, NONE)
            internalServerErrorHandler(w)
            return
        }
    }
    ```
- **Internal Server Error (500)**
    ```go
    func internalServerErrorHandler(w http.ResponseWriter) {
        w.WriteHeader(http.StatusInternalServerError)
        t, err := template.ParseFiles("templates/500.html")
        if err != nil {
            log.Printf("%v Error executing 500 template: %v%v", RED, err, NONE)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
        err = t.Execute(w, nil)
        if err != nil {
            log.Printf("Error executing 500 template: %v", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
    }
    ```

Each handler ensures appropriate HTTP status codes are set and specific templates are rendered, logging any errors encountered.

### fileopen.go

The `fileopen.go` file contains a function to read the content of a file and return it as a string.

```go
package core

import (
	"os"
	"strings"
)

func FileOpen(filename string) string {
	f, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	content := strings.ReplaceAll(string(f), "\r\n", "\n")
	return content
}

```

### page.go

The `page.go` file defines the `Page` struct and functions to save and load page data.

```go
package core

import "os"

type Page struct {
	Title string
	Body  []byte
}

func Save(lines []string) error {
	str := ""
	for _, line := range lines {
		str += line + "\n"
	}
	return os.WriteFile("content.txt", []byte(str), 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename)
	return &Page{Title: title, Body: body}, nil
}
```

### ascii-art.go

The `ascii-art.go` file contains the main function to generate ASCII art from input text and a specified banner style.

```go
package asciiart

import (
	"fmt"
	"strings"
)

func GetAscii(input, style string) []string {
	bannerFile, err := GetBannerFile(style)
	if err != nil {
		fmt.Println("Error:", err)
		return []string{}
	}
	lines := make([]string, 0)
	words := strings.Split(input, "\n")

	for _, word := range words {
		if word == "" {
			lines = append(lines, "")
			continue
		}
		lines = append(lines, GetWord(word, bannerFile)...)
	}

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.ReplaceAll(lines[i], " ", "&nbsp;")
	}

	return lines
}
```

### getbanner.go

The `getbanner.go` file provides a function to get the path of the banner file based on the style argument.

```go
package asciiart

import "fmt"

func GetBannerFile(style string) (string, error) {
	switch style {
	case "standard":
		return "banners/standard.txt", nil
	case "shadow":
		return "banners/shadow.txt", nil
	case "thinkertoy":
		return "banners/thinkertoy.txt", nil
	default:
		return "", fmt.Errorf("unknown style: %s. Available styles: standard, shadow, thinkertoy", style)
	}
}
```

### getletter.go

The `getletter.go` file defines the function to retrieve the ASCII representation of a single character from a banner.

```go
package asciiart

import (
	"strings"
)

const LETTER_HEIGHT = 8

func GetLetter(content string, ascii int) string {
	if ascii == 32 {
		s := ""
		for i := 0; i < 8; i++ {
			if i != 7 {
				s += "    " + "\n"
				continue
			}
			s += "    "
		}
		return s
	}

	str := ""
	lines := strings.Split(content, "\n")

	place := ascii - 31
	times := (place - 1) * LETTER_HEIGHT
	beginning := (ascii - 30) + times

	for i := beginning; i < beginning+LETTER_HEIGHT; i++ {
		if i != (beginning+LETTER_HEIGHT)-1 {
			str += lines[i-1] + "\n"
		} else {
			str += lines[i-1]
		}
	}

	return str
}
```

### getword.go

The `getword.go` file provides the function to generate ASCII art for an entire word by combining the ASCII representations of each character.

```go
package asciiart

import (
	"strings"
)

func GetWord(input string, bannerFile string) []string {
	lines := make([]string, 8)
	content := FileOpen(bannerFile)

	for _, char := range input {
		c := strings.Split(GetLetter(content, int(char)), "\n")
		for i := 0; i < len(lines); i++ {
			lines[i] += c[i]
		}
	}

	return lines
}
```