# Hotel Booking System

This repository contains a hotel booking system with a backend API built in Go using the Echo framework and a frontend built with Vue 3 and Vite. The system supports user registration, authentication, room searching, and booking management. You can view the deployed app version on [https://convertium.luqmanzakariya.com/](https://convertium.luqmanzakariya.com/).

## Repository Structure
- **/front**: Frontend code built with Vue 3 and Vite.
- **go.mod**, **go.sum**, **main.go**: Backend Go application and dependencies.

## Backend Setup
- **Requirements**: Go 1.x
- **Run Locally**:
  ```bash
  go run main.go
  ```
- **Port**: Defaults to `8080` unless overridden by the `PORT` environment variable.

## Frontend Setup
- **Requirements**: Node.js (version 20.19+ or 22.12+), npm
  - Note: Some templates may require a higher Node.js version. Upgrade if your package manager warns about it.
- **Run Locally**:
  ```bash
  cd front
  npm install
  npm run dev
  ```
- **Deployed Version**: [https://convertium.luqmanzakariya.com/](https://convertium.luqmanzakariya.com/)

## API Documentation

**Base URL**: `http://localhost:8080` (or specified `PORT`)

**Authentication**: Most endpoints (`/bookings`) require a JWT token in the `Authorization` header: `Authorization: <JWT_TOKEN>`.

### Endpoints

1. **Register User**
   - **Route**: `/register`
   - **Method**: `POST`
   - **Headers**: `Content-Type: application/json`
   - **Body**: `{ "name": "string", "email": "string", "password": "string" }`
   - **Description**: Registers a new user.
   - **Response**:
     - **200**: `{"message": "Registered"}`
     - **400**: `{"error": "User exists"}`

2. **Login**
   - **Route**: `/login`
   - **Method**: `POST`
   - **Headers**: `Content-Type: application/json`
   - **Body**: `{ "email": "string", "password": "string" }`
   - **Description**: Authenticates user and returns JWT token (valid 24 hours).
   - **Response**:
     - **200**: `{"token": "<JWT_TOKEN>"}`
     - **401**: `{"error": "Invalid credentials"}`

3. **Search Rooms**
   - **Route**: `/rooms`
   - **Method**: `GET`
   - **Query Params**: `checkIn=YYYY-MM-DD`, `checkOut=YYYY-MM-DD` (optional)
   - **Description**: Lists available rooms, filtered by date availability if provided.
   - **Response**:
     - **200**: Array of rooms (`id`, `name`, `subtitle`, `description`, `price`, `imageUrl`)

4. **Get Room by ID**
   - **Route**: `/rooms/:id`
   - **Method**: `GET`
   - **Description**: Retrieves room details by ID.
   - **Response**:
     - **200**: Room object
     - **404**: `{"error": "Room not found"}`

5. **Create Booking**
   - **Route**: `/bookings`
   - **Method**: `POST`
   - **Headers**: `Content-Type: application/json`, `Authorization: <JWT_TOKEN>`
   - **Body**: `{ "roomId": 1, "checkIn": "YYYY-MM-DD", "checkOut": "YYYY-MM-DD", "title": "Mr|Ms|Mrs", "contactName": "string", "contactEmail": "string", "numberOfGuests": 1 }`
   - **Description**: Creates a booking with 9% tax. Validates title and availability.
   - **Response**:
     - **200**: Booking object (`id`, `userId`, `roomId`, `checkIn`, `checkOut`, etc.)
     - **400**: `{"error": "<error message>"}`
     - **401**: `{"error": "Unauthorized"}`

6. **List Bookings**
   - **Route**: `/bookings`
   - **Method**: `GET`
   - **Headers**: `Authorization: <JWT_TOKEN>`
   - **Description**: Lists user's non-cancelled bookings, split into past and upcoming.
   - **Response**:
     - **200**: `{ "past": [bookings], "upcoming": [bookings] }`
     - **401**: `{"error": "Unauthorized"}`

7. **Get Booking by ID**
   - **Route**: `/bookings/:id`
   - **Method**: `GET`
   - **Headers**: `Authorization: <JWT_TOKEN>`
   - **Description**: Retrieves a specific booking for the authenticated user.
   - **Response**:
     - **200**: Booking object
     - **404**: `{"error": "Booking not found or unauthorized"}`

8. **Cancel Booking**
   - **Route**: `/bookings/:id`
   - **Method**: `DELETE`
   - **Headers**: `Authorization: <JWT_TOKEN>`
   - **Description**: Cancels a booking if check-in is in the future.
   - **Response**:
     - **200**: `{"message": "Cancelled"}`
     - **400**: `{"error": "Cannot cancel"}`
     - **401**: `{"error": "Unauthorized"}`

## Notes
- **Dates**: Input as `YYYY-MM-DD`, output in ISO 8601 (`YYYY-MM-DDThh:mm:ssZ`).
- **Storage**: In-memory (non-persistent).
- **JWT Secret**: Hardcoded as `secret-key` (use environment variable in production).
- **CORS**: Enabled for cross-origin requests.