package services

import (
	"nikma/internal/config"
	"nikma/internal/models"
	"nikma/internal/repository"
)

// ContentService handles business logic for content
type ContentService struct {
	repo *repository.ContentRepository
}

// NewContentService creates a new content service
func NewContentService(repo *repository.ContentRepository) *ContentService {
	return &ContentService{repo: repo}
}

// GetContent retrieves the current content
func (s *ContentService) GetContent() (*models.ContentData, error) {
	return s.repo.Load()
}

// UpdateContent updates the content
func (s *ContentService) UpdateContent(content *models.ContentData) error {
	return s.repo.Save(content)
}

// AuthService handles authentication logic
type AuthService struct {
	config *config.Config
}

// NewAuthService creates a new auth service
func NewAuthService(cfg *config.Config) *AuthService {
	return &AuthService{config: cfg}
}

// ValidateCredentials checks if the provided credentials are valid
func (s *AuthService) ValidateCredentials(username, password string) bool {
	return username == s.config.AuthUsername && password == s.config.AuthPassword
}

// UploadService handles file upload logic
type UploadService struct {
	config *config.Config
}

// NewUploadService creates a new upload service
func NewUploadService(cfg *config.Config) *UploadService {
	return &UploadService{config: cfg}
}

// GetUploadsDir returns the uploads directory path
func (s *UploadService) GetUploadsDir() string {
	return s.config.UploadsDir
}

// GetMaxUploadSize returns the maximum upload size
func (s *UploadService) GetMaxUploadSize() int64 {
	return s.config.MaxUploadSize
}
