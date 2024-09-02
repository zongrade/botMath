import sys
from docx import Document
from gtts import gTTS
import os

def extract_text_under_heading(doc_path, heading_text):
    doc = Document(doc_path)
    found = False
    result = []

    for para in doc.paragraphs:
        if para.style.name.startswith('Heading'):
            if found:
                break
            if heading_text in para.text:
                found = True
        elif found:
            result.append(para.text)

    if not found:
        raise ValueError(f"Заголовок '{heading_text}' не найден")

    return '\n'.join(result)

def text_to_speech(text, output_file):
    tts = gTTS(text, lang='ru')  # Используйте 'ru' для русского языка
    tts.save(output_file)

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Использование: python main.py 'Название заголовка'")
        sys.exit(1)

    doc_path = 'pod.docx'
    heading_text = sys.argv[1]

    try:
        # Извлекаем текст под заголовком
        text = extract_text_under_heading(doc_path, heading_text)
        
        # Сохраняем текст в аудиофайл
        audio_file = 'output.mp3'
        text_to_speech(text, audio_file)

        print(f"Текст под заголовком '{heading_text}':\n{text.encode('windows-1251', 'ignore').decode('windows-1251')}")

        print(f"Аудиофайл сохранен как {audio_file}")

    except ValueError as e:
        print(e)
