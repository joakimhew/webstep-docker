# Workshop

## Instruktioner

I mappen `exercise` finns det mappar för dotnet, java, nodejs och go. Välj det som du är bekväm med, eller något du inte är bekväm med för lite extra utmaning ;)

Varje mapp har en `Dockerfile` och en `docker-compose.yml` som du behöver fylla i. Skulle du fastna så finns svaren i `solution` mappen.

Testa att bygga din image och starta en container genom att köra:
`docker-compose up --build`

När du lyckats bygga imagen är det dags att ladda upp den till docker hub. Detta gör du genom att :

1. Gå till https://hub.docker.com/
2. Tryck på `Create repository` och skapa en ny repository
3. `docker login`
4. `docker tag local-image:tagname new-repo:tagname`
5. `docker push new-repo:tagname`

Nu kan vi testa att köra den nya imagen i en ny miljö:

1. Gå till https://labs.play-with-docker.com/
2. Skapa en ny instans (Add new instance)
3. Hämta imagen som du precis pushade: `docker pull new-repo:tagname`
4. Starta imagen genom att köra `docker run ...`

## Tips

### Dockerfile

Referens: https://docs.docker.com/engine/reference/builder/

Vanliga instruktioner:

```DockerFIle
FROM # Används för imagens bas
WORKDIR # Används för att
COPY # Används för att kopiera något från lokalt till image
RUN # Används för att köra kommandon för imagen
CMD [ ] # Används för att specificera vad containern skall köra vid start
```

### Docker-compose

Exempelfil:

```docker
version: "3"
services:
  servicenamn:
    build: .
    ports:
      - "80:80"
```

### Kommandon

```sh
docker images # Lista images
docker image build . -t namn:tagg # Bygg images
docker run -p 80:80 -it namn:tagg # Starta en container med port 80 ->  80
```
