package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"nikma/internal/middleware"
	"nikma/internal/models"
	"nikma/internal/services"
)

// ContentHandler handles content-related HTTP requests
type ContentHandler struct {
	service *services.ContentService
}

// NewContentHandler creates a new content handler
func NewContentHandler(service *services.ContentService) *ContentHandler {
	return &ContentHandler{service: service}
}

// GetContent handles GET /api/content
func (h *ContentHandler) GetContent(w http.ResponseWriter, r *http.Request) {
	content, err := h.service.GetContent()
	if err != nil {
		http.Error(w, "Failed to load content", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(content)
}

// UpdateContent handles POST /api/content
func (h *ContentHandler) UpdateContent(w http.ResponseWriter, r *http.Request) {
	var content models.ContentData
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateContent(&content); err != nil {
		http.Error(w, "Failed to save content", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.SuccessResponse("Content updated successfully"))
}

// AuthHandler handles authentication requests
type AuthHandler struct {
	service *services.AuthService
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Authenticate handles POST /api/authenticate
func (h *AuthHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	var creds models.AuthCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if h.service.ValidateCredentials(creds.Username, creds.Password) {
		json.NewEncoder(w).Encode(map[string]bool{"success": true})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]bool{"success": false})
	}
}

// UploadHandler handles file upload requests
type UploadHandler struct {
	uploadService *services.UploadService
	authMiddleware *middleware.AuthMiddleware
}

// NewUploadHandler creates a new upload handler
func NewUploadHandler(uploadService *services.UploadService, authMiddleware *middleware.AuthMiddleware) *UploadHandler {
	return &UploadHandler{
		uploadService: uploadService,
		authMiddleware: authMiddleware,
	}
}

// UploadImage handles POST /api/upload-image
func (h *UploadHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form with max size
	r.ParseMultipartForm(h.uploadService.GetMaxUploadSize())

	// Get the file from the request
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create uploads directory if it doesn't exist
	uploadsDir := h.uploadService.GetUploadsDir()
	if err := os.MkdirAll(uploadsDir, os.ModePerm); err != nil {
		http.Error(w, "Error creating uploads directory", http.StatusInternalServerError)
		return
	}

	// Create destination file
	destPath := filepath.Join(uploadsDir, handler.Filename)
	dst, err := os.Create(destPath)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy uploaded file to destination
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// Return success response with image path
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.UploadResponse{
		Status:    "success",
		Message:   "Image uploaded successfully",
		ImagePath: "/uploads/" + handler.Filename,
	})
}

// PageHandler handles page serving requests
type PageHandler struct{}

// NewPageHandler creates a new page handler
func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

// GetIndex serves the index.html file
func (h *PageHandler) GetIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// GetDashboard serves the dashboard.html file
func (h *PageHandler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dashboard.html")
}
