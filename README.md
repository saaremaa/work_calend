# work-calend

work-calend

Простой сервис, который выдает является ли день рабочим.

В качестве данных используется Производственный календарь → https://data.gov.ru/opendata/7708660670-proizvcalendar

# Запуск
1. git clone https://mo.enicom.ru:3001/krb/work-calend.git
2. ***docker-compose build***
3. Копируем docker-compose.yml например в папку /home/docker/app/work_api
4. Создаем папку дата и кладем туда календарь в формате CSV
4. ***docker-compose up -d***
5. Просмотр логов по команде ***journalctl -fxe CONTAINER_NAME=work_api***

# API

- "/" -  стартовая страница при обращении выкидывает на сайт компании
- "/api/v1/health"- проверка доступности API для Docker
- "/api/v1/stat" - данные по датасету
- "/api/v1/stats"- статистика по обращению к сервису
- "/api/v1/check_day/:date" - проверка является ли день выходным (пример http://127.0.0.1:2081/api/v1/check_day/15.05.2020)