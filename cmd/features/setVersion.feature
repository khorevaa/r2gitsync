# language: ru

Функционал: Установка версии в файл версий (команды set-version)
    Как Пользователь
    Я хочу выполнять автоматическую установку последней синхронизированной версии с хранилищем 1С
    Чтобы автоматизировать свою работы с хранилищем с git

Контекст: Тестовый контекст set-version
    Когда Я создаю временный каталог и сохраняю его в переменной "ПутьКаталогаИсходников"
    И Я инициализирую репозиторий в каталоге из переменной "ПутьКаталогаИсходников"
    И Я создаю тестовой файл "VERSION" в каталоге из переменной "ПутьКаталогаИсходников" с текстом:
    """
<?xml version="1.0" encoding="UTF-8"?>
<VERSION>0</VERSION>
    """
    И Я создаю временный каталог и сохраняю его в переменной "ВременнаяДиректория"

Сценарий: Установка версии без коммита
    Допустим Я добавляю параметр "--debug"
    И Я добавляю параметр "--tempdir" из переменной "ВременнаяДиректория"
    И Я добавляю параметр "set-version"
    И Я добавляю параметр "10"
    И Я добавляю параметр из переменной "ПутьКаталогаИсходников"
    Когда Я выполняю приложение
    Тогда Файл "VERSION" в каталоге из переменной "ПутьКаталогаИсходников" содержит "<VERSION>10</VERSION>"

Сценарий: Установка версии с коммитом
    Допустим Я добавляю параметр "--debug"
    И Я добавляю параметр "--tempdir" из переменной "ВременнаяДиректория"
    И Я добавляю параметр "set-version"
    И Я добавляю параметр "-c"
    И Я добавляю параметр "5"
    И Я добавляю параметр из переменной "ПутьКаталогаИсходников"
    Когда Я выполняю приложение
    Тогда Файл "VERSION" в каталоге из переменной "ПутьКаталогаИсходников" содержит "<VERSION>5</VERSION>"

Сценарий: Установка версии с использованием переменных окружения
    Допустим Я добавляю параметр "--debug"
    И Я добавляю параметр "--tempdir" из переменной "ВременнаяДиректория"
    И Я добавляю параметр "set-version"
    И Я добавляю параметр "1"
    И Я устанавливаю переменную окружения "GITSYNC_WORKDIR" из переменной "ПутьКаталогаИсходников"
    Когда Я выполняю приложение
    Тогда Файл "VERSION" в каталоге из переменной "ПутьКаталогаИсходников" содержит "<VERSION>1</VERSION>"
    И Я очищаю значение переменных окружения
    |GITSYNC_WORKDIR|