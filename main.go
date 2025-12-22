package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

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
	Image string `json:"image"`  // Path to the image file
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

// Global variable to hold our content data
var contentData ContentData

// File path for storing content data
const contentFilePath = "content.json"

// LoadContent loads content from the JSON file
func LoadContent() error {
	// Check if the file exists
	if _, err := os.Stat(contentFilePath); os.IsNotExist(err) {
		// If file doesn't exist, initialize with default content
		InitializeDefaultContent()
		return SaveContent()
	}

	// Read the file
	data, err := ioutil.ReadFile(contentFilePath)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data
	return json.Unmarshal(data, &contentData)
}

// SaveContent saves content to the JSON file
func SaveContent() error {
	data, err := json.MarshalIndent(contentData, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(contentFilePath, data, 0644)
}

// InitializeDefaultContent initializes the content with default values
func InitializeDefaultContent() {
	contentData = ContentData{
		Hero: HeroSection{
			Title:           "Bimbel Tuntas, Nilai Pintas.",
			Subtitle:        "Bimbingan Matematika & IPA oleh Sarjana Fisika.",
			Description:     "Bimbingan intensif dan personal untuk siswa SD dan SMP. Ubah kesulitan belajar menjadi prestasi nyata, fokus pada pemahaman konsep dasar.",
			WhatsappNumber:  "6281234567890",
			WhatsappMessage: "Halo Kak Nikma, saya tertarik dengan bimbel Belajar Bareng Nikma. Saya ingin mendaftar dan mendapatkan sesi perkenalan gratis.",
		},
		About: AboutSection{
			Title:        "Kenalan dengan Kak Nikma",
			Description1: "Kak Nikma adalah lulusan Sarjana Fisika yang memiliki <span class='font-semibold text-gold'>passion mendalam dalam mengajar Matematika dan IPA (Fisika/Biologi)</span>.",
			Description2: "Dengan metode pengajaran yang sabar, terstruktur, dan fokus pada pemecahan masalah, Kak Nikma membantu membangun kepercayaan diri siswa SD dan SMP dalam belajar.",
			Description3: "Tujuan utamanya adalah mengubah kesulitan belajar menjadi prestasi nyata melalui pemahaman konsep dasar yang kuat.",
		},
		Program: ProgramSection{
			Title: "Program Belajar Bareng Nikma",
			Sd: ProgramItem{
				Title:       "Program SD (Kelas 4-6)",
				Description: "Fokus pada Dasar-dasar Matematika dan Sains.",
				Features: []string{
					"Penguatan Calitung (Catur, Literasi, Hitung)",
					"Pemahaman konsep dasar Matematika",
					"Latihan soal rutin",
				},
			},
			Smp: ProgramItem{
				Title:       "Program SMP (Kelas 7-9)",
				Description: "Fokus pada Matematika dan Fisika.",
				Features: []string{
					"Pemecahan Masalah Aljabar",
					"Konsep Dasar Fisika",
					"Persiapan Ujian Sekolah/Daerah",
				},
			},
		},
		Gallery: GallerySection{
			Title: "Galeri Kegiatan Belajar Bareng Nikma",
			Items: []GalleryItem{
				{Title: "Aktivitas Belajar Interaktif", Image: "/assets/images/gallery1.jpg"},
				{Title: "Murid-Murid Bahagia", Image: "/assets/images/gallery2.jpg"},
				{Title: "Sesi Belajar Kelompok", Image: "/assets/images/gallery3.jpg"},
				{Title: "Penghargaan Prestasi", Image: "/assets/images/gallery4.jpg"},
				{Title: "Kegiatan Praktikum IPA", Image: "/assets/images/gallery5.jpg"},
				{Title: "Sesi Evaluasi Mingguan", Image: "/assets/images/gallery6.jpg"},
			},
		},
		Testimonials: TestimonialsSection{
			Title: "Apa Kata Mereka?",
			Items: []TestimonialItem{
				{
					Text:   "Nilai matematika adik saya naik drastis setelah les di sini. Kak Nikma sabar dan bisa menjelaskan pelajaran dengan cara yang mudah dimengerti.",
					Author: "Ortu Murid SD",
				},
				{
					Text:   "Alhamdulillah, anak saya jadi lebih percaya diri saat ujian karena paham konsepnya. Terima kasih Kak Nikma!",
					Author: "Ortu Murid SMP",
				},
				{
					Text:   "Metode belajarnya seru dan ga bikin bosan. Anak saya malah semangat belajar IPA setiap minggu.",
					Author: "Ortu Murid SD",
				},
			},
		},
		Contact: ContactSection{
			Title:       "Hubungi Kami",
			Description: "Siap memulai perjalanan belajar yang menyenangkan dan efektif? Hubungi kami melalui WhatsApp!",
			ServiceArea: "Online dan Tatap Muka Area Bunder, Banyuwangi",
			ButtonText:  "Hubungi via WhatsApp",
		},
		Footer: FooterSection{
			Text: "&copy; 2025 Belajar Bareng Nikma. Hak Cipta Dilindungi.",
		},
	}
}

// GetContent returns the current content
func GetContent(w http.ResponseWriter, r *http.Request) {
	// No authentication required for getting content (public endpoint)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contentData)
}

// UpdateContent updates the content
func UpdateContent(w http.ResponseWriter, r *http.Request) {
	// Check if the request is authorized
	if !checkAuth(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	
	// Decode the request body into contentData
	if err := json.NewDecoder(r.Body).Decode(&contentData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Save the updated content to file
	if err := SaveContent(); err != nil {
		http.Error(w, "Failed to save content", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Content updated successfully"})
}

// AuthCredentials holds the username and password for authentication
type AuthCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Authenticate handles the authentication request
func Authenticate(w http.ResponseWriter, r *http.Request) {
	// Handle CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	var creds AuthCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check credentials (hardcoded server-side)
	if creds.Username == "nikma" && creds.Password == "250200" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{"success": true})
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]bool{"success": false})
	}
}

// checkAuth checks if the request contains valid authentication
func checkAuth(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false
	}

	// Check if it's a basic auth header
	if strings.HasPrefix(authHeader, "Basic ") {
		encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
		decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			return false
		}

		credentials := string(decodedCredentials)
		parts := strings.Split(credentials, ":")
		if len(parts) != 2 {
			return false
		}

		username := parts[0]
		password := parts[1]

		// Validate credentials (hardcoded server-side)
		return username == "nikma" && password == "280200"
	}

	return false
}

// GetIndex serves the index.html file
func GetIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// GetDashboard serves the dashboard.html file
func GetDashboard(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dashboard.html")
}

func main() {
	// Load content from file
	if err := LoadContent(); err != nil {
		log.Fatal("Failed to load content:", err)
	}

	// Create a new ServeMux
	r := http.NewServeMux()

	// API routes
	r.HandleFunc("/api/content", func(w http.ResponseWriter, r *http.Request) {
		// Handle CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		// Check authentication for POST requests (updates), but not for GET requests (retrieval)
		if r.Method == "POST" && !checkAuth(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		switch r.Method {
		case "GET":
			GetContent(w, r)
		case "POST":
			UpdateContent(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Authentication route
	r.HandleFunc("/api/authenticate", func(w http.ResponseWriter, r *http.Request) {
		// Handle CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		if r.Method == "POST" {
			Authenticate(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Image upload route
	r.HandleFunc("/api/upload-image", func(w http.ResponseWriter, r *http.Request) {
		// Handle CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		// Check authentication
		if !checkAuth(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if r.Method == "POST" {
			UploadImage(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Web routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			GetIndex(w, r)
		} else {
			// Serve static files for other paths
			http.FileServer(http.Dir("./")).ServeHTTP(w, r)
		}
	})
	r.HandleFunc("/dashboard", GetDashboard)

	// Start the server
	port := ":8085"
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

// UploadImage handles image uploads
func UploadImage(w http.ResponseWriter, r *http.Request) {
	// Maximum upload size of 10 MB
	r.ParseMultipartForm(10 << 20)

	// Get the file from the request
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create uploads directory if it doesn't exist
	os.MkdirAll("uploads", os.ModePerm)
	
	// Create destination file
	dst, err := os.Create("uploads/" + handler.Filename)
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
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success", 
		"message": "Image uploaded successfully", 
		"imagePath": "/uploads/" + handler.Filename,
	})
}
