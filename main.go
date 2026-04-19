package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"nikma/internal/config"
	"nikma/internal/handlers"
	"nikma/internal/middleware"
	"nikma/internal/repository"
	"nikma/internal/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize repository
	contentRepo := repository.NewContentRepository(cfg.ContentFile)

	// Initialize services
	contentService := services.NewContentService(contentRepo)
	authService := services.NewAuthService(cfg)
	uploadService := services.NewUploadService(cfg)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// Initialize handlers
	contentHandler := handlers.NewContentHandler(contentService)
	authHandler := handlers.NewAuthHandler(authService)
	uploadHandler := handlers.NewUploadHandler(uploadService, authMiddleware)
	pageHandler := handlers.NewPageHandler()

	// Parse templates
	indexTemplate, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("Failed to parse index template: %v", err)
	}

	dashboardTemplate, err := template.ParseFiles("dashboard.html")
	if err != nil {
		log.Fatalf("Failed to parse dashboard template: %v", err)
	}

	// Setup routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		content, err := contentService.GetContent()
		if err != nil {
			http.Error(w, "Failed to load content", http.StatusInternalServerError)
			return
		}
		err = indexTemplate.Execute(w, content)
		if err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		content, err := contentService.GetContent()
		if err != nil {
			http.Error(w, "Failed to load content", http.StatusInternalServerError)
			return
		}
		err = dashboardTemplate.Execute(w, content)
		if err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	})

	// API routes
	http.HandleFunc("/api/content", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			contentHandler.GetContent(w, r)
		case http.MethodPost:
			contentHandler.UpdateContent(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/authenticate", authHandler.Authenticate)
	http.HandleFunc("/api/upload-image", uploadHandler.UploadImage)

	// Static files
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fs = http.FileServer(http.Dir("./uploads"))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	port := cfg.Port
	fmt.Printf("Server running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
