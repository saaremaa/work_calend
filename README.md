# work-calend

work-calend

Простой сервис, который выдает является ли день рабочим.

В качестве данных используется Производственный календарь → https://data.gov.ru/opendata/7708660670-proizvcalendar

# Запуск
1. git clone https://github.com/saaremaa/work_calend
2. ***docker-compose build***
3. Копируем docker-compose.yml например в папку /home/docker/app/work_api
4. Создаем папку дата и кладем туда календарь в формате CSV
4. ***docker-compose up -d***
5. Просмотр логов по команде ***journalctl -fxe CONTAINER_NAME=work_api***