package main

func computeDelay(Endpoints, s CacheServer, r Request) {
	var delay int64

	for index := 0; index < len(Endpoints); index++ {
		actualEndpoint = Endpoints[index]
		if r.Re != index { 
        //Este endpoint no hace esa request
            if actualEndpoint.CacheServers[index].C == s.C  {
                 //Que el endpoint este enganchado al CacheServer       
                 delay=delay+actualEndpoint.CacheServers[index].Lc
			}
            else{
                delay=delay+actualEndpoint.LD
            }

		}

	}

//delay completo de la request ejecutada por ese CacheServer
return delay
}


