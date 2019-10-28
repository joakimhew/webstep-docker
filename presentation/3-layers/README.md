```sh
docker image save imagename -o imagename.tar
mkdir extract
tar xvf imagename.tar -C extract
cd extract
docker history imagename
jq '.' manifest.json
```
