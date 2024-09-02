package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"golang.org/x/text/encoding/charmap"
)

// Вызов и исполнение py файла
func runMainPy(heading string) (string, error) {
	cmd := exec.Command("python", "main.py", heading)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении main.py: %v", err)
	}
	utf8Out, err := charmap.Windows1251.NewDecoder().Bytes(out)
	if err != nil {
		return "", fmt.Errorf("ошибка при преобразовании кодировки: %v", err)
	}
	return string(utf8Out), nil
}

// Функция для генерации аудиофайла из текста
func textToSpeech(text string) (string, error) {
	cmd := exec.Command("python", "tts.py", text)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении tts.py: %v", err)
	}
	return strings.TrimSpace(string(out)), nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env")
	}

	key := os.Getenv("TGKEY")
	if len(key) < 1 {
		log.Fatal("Отсутствует TGKEY из .env")
	}

	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		heading := update.Message.Text
		result, err := runMainPy(heading)
		if err != nil {
			result = fmt.Sprintf("Ошибка: %v", err)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
		bot.Send(msg)

		// Генерируем аудиофайл из текста
		textToSpeech(result)

		audioFilePath := "./output.mp3"
		audioFile, err := os.Open(audioFilePath)
		if err != nil {
			log.Printf("Ошибка при взятии файла: %v", err)
			continue
		}
		// Создаем объект для отправки аудио
		audioConfig := tgbotapi.NewAudio(update.Message.Chat.ID, tgbotapi.FileReader{
			Reader: audioFile,
			Name:   "audio.mp3", // Имя файла, которое будет отображаться в Telegram
		})
		audioConfig.Caption = "Это мое аудио сообщение"
		_, err = bot.Send(audioConfig)
		if err != nil {
			log.Println(err)
		}
		audioFile.Close()
	}
}
