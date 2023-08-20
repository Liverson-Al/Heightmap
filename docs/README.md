# Heightmap

 - EarthExplorer - отсутсвует API, следовательно нужно будет парсить и знать какие к каждой территории нужны датасеты 
(https://earthexplorer.usgs.gov/)
![image](https://github.com/Liverson-Al/Heightmap/assets/80148366/1859b4b2-7668-4a47-9d6b-d096286421c4)


 - Heightmapper - готовая карта высот всего мира в одном веб приложении с открытым кодом. API в моём понятии там нет, но можно использовать внутреннюю функцию renderView (371 строка) из файла main.js, который результатом работы выдаёт png снимок карты высот 
(https://tangrams.github.io/heightmapper/)
![image](https://github.com/Liverson-Al/Heightmap/assets/80148366/2d7da8de-2fde-40ac-a24d-6ad347007f83)


 - Dwtkns - Нужно парсить сайт, необходима регистрация на сайте NASA с последующей выдачей доступа к сервису. Даннные хранятся по тайлам. API отсутствует. Ещё один их сервис есть, но он больше не функционирует и все его ссылки умерли 
(https://dwtkns.com/srtm30m/)
![image](https://github.com/Liverson-Al/Heightmap/assets/80148366/f160b415-250f-4e61-8d9e-7c6df94178fc)


 - OpenTopography - у него есть API, но доступ к нему есть только у университетов США и снимки там в основном по территории США. Сейчас попробую запросить их ключ к API, может там больше возможностей 
(https://portal.opentopography.org/)
![image](https://github.com/Liverson-Al/Heightmap/assets/80148366/036c5fe1-0458-483a-a124-dd8ebc0fee21)


 - Opentopodata - бесплатный REST API сервер с кучей датасетов, в том числе и с глобальным - aster30m. Можно получить с помощью запросов или отдельно всё скачать
(https://www.opentopodata.org/)
![image](https://github.com/Liverson-Al/Heightmap/assets/80148366/2a29d1a7-e624-46af-9b2f-52d498458894)


 - Open-Elevation - бесплатный REST API сервер, датасет неизвестен. Документация - https://github.com/Jorl17/open-elevation/blob/master/docs/api.md, пример GET запроса - https://api.open-elevation.com/api/v1/lookup?locations=41.161758,-8.583933
(https://www.open-elevation.com/)


 - Yandex API - ограниченно-бесплатный сервис, но в рассмотренных продуктах не было замеченно получение высоты по координате. Но для последующей отрисовки глобальной карты можно воспользоваться их JavaScript API
(https://yandex.ru/dev/maps/commercial/?utm_source=undefined&utm_medium=undefined&utm_campaign=undefined#addition)



  Рассмотреть
   - Элементарный пространственный анализ - https://kpfu.ru/staff_files/F2134311690/KURS_LEKCIJ_GIT.pdf
   - https://ru.wikipedia.org/wiki/%D0%9A%D0%BE%D1%8D%D1%84%D1%84%D0%B8%D1%86%D0%B8%D0%B5%D0%BD%D1%82_%D0%BC%D0%B0%D1%81%D1%88%D1%82%D0%B0%D0%B1%D0%B0
   
