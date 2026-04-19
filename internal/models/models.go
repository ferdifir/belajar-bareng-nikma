package models

// ContentData represents the structure of our website content
type ContentData struct {
	Hero         HeroSection         `json:"hero"`
	About        AboutSection        `json:"about"`
	Program      ProgramSection      `json:"program"`
	Gallery      GallerySection      `json:"gallery"`
	Testimonials TestimonialsSection `json:"testimonials"`
	Contact      ContactSection      `json:"contact"`
	Footer       FooterSection       `json:"footer"`
}

// HeroSection represents the hero section content
type HeroSection struct {
	Title           string `json:"title"`
	Subtitle        string `json:"subtitle"`
	Description     string `json:"description"`
	WhatsappNumber  string `json:"whatsappNumber"`
	WhatsappMessage string `json:"whatsappMessage"`
}

// AboutSection represents the about section content
type AboutSection struct {
	Title        string `json:"title"`
	Description1 string `json:"description1"`
	Description2 string `json:"description2"`
	Description3 string `json:"description3"`
}

// ProgramSection represents the program section content
type ProgramSection struct {
	Title string      `json:"title"`
	Sd    ProgramItem `json:"sd"`
	Smp   ProgramItem `json:"smp"`
}

// ProgramItem represents a single program (SD or SMP)
type ProgramItem struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Features    []string `json:"features"`
}

// GallerySection represents the gallery section content
type GallerySection struct {
	Title string        `json:"title"`
	Items []GalleryItem `json:"items"`
}

// GalleryItem represents a single gallery item
type GalleryItem struct {
	Title string `json:"title"`
	Image string `json:"image"` // Path to the image file
}

// TestimonialsSection represents the testimonials section content
type TestimonialsSection struct {
	Title string            `json:"title"`
	Items []TestimonialItem `json:"items"`
}

// TestimonialItem represents a single testimonial
type TestimonialItem struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

// ContactSection represents the contact section content
type ContactSection struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ServiceArea string `json:"serviceArea"`
	ButtonText  string `json:"buttonText"`
}

// FooterSection represents the footer content
type FooterSection struct {
	Text string `json:"text"`
}

// AuthCredentials holds the username and password for authentication
type AuthCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UploadResponse represents the response for image upload
type UploadResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	ImagePath string `json:"imagePath"`
}

// APIResponse represents a generic API response
type APIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// SuccessResponse creates a success API response
func SuccessResponse(message string) APIResponse {
	return APIResponse{
		Status:  "success",
		Message: message,
	}
}

// ErrorResponse creates an error API response
func ErrorResponse(message string) APIResponse {
	return APIResponse{
		Status:  "error",
		Message: message,
	}
}
