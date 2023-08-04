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


  Рассмотреть
   - https://www.opentopodata.org
   - https://www.open-elevation.com
   - https://en.wikipedia.org/wiki/Shuttle_Radar_Topography_Mission
   - https://project.linuxfoundation.org/hubfs/Overture%20Maps/Building%20Heights%20Whitepaper_041423.pdf
   - https://yandex.ru/dev/maps/commercial/?utm_source=undefined&utm_medium=undefined&utm_campaign=undefined
