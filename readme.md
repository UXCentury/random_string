# Random String Generator Web Server

This Go (Golang) program creates a simple web server that generates random strings based on specified query parameters. It is designed to provide flexibility in generating random strings with various options.

## Table of Contents
- [Usage](#usage)
- [Customization](#customization)
- [Getting Started](#getting-started)

## Usage

To use this web server, you can make HTTP requests to the root URL (`/`) and customize the random string generation using query parameters. Here are the available query parameters:

- `length`: The length of the random string.
- `alpha`: Whether to include alphabetic characters (true/false).
- `numeric`: Whether to include numeric characters (true/false).
- `special`: Whether to include special characters (true/false).
- `useUpper`: Whether to include uppercase letters (true/false).
- `useLower`: Whether to include lowercase letters (true/false).

### Example Request

To generate a random string with a length of 12, including both alphabetic and numeric characters, you can make a request like this:

http://localhost:8080/?length=12&alpha=true&numeric=true


The server will respond with a JSON object containing the generated random string.

## Customization

You can customize the behavior of the random string generator by modifying the query parameters in your requests. The server provides flexibility to tailor the generated strings to your specific requirements.

## Getting Started

Follow these steps to run the web server locally:

1. Make sure you have Go (Golang) installed on your machine.

2. Clone this repository to your local system.

3. Open a terminal and navigate to the repository directory.

4. Run the following command to start the web server:

   ```shell
   go run main.go

The server will start and listen on port 8080.

You can make HTTP requests to the server with the desired query parameters to generate random strings.

