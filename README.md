# Lema Take Home Assignment

Lema test backend service.
A simple web application built with the [Fiber framework](https://github.com/gofiber/fiber) and [SQLite](https://www.sqlite.org/), designed for performance, simplicity, and scalability.

## Table of Contents
- [Installation](#installation)
- [Configuration](#configuration)

## Installation

To get started with this application, follow the steps below:

1. Navigate into the project directory:
    ```bash
    cd lema
    ```

2. Install dependencies using Go modules:
    ```bash
    go mod tidy
    ```

3. If you don't have `SQLite` installed, you can follow their [installation guide](https://www.sqlite.org/download.html).

## Configuration

1. Edit the `.env` file in the root of the project. Add your database name and app port:


2. You can also configure additional settings like port and environment:

    ```plaintext
    PORT=3000
    DB_NAME=lema.sqlite
    ```


