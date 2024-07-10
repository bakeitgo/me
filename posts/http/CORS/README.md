## CORS

* CORS (Cross Origin Resource Sharing) is a protocol layered on top of HTTP, which adds limitations onto web browser behaviour when it comes to requesting resources from domain another than actual web site is hosted on.
    * By default browser requests API like fetch() and XMLHttpRequest follow **same-origin policy** so you cant make a request to other origin **unless server to which you want to make request, allows web site domain for requests**
    * So, when you can make it? Browser sends request to the "domain different" server with Origin="originFromWhichImMakingRequest" header, and server responds with requested resource and with Access-Control-Allow-Origin="theDomainListForWhichIAllowToReadResponse" header. If the domain From request Origin header matches the domain list in Access-Control-Allow-Origin, the response with resource is successfully going onto browser, if not (guess what), its not.
    * The upper scenario is for GET or POST request with certain MIME types, for rest methods the "preflight" OPTION request to the cross origin server, and upon approval the real one.

