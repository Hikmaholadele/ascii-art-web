# ASCII Art Web

## Description

ASCII Art Web is a web application written in Go that converts user input text into ASCII art using different banner styles. Users can enter text through a web interface and receive the generated ASCII art output directly in their browser.

## Features

* Convert text into ASCII art
* Multiple banner styles:

  * standard
  * shadow
  * thinkertoy
* Simple and responsive web interface
* Error handling for invalid inputs
* Built using Go's `net/http` package

---

## Project Structure

```text
ascii-art-web/
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── templates/
│   └── index.html
├── main.go
├── go.mod
└── README.md
```

---

