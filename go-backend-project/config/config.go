package config

import (
	"log"
	"os"

	"github.com/joho/godotenv" // Bu paket .env dosyasını okumak için kullanılıyor
)

// Config yapısı, uygulama konfigürasyon ayarlarını tutar
type Config struct {
	Port     string // Uygulamanın dinleyeceği port
	Database string // Veritabanı bağlantı dizesi
}

// LoadConfig fonksiyonu, .env dosyasından konfigürasyon ayarlarını yükler
func LoadConfig() Config {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file") // Eğer dosya yüklenemezse hata ver
	}

	// Config yapısını oluştur
	return Config{
		Port:     getEnv("PORT", "8080"), // PORT değişkenini al, yoksa varsayılan 8080 kullan
		Database: getEnv("DATABASE_URL", "user:password@tcp(localhost:3306)/dbname"), // DATABASE_URL değişkenini al
	}
}

// getEnv fonksiyonu, belirtilen anahtara göre ortam değişkenini alır
// Eğer ortam değişkeni yoksa varsayılan değeri döner
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value // Eğer ortam değişkeni varsa onu döner
	}
	return defaultValue // Yoksa varsayılan değeri döner
}