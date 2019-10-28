Förberedelser:

Starta 2 alpine shells:

`docker run --privileged --name webstep-alpine -it alpine sh`
`docker exec -it webstep-alpine sh`

Kopiera över node.tar.gz till /app
`mkdir /app`

`docker cp node.tar.gz webstep-alpine:/app/`

Installera dependencies:
`apk add --no-cache libgcc libstdc++ curl bash`

DEMO:

Skapa en ny shell mapp
`mkdir /new-shell`

Kopiera mappar som behövs till `new-shell`
`cd new-shell`
`cp -R /bin .`
`cp -R /lib .`
`cp -R /usr .`

Skapa en proc mount
`mkdir /proc`
`mount --bind /proc /new-shell/proc`

Skapa en fil:
`echo "Hello" > webstep`

Kör chroot från första terminalen
`chroot .`

Packa upp `node.tar.gz` från den andra terminalen
`cd /app`
`tar xzvf node.tar.gz -C /new-shell --strip-components=1`
