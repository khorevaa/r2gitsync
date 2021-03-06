# language: ru
@ext
Функциональность: Синхронизация хранилища расширения конфигурации 1С и гит (команды sync)
  Как Пользователь
  Я хочу выполнять автоматическую синхронизацию конфигурации из хранилища расширения
  Чтобы автоматизировать свою работы с хранилищем с git

  Контекст: Тестовый контекст синхронизации
    Когда Я создаю временный каталог и сохраняю его в переменной "КаталогХранилища1С"
    И я скопировал каталог "./tests/fixtures/extension_storage" в каталог из переменной "КаталогХранилища1С"
    И Я создаю временный каталог и сохраняю его в переменной "ПутьКаталогаИсходников"
    И Я инициализирую репозиторий в каталоге из переменной "ПутьКаталогаИсходников"
    И Я создаю тестовой файл AUTHORS
    И Я записываю "0" в файл VERSION

  Сценарий: Синхронизация хранилища расширения с git-репозиторием
    Допустим Я устанавливаю авторизацию в хранилище пользователя "Администратор" с паролем ""
    И Я устанавливаю версию платформы "8.3"
    Когда Я выполняю выполняют синхронизацию для расширения "test"
    Тогда Файл VERSION содержит 4
