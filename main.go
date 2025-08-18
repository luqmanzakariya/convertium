package main

import (
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Room struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Subtitle    string  `json:"subtitle"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl,omitempty"`
}

type Booking struct {
	ID               int       `json:"id"`
	UserID           int       `json:"userId"`
	RoomID           int       `json:"roomId"`
	CheckIn          time.Time `json:"checkIn"`
	CheckOut         time.Time `json:"checkOut"`
	Title            string    `json:"title"`
	ContactName      string    `json:"contactName"`
	ContactEmail     string    `json:"contactEmail"`
	NumberOfGuests   int       `json:"numberOfGuests"`
	TotalPrice       float64   `json:"totalPrice"`
	Tax              float64   `json:"tax"`
	ConfirmationCode string    `json:"confirmationCode"`
	Confirmed        bool      `json:"confirmed"`
	Cancelled        bool      `json:"cancelled"`
}

var (
	users = []User{}
	rooms = []Room{
		{1, "Standard Room", "Basic room with bed", `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.`, 100.0, "https://storage.googleapis.com/images.luqmanzakariya.com/others/basic-room-with-bed.webp"},
		{2, "Deluxe Room", "Room with bamboo view", `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.`, 150.0, "https://storage.googleapis.com/images.luqmanzakariya.com/others/room-with-bamboo.webp"},
		{3, "Suite", "Luxury suite", `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.`, 200.0, "https://storage.googleapis.com/images.luqmanzakariya.com/others/room-suite.webp"},
		{4, "Executive Suite", "Spacious suite with work area and premium amenities", `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.`, 250.0, "https://storage.googleapis.com/images.luqmanzakariya.com/others/room-executive.webp"},
		{5, "Family Room", "Large room with two queen beds, suitable for families", `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.`, 180.0, "https://storage.googleapis.com/images.luqmanzakariya.com/others/room-family.webp"},
		{6, "Ocean View Room", "Room with balcony overlooking the ocean", `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.`, 220.0, "https://storage.googleapis.com/images.luqmanzakariya.com/others/room-ocean-view.webp"},
	}
	bookings      = []Booking{}
	nextUserID    = 1
	nextBookingID = 1
	mutex         sync.Mutex
	jwtSecret     = []byte("secret-key")
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.POST("/register", register)
	e.POST("/login", login)

	e.GET("/rooms", searchRooms)
	e.GET("/rooms/:id", getRoomByID)

	// Restricted group
	r := e.Group("/bookings")
	r.Use(authMiddleware)

	r.POST("", createBooking)
	r.GET("", listBookings)
	r.GET("/:id", getBookingByID)
	r.DELETE("/:id", cancelBooking)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}

	e.Logger.Fatal(e.Start(":" + port))
}

func register(c echo.Context) error {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	mutex.Lock()
	defer mutex.Unlock()
	for _, u := range users {
		if u.Email == req.Email {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "User exists"})
		}
	}
	user := User{ID: nextUserID, Email: req.Email, Password: req.Password}
	users = append(users, user)
	nextUserID++
	return c.JSON(http.StatusOK, echo.Map{"message": "Registered"})
}

func login(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	mutex.Lock()
	defer mutex.Unlock()
	for _, u := range users {
		if u.Email == req.Email && u.Password == req.Password {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
				Subject:   strconv.Itoa(u.ID),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			})
			tokenStr, _ := token.SignedString(jwtSecret)
			return c.JSON(http.StatusOK, echo.Map{"token": tokenStr})
		}
	}
	return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
}

func generateConfirmationCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 10
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func searchRooms(c echo.Context) error {
	checkInStr := c.QueryParam("checkIn")
	checkOutStr := c.QueryParam("checkOut")
	checkIn, _ := time.Parse("2006-01-02", checkInStr)
	checkOut, _ := time.Parse("2006-01-02", checkOutStr)

	available := []Room{}
	mutex.Lock()
	for _, room := range rooms {
		conflict := false
		for _, b := range bookings {
			if b.RoomID == room.ID && !b.Cancelled && b.Confirmed &&
				!(checkOut.Before(b.CheckIn) || checkIn.After(b.CheckOut)) {
				conflict = true
				break
			}
		}
		if !conflict {
			available = append(available, room)
		}
	}
	mutex.Unlock()
	return c.JSON(http.StatusOK, available)
}

func getRoomByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid room ID"})
	}

	mutex.Lock()
	defer mutex.Unlock()
	for _, room := range rooms {
		if room.ID == id {
			return c.JSON(http.StatusOK, room)
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "Room not found"})
}

func getBookingByID(c echo.Context) error {
	userID := c.Get("userID").(int)
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid booking ID"})
	}

	mutex.Lock()
	defer mutex.Unlock()
	for _, booking := range bookings {
		if booking.ID == id && booking.UserID == userID {
			return c.JSON(http.StatusOK, booking)
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "Booking not found or unauthorized"})
}

func createBooking(c echo.Context) error {
	userID := c.Get("userID").(int)
	var req struct {
		RoomID         int    `json:"roomId"`
		CheckIn        string `json:"checkIn"`
		CheckOut       string `json:"checkOut"`
		Title          string `json:"title"`
		ContactName    string `json:"contactName"`
		ContactEmail   string `json:"contactEmail"`
		NumberOfGuests int    `json:"numberOfGuests"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// Validate title
	if req.Title != "Mr" && req.Title != "Ms" && req.Title != "Mrs" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid title. Must be Mr, Ms, or Mrs"})
	}

	// Validate number of guests
	if req.NumberOfGuests < 1 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Number of guests must be at least 1"})
	}

	checkIn, err := time.Parse("2006-01-02", req.CheckIn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid check-in date"})
	}
	checkOut, err := time.Parse("2006-01-02", req.CheckOut)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid check-out date"})
	}

	mutex.Lock()
	defer mutex.Unlock()

	// Find room to get price
	var roomPrice float64
	for _, room := range rooms {
		if room.ID == req.RoomID {
			roomPrice = room.Price
			break
		}
	}
	if roomPrice == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Room not found"})
	}

	// Check for booking conflicts
	for _, b := range bookings {
		if b.RoomID == req.RoomID && !b.Cancelled && b.Confirmed &&
			!(checkOut.Before(b.CheckIn) || checkIn.After(b.CheckOut)) {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Room unavailable"})
		}
	}

	// Calculate total price: (room price * nights) + 9% tax
	nights := checkOut.Sub(checkIn).Hours() / 24
	if nights < 1 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Check-out must be after check-in"})
	}
	basePrice := roomPrice * nights
	tax := basePrice * 0.09
	totalPrice := basePrice + tax

	booking := Booking{
		ID:               nextBookingID,
		UserID:           userID,
		RoomID:           req.RoomID,
		CheckIn:          checkIn,
		CheckOut:         checkOut,
		Title:            req.Title,
		ContactName:      req.ContactName,
		ContactEmail:     req.ContactEmail,
		NumberOfGuests:   req.NumberOfGuests,
		TotalPrice:       totalPrice,
		Tax:              tax,
		ConfirmationCode: generateConfirmationCode(),
		Confirmed:        true,
	}
	bookings = append(bookings, booking)
	nextBookingID++
	return c.JSON(http.StatusOK, booking)
}

func listBookings(c echo.Context) error {
	userID := c.Get("userID").(int)
	now := time.Now()
	past := []Booking{}
	upcoming := []Booking{}
	mutex.Lock()
	for _, b := range bookings {
		if b.UserID == userID && !b.Cancelled {
			if b.CheckOut.Before(now) {
				past = append(past, b)
			} else {
				upcoming = append(upcoming, b)
			}
		}
	}
	mutex.Unlock()
	return c.JSON(http.StatusOK, echo.Map{"past": past, "upcoming": upcoming})
}

func cancelBooking(c echo.Context) error {
	userID := c.Get("userID").(int)
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	mutex.Lock()
	defer mutex.Unlock()
	for i, b := range bookings {
		if b.ID == id && b.UserID == userID && !b.Cancelled && b.CheckIn.After(time.Now()) {
			bookings[i].Cancelled = true
			return c.JSON(http.StatusOK, echo.Map{"message": "Cancelled"})
		}
	}
	return c.JSON(http.StatusBadRequest, echo.Map{"error": "Cannot cancel"})
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := c.Request().Header.Get("Authorization")
		if tokenStr == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Unauthorized"})
		}
		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
		}
		userID, _ := strconv.Atoi(claims.Subject)
		c.Set("userID", userID)
		return next(c)
	}
}
