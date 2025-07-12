# RSS Aggregator

This is a backend service for an RSS aggregator, written in Go. It allows users to register, subscribe to RSS feeds, and fetch the latest posts from those feeds.

## Features

*   **User Authentication:** Users can create an account and receive an API key for authentication.
*   **Feed Subscriptions:** Authenticated users can subscribe to RSS feeds by providing a URL.
*   **Feed Following:** Users can follow and unfollow feeds to customize their post feed.
*   **Post Fetching:** The service periodically fetches new posts from the subscribed RSS feeds.
*   **Personalized Post Feed:** Users can retrieve a list of the latest posts from all the feeds they follow.

## Technologies Used

*   **Go:** The primary programming language for the backend service.
*   **PostgreSQL:** The database used to store user data, feeds, and posts.
*   **Chi:** A lightweight, idiomatic and composable router for building Go HTTP services.
*   **godotenv:** A library to load environment variables from a `.env` file.
*   **pq:** A pure Go Postgres driver for the `database/sql` package.
*   **uuid:** A library for generating and working with UUIDs.

## API Endpoints

The following API endpoints are available under the `/v1` path:

*   `GET /ready`: Health check endpoint to verify if the service is running.
*   `GET /error`: An endpoint to test error handling.
*   `POST /users`: Create a new user.
*   `GET /users`: Get the current authenticated user's information.
*   `POST /feeds`: Create a new feed subscription.
*   `GET /feeds`: Get all feeds in the system.
*   `POST /feed_follows`: Follow a specific feed.
*   `GET /feed_follows`: Get all the feeds followed by the authenticated user.
*   `DELETE /feed_follows/{feedFollowID}`: Unfollow a specific feed.
*   `GET /posts`: Get the latest posts from all the feeds followed by the authenticated user.

## Getting Started

To get started with this project locally, you will need to have Go and PostgreSQL installed.

1.  **Clone the repository:**
    ```bash
    git clone <repository-url>
    cd rssagg
    ```

2.  **Create a `.env` file:**
    Create a `.env` file in the root of the project with the following content, replacing the placeholder values with your actual database credentials:
    ```
    PORT=8080
    DB_URL=postgres://user:password@localhost:5432/rssagg?sslmode=disable
    ```

3.  **Set up the database:**
    The project uses `sqlc` to generate type-safe Go code from SQL queries. You will need to have the database schema set up. The SQL files for creating the schema are located in the `sql/schema` directory.

4.  **Run the application:**
    ```bash
    go run main.go
    ```

The server will start on the port specified in your `.env` file (e.g., 8080).

