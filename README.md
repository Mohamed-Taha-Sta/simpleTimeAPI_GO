# simple Time API using GoLang

# Timezone API

- This is a simple HTTP server written in Go that provides the current time for various locations around the world.
- I created this API primarily to help me understand the fundamentals of Go and to assist me with future projects. 
Utilizing web server APIs can sometimes lead to inconsistent availability and, on occasion, incur significant costs.

## Features

- Provides the current time for various locations around the world. Feel Free to add any timezones if needed, simply by
appending to the map containing the time differences from UTC for each location.
- Returns the current time for a given location in a JSON response, in a counts for most use cases.

## Usage

Make a request to the "/time" endpoint with a location as a query parameter. The server will calculate the current time for that location and return it in a JSON response.

Example of request:

![ExampleUsingPostMan](misc/examplePostman.png)

Setup
1. Clone the repository.
2. Run go build to compile the application.
3. Run the application with ./APIProject.

Any contributions are more than welcome!