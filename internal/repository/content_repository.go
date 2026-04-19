package repository

import (
	"encoding/json"
	"os"

	"nikma/internal/models"
)

// ContentRepository handles content data persistence
type ContentRepository struct {
	contentPath string
}

// NewContentRepository creates a new content repository
func NewContentRepository(contentPath string) *ContentRepository {
	return &ContentRepository{contentPath: contentPath}
}

// Load loads content from the JSON file
func (r *ContentRepository) Load() (*models.ContentData, error) {
	// Check if the file exists
	if _, err := os.Stat(r.contentPath); os.IsNotExist(err) {
		// If file doesn't exist, return default content
		defaultContent := r.getDefaultContent()
		if saveErr := r.Save(defaultContent); saveErr != nil {
			return nil, saveErr
		}
		return defaultContent, nil
	}

	// Read the file
	data, err := os.ReadFile(r.contentPath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data
	var content models.ContentData
	if err := json.Unmarshal(data, &content); err != nil {
		return nil, err
	}

	return &content, nil
}

// Save saves content to the JSON file
func (r *ContentRepository) Save(content *models.ContentData) error {
	data, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.contentPath, data, 0644)
}

// getDefaultContent returns the default content
func (r *ContentRepository) getDefaultContent() *models.ContentData {
	return &models.ContentData{
		Hero: models.HeroSection{
			Title:           "Bimbel Tuntas, Nilai Pintas.",
			Subtitle:        "Bimbingan Matematika & IPA oleh Sarjana Fisika.",
			Description:     "Bimbingan intensif dan personal untuk siswa SD dan SMP. Ubah kesulitan belajar menjadi prestasi nyata, fokus pada pemahaman konsep dasar.",
			WhatsappNumber:  "6281234567890",
			WhatsappMessage: "Halo Kak Nikma, saya tertarik dengan bimbel Belajar Bareng Nikma. Saya ingin mendaftar dan mendapatkan sesi perkenalan gratis.",
		},
		About: models.AboutSection{
			Title:        "Kenalan dengan Kak Nikma",
			Description1: "Kak Nikma adalah lulusan Sarjana Fisika yang memiliki <span class='font-semibold text-gold'>passion mendalam dalam mengajar Matematika dan IPA (Fisika/Biologi)</span>.",
			Description2: "Dengan metode pengajaran yang sabar, terstruktur, dan fokus pada pemecahan masalah, Kak Nikma membantu membangun kepercayaan diri siswa SD dan SMP dalam belajar.",
			Description3: "Tujuan utamanya adalah mengubah kesulitan belajar menjadi prestasi nyata melalui pemahaman konsep dasar yang kuat.",
		},
		Program: models.ProgramSection{
			Title: "Program Belajar Bareng Nikma",
			Sd: models.ProgramItem{
				Title:       "Program SD (Kelas 4-6)",
				Description: "Fokus pada Dasar-dasar Matematika dan Sains.",
				Features: []string{
					"Penguatan Calitung (Catur, Literasi, Hitung)",
					"Pemahaman konsep dasar Matematika",
					"Latihan soal rutin",
				},
			},
			Smp: models.ProgramItem{
				Title:       "Program SMP (Kelas 7-9)",
				Description: "Fokus pada Matematika dan Fisika.",
				Features: []string{
					"Pemecahan Masalah Aljabar",
					"Konsep Dasar Fisika",
					"Persiapan Ujian Sekolah/Daerah",
				},
			},
		},
		Gallery: models.GallerySection{
			Title: "Galeri Kegiatan Belajar Bareng Nikma",
			Items: []models.GalleryItem{
				{Title: "Aktivitas Belajar Interaktif", Image: "/assets/images/gallery1.jpg"},
				{Title: "Murid-Murid Bahagia", Image: "/assets/images/gallery2.jpg"},
				{Title: "Sesi Belajar Kelompok", Image: "/assets/images/gallery3.jpg"},
				{Title: "Penghargaan Prestasi", Image: "/assets/images/gallery4.jpg"},
				{Title: "Kegiatan Praktikum IPA", Image: "/assets/images/gallery5.jpg"},
				{Title: "Sesi Evaluasi Mingguan", Image: "/assets/images/gallery6.jpg"},
			},
		},
		Testimonials: models.TestimonialsSection{
			Title: "Apa Kata Mereka?",
			Items: []models.TestimonialItem{
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
		Contact: models.ContactSection{
			Title:       "Hubungi Kami",
			Description: "Siap memulai perjalanan belajar yang menyenangkan dan efektif? Hubungi kami melalui WhatsApp!",
			ServiceArea: "Online dan Tatap Muka Area Bunder, Banyuwangi",
			ButtonText:  "Hubungi via WhatsApp",
		},
		Footer: models.FooterSection{
			Text: "&copy; 2025 Belajar Bareng Nikma. Hak Cipta Dilindungi.",
		},
	}
}
