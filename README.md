# go-getfile
Go module to download files

```go
// Create a getfile client
client := getfile.NewClient()

// Use a custom user agent
client.SetUserAgent("my-user-agent")

// Set a delay, e.g. duration to wait before each download
client.SetDelay(1 * time.Second)

// The free eBook "An Introduction to Programming in Go"
url := "https://www.golang-book.com/public/pdf/gobook.pdf"
targetFile := "gobook.pdf"

// Download file from url,
// overwrite existing target file
if err := client.Get(url, targetFile); err != nil {
    fmt.Printf("Download failed: %v\n", err)
}

// Download file from url,
// but only if the target file doesn't exist already
if err := client.GetIfNotExists(url, targetFile); err != nil {
    fmt.Printf("Download failed: %v\n", err)
}

// Downfile file from url,
// but only if the target files doesn't exist already or
// it is older than the given time duration (24 hours)
if err := client.GetIfOutdated(url, targetFile, 24 * time.Hour); err != nil {
    fmt.Printf("Download failed: %v\n", err)
}
```
