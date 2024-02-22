# README for Your CLI Application

## Overview

This CLI application is a powerful and flexible tool designed to fetch and display jokes from various categories, including programming and miscellaneous, as well as philosophical quotes. Built with Go (Golang) and leveraging the Cobra library, it showcases the ease of creating structured and user-friendly command-line interfaces.

## Features

- **Fetch Jokes:** Retrieve jokes from specified categories, making the CLI a fun tool to lighten up your day.
- **Philosophical Quotes:** Gain insights and inspiration from philosophical quotes, fetched directly into your terminal.
- **Extensibility:** Designed with scalability in mind, allowing for easy addition of new commands and functionalities.

## Getting Started

### Prerequisites

- Go (Golang) installed on your machine (version 1.15+ recommended).
- Basic understanding of CLI operations and Go programming.

### Installation

1. **Build the Application**

   Compile the application using Go:

   ```sh
   go build -o chatcli
   ```

   This command builds the application and creates an executable named `chatcli`.

2. **Run the Application**

   Execute the CLI directly to explore its functionalities:

   ```sh
   ./chatcli help
   ```

### Usage

The CLI provides several commands and options for fetching jokes and philosophical quotes:

- **Fetch a Joke:**

  ```sh
  ./chatcli joke --category Programming
  ```

  Replace `Programming` with `Miscellaneous` to fetch jokes from the miscellaneous category.

- **Fetch a Philosophical Quote:**

  ```sh
  ./chatcli philosophy
  ```

- **Help:**

  For more information on available commands and their usage, run:

  ```sh
  ./chatcli
  ```
