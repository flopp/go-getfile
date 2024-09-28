# go-getfile
Go module to download files

````
client := getfile.NewClient()

// use a custom user agent
client.SetUserAgent("my-user-agent")

// Download file from url, overwrite existing file
url := "https://www.golang-book.com/public/pdf/gobook.pdf"
if err := client.Get(url, "gobook.pdf"); err != nil {
    fmt.Printf("Download failed: %v\n", err)
}

// Download file from url, but only if the target files doesn't exist already
if err := client.GetIfNotExists(url, "gobook.pdf"); err != nil {
    fmt.Printf("Download failed: %v\n", err)
}

// Downfile file from url, but only if the target files doesn't exist already or is older than the given time duration (24 hours)
if err := client.GetIfOutdated(url, "gobook.pdf", 24 * time.Hour); err != nil {
    fmt.Printf("Download failed: %v\n", err)
}
````
