Телеграмм Бот позволяющий списать с вашего док файла

**Как работает**?
1. Он ищёт в файле первое упоминание указанной фразы
2. Формирует текст границой которого выступает Заголовок\
![Пример текста](https://github.com/zongrade/botMath/raw/main/text.jpg)
3. Формирует аудио файл, чтобы можно было в наушниках списать\
![Пример аудио](https://github.com/zongrade/botMath/raw/main/audio.jpg)
4. Отправляет в качестве ответа
 
**Как использовать**?\
Нужно иметь 2 файла в той же директории, что и клонированный репозиторий
____
.env в нём указать
TGKEY=<Ключ от бота>
____
pod.docx
он должен быть хорошо структурирован!!!
____

**Как должен быть оформлен docx файл?**
1. Каждый отдельный билет/вопрос/тема должна начинаться с нового заголовка\
![Вот такой заголовок](https://github.com/zongrade/botMath/raw/main/otdelenie.jpg)\
![Вот такой заголовок](https://github.com/zongrade/botMath/raw/main/zagolovok.jpg)
2. Избегайте специальных символов (пример: ↓), они будут проигнорированы при текстовом ответе
3. Отсутствие картинок

**Что по аудио файлу?**\
Аудио файл учтёт все символы и будет озвучен Google Text-to-Speech

Для адекватной работы не рекомендуется использовать слишком длинные темы. Аудио файл примерно на 1-2 страницы будет звучать около 3 минут и гененироваться ~20 секунд
