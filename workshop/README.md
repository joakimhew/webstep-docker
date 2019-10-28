# Workshop

## Instruktioner

I mappen `exercise` finns det mappar för dotnet, java, nodejs och go. Välj det som du är bekväm med, eller något du inte är bekväm med för lite extra utmaning ;)

Varje mapp har en `Dockerfile` och en `docker-compose.yml` som du behöver fylla i. Skulle du fastna så finns svaren i `solution` mappen.

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
