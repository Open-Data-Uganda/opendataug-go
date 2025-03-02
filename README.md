# OpenDataUG Go Client

A Go client library for accessing Uganda's administrative divisions data through the Open Data API.

## Installation

```bash
go get github.com/Open-Data-Uganda/opendataug-go
```

## Usage

### Initialize the Client

```go
package main

import (
    "fmt"
    "log"
    "github.com/Open-Data-Uganda/opendataug-go"
)

func main() {
    client := opendataug.NewClient()
}
```

### Working with Villages

```go
villages, err := client.GetVillages()
if err != nil {
    log.Fatalf("Error getting villages: %v", err)
}

for _, village := range villages {
    fmt.Printf("Village: %s (ID: %s)\n", village.Name, village.ID)
}

village, err := client.GetVillage("village-123")
if err != nil {
    log.Fatalf("Error getting village: %v", err)
}
fmt.Printf("Village details: %+v\n", village)

parishVillages, err := client.GetVillagesByParish("parish-456")
if err != nil {
    log.Fatalf("Error getting villages by parish: %v", err)
}
fmt.Printf("Found %d villages in the parish\n", len(parishVillages))
```

### Administrative Divisions

The library provides comprehensive access to Uganda's administrative divisions:

#### Districts

```go
districts, err := client.GetDistricts()
district, err := client.GetDistrict("district-789")
```

#### Counties/Subcounties

```go
subcounties, err := client.GetSubcountiesByDistrict("district-789")
subcounty, err := client.GetSubcounty("subcounty-012")
```

#### Parishes

```go
parishes, err := client.GetParishesBySubcounty("subcounty-012")
parish, err := client.GetParish("parish-456")
```

## Data Models

The library provides the following data models that map to the API's JSON responses:

```go
type District struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

type Subcounty struct {
    ID         string `json:"id"`
    Name       string `json:"name"`
    DistrictID string `json:"district_id"`
}

type Parish struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    SubcountyID string `json:"subcounty_id"`
}

type Village struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    ParishID  string `json:"parish_id"`
}
```

## Error Handling

The library uses standard Go error handling patterns:

```go
villages, err := client.GetVillages()
if err != nil {
    switch e := err.(type) {
    case *opendataug.APIError:
        log.Printf("API Error: %v (Status: %d)", e.Message, e.StatusCode)
    default:
        log.Printf("Error: %v", err)
    }
}
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the [MIT License](LICENSE) - see the LICENSE file for details.

## Support

If you encounter any issues or have questions, please file an issue on the GitHub repository.
