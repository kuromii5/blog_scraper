# Blog Scraper REST API

This project is a REST API built with GoLang that serves as a blog scraper. It collects posts from XML documents, stores them in a PostgreSQL database, and provides endpoints to interact with the data.

## Features

- **PostgreSQL Database:** Utilizes PostgreSQL as the database to store scraped blog posts.
- **SQLC:** Uses sqlc to generate type-safe Go code for interacting with the PostgreSQL database.
- **REST API:** Provides RESTful endpoints for creating users, following feeds, and retrieving scraped blog posts.
- **User Management:** Users can register, login, and manage their followed feeds.
- **Scraping:** Automatically scrapes posts from websites.

## Tech Stack

- **GoLang:** Main programming language for the project.
- **PostgreSQL:** Database management system.
- **sqlc:** Go code generation tool for SQL.
- **Chi:** Lightweight, idiomatic router for Go.
- **XML Parsing:** Used to extract blog post data from XML documents.
