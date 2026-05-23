# Web Crawler - Go version

A CLI web crawler written in Go that generates an internal links report for any website. Built as part of the Boot.dev curriculum.

## Usage

```bash
go run . <url>
```

## Features

- Crawls all pages of a given website
- Reports internal link counts per page
- Respects the same domain (no external links followed)  
