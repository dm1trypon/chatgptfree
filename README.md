# go-chatgptfree-api

The `go-chatgptfree-api` package is a simple Golang library for interacting with a ChatGPT-like API to generate text completions based on user prompts. The package provides a function to send input to the API and retrieve a response, making it easy to integrate AI-driven conversational capabilities into your applications.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Functions](#functions)
- [Error Handling](#error-handling)
- [Contributions](#contributions)
- [License](#license)

## Installation

To install the `go-chatgptfree-api` package, you can use Go's package manager. Run the following command in your terminal:

```bash
go get github.com/dm1trypon/go-chatgptfree-api
```

Replace `github.com/dm1trypon/go-chatgptfree-api` with the appropriate module path if necessary.

## Usage

Here is a basic example of how to use the `GenerateText` function to get a text completion from the ChatGPT API:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dm1trypon/go-chatgptfree-api"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	prompt := "What are the benefits of using Golang?"
	response, err := chatgptfree.GenerateText(ctx, prompt)
	if err != nil {
		log.Fatalf("Error while generating text: %v", err)
	}

	fmt.Println("Response from API:", string(response))
}
```

In this example, we're creating a context with a timeout of 5 seconds to make sure the request does not hang indefinitely. We then call the `GenerateText` function with our prompt and print out the response.

## Functions

### `GenerateText(ctx context.Context, prompt string) ([]byte, error)`

- **Description**: Generates text based on the provided prompt by sending a request to the ChatGPT API.
- **Parameters**:
    - `ctx`: A context for managing request timeouts.
    - `prompt`: The text prompt that you want the AI to respond to.
- **Returns**:
    - `[]byte`: The generated text response from the AI.
    - `error`: An error if something went wrong during the request.

## Error Handling

The package defines a few error types that you can use to handle specific error scenarios:

- `errResponseCodeIsNot200`: Returned when the HTTP response code is not 200 (OK).
- `errEmptyRespChoices`: Returned when the response from the API does not contain any choices.

You can check for these errors in your code when calling `GenerateText`.
