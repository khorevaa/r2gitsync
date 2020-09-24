# language: ru

Функциональность: Проверка математики

  Структура сценария: add two digits

    Допустим I add <one> and <two>
    Когда I add <one> and <two>
    Тогда I the result should equal <res>

    Примеры:
      | one | two | res |
      | 2   | 2   | 8   |
      | 5   | 2   | 14  |

  Сценарий: add two digits
    Когда I add 1 and1 2
    Тогда I the result should equal 3
