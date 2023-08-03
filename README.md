# Heightmap

 - EarthExplorer - отсутсвует API, следовательно нужно будет парсить и знать какие к каждой территории нужны датасеты 
(https://earthexplorer.usgs.gov/)


 - Heightmapper - готовая карта высот всего мира в одном веб приложении с открытым кодом. API в моём понятии там нет, но можно использовать внутреннюю функцию renderView (371 строка) из файла main.js, который результатом работы выдаёт png снимок карты высот 
(https://tangrams.github.io/heightmapper/)


 - Dwtkns - Нужно парсить сайт, необходима регистрация на сайте NASA с последующей выдачей доступа к сервису. Даннные хранятся по тайлам. API отсутствует. Ещё один их сервис есть, но он больше не функционирует и все его ссылки умерли 
(https://dwtkns.com/srtm30m/)


 - OpenTopography - у него есть API, но доступ к нему есть только у университетов США и снимки там в основном по территории США. Сейчас попробую запросить их ключ к API, может там больше возможностей 
(https://portal.opentopography.org/)


  Рассмотреть
  https://www.opentopodata.org
  https://www.open-elevation.com
  https://en.wikipedia.org/wiki/Shuttle_Radar_Topography_Mission
  https://project.linuxfoundation.org/hubfs/Overture%20Maps/Building%20Heights%20Whitepaper_041423.pdf
