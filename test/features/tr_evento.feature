Feature: Validate API responses
  EVENTOS_CAMPUS_CRUD
  probe JSON reponses


Scenario Outline: To probe route code response  /tr_evento
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      

  Examples: 
  |method|route        |bodyreq               |codres       |
  |GET   |/v1/tr_evento|./files/req/Vacio.json|404 Not Found|
  |GET   |/v1/tr_event |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/tr_event |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/tr_event |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/tr_event |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /tr_evento       
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      
  And the response should match json "<bodyres>"

  Examples: 
  |method|route        |bodyreq               |codres         |bodyres                |
  |GET   |/v1/tr_evento|./files/req/Vacio.json|404 Not Found  |./files/res7/Ierr0.json|
  |POST  |/v1/tr_evento|./files/req/Vacio.json|400 Bad Request|./files/res7/Ierr1.json|
   