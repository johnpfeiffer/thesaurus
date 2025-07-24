package main

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

func mdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the file path from the query parameter
	filename := r.URL.Query().Get("file")
	if filename == "" {
		// List available markdown files if no specific file requested
		listMarkdownFiles(w, r)
		return
	}

	// Sanitize the filename to prevent directory traversal
	filename = filepath.Base(filename)
	if !strings.HasSuffix(filename, ".md") {
		http.Error(w, "Only .md files are allowed", http.StatusBadRequest)
		return
	}

	// Construct the full path
	filePath := filepath.Join("data", filename)

	// Read the markdown file
	content, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// Convert markdown to HTML using goldmark
	md := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(), // Allow raw HTML
		),
	)

	var buf bytes.Buffer
	if err := md.Convert(content, &buf); err != nil {
		http.Error(w, "Error converting markdown", http.StatusInternalServerError)
		return
	}

	// Render the HTML page
	tmpl := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="/static/style.css">
    <style>
        .markdown-content {
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
            line-height: 1.6;
        }
        .markdown-content pre {
            background-color: #f6f8fa;
            padding: 16px;
            overflow: auto;
            border-radius: 6px;
        }
        .markdown-content code {
            background-color: #f6f8fa;
            padding: 2px 4px;
            border-radius: 3px;
            font-family: monospace;
        }
        .markdown-content blockquote {
            border-left: 4px solid #ddd;
            margin: 0;
            padding-left: 16px;
            color: #666;
        }
        .back-link {
            margin-bottom: 20px;
        }
    </style>
</head>
<body>
    <div class="markdown-content">
        <div class="back-link">
            <a href="/md">← Back to file list</a> | <a href="/">← Home</a>
        </div>
        <h1>{{.Title}}</h1>
        {{.Content}}
    </div>
</body>
</html>`

	t, err := template.New("markdown").Parse(tmpl)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Content template.HTML
	}{
		Title:   strings.TrimSuffix(filename, ".md"),
		Content: template.HTML(buf.String()),
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
		return
	}
}

func listMarkdownFiles(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("data")
	if err != nil {
		http.Error(w, "Error reading data directory", http.StatusInternalServerError)
		return
	}

	var mdFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			mdFiles = append(mdFiles, file.Name())
		}
	}

	tmpl := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Markdown Files</title>
    <link rel="stylesheet" href="/static/style.css">
    <style>
        .file-list {
            max-width: 600px;
            margin: 50px auto;
            padding: 20px;
        }
        .file-list ul {
            list-style-type: none;
            padding: 0;
        }
        .file-list li {
            margin: 10px 0;
            padding: 10px;
            background-color: #f5f5f5;
            border-radius: 5px;
        }
        .file-list a {
            text-decoration: none;
            color: #333;
            font-weight: bold;
        }
        .file-list a:hover {
            color: #0066cc;
        }
    </style>
</head>
<body>
    <div class="file-list">
        <h1>Available Markdown Files</h1>
        <a href="/">← Back to Home</a>
        {{if .}}
        <ul>
            {{range .}}
            <li><a href="/md?file={{.}}">{{.}}</a></li>
            {{end}}
        </ul>
        {{else}}
        <p>No markdown files found in the data directory.</p>
        {{end}}
    </div>
</body>
</html>`

	t, err := template.New("filelist").Parse(tmpl)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, mdFiles); err != nil {
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
		return
	}
}
