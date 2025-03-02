OpenDataUG Go Client

A Go client library for accessing Uganda's administrative divisions data through the Open Data API.

## Installation

```go
go get github.com/yourusername/opendataug
```

## Usage

### Initialize the client

go
import "github.com/yourusername/opendataug"
func main() {
// Create a new client
client := opendataug.NewClient()
// Optionally set a custom base URL if needed
// client.BaseURL = "https://custom-api-url.com/api"
// Use the client to access various endpoints
// ...
}

### Working with Villages

// Get all villages
villages, err := client.GetVillages()
if err != nil {
log.Fatalf("Error getting villages: %v", err)
}
for , village := range villages {
fmt.Printf("Village: %s (ID: %s)\n", village.Name, village.ID)
}
// Get a specific village by ID
village, err := client.GetVillage("some-village-id")
if err != nil {
log.Fatalf("Error getting village: %v", err)
}
fmt.Printf("Village details: %+v\n", village)
// Get all villages in a specific parish
parishVillages, err := client.GetVillagesByParish("some-parish-id")
if err != nil {
log.Fatalf("Error getting villages by parish: %v", err)
}
fmt.Printf("Found %d villages in the parish\n", len(parishVillages))

### Working with other administrative divisions

The library also provides methods to work with other administrative divisions:

- Districts
- Counties/Subcounties
- Parishes

Check the documentation for more details on these methods.

## Data Models

The library provides Go structs that map to the API's JSON responses, including:

- Village
- Parish
- Subcounty
- District

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)
