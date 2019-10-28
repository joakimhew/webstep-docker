1. Bygg image utan cache och utan att ta bort intermediate containers:
   `docker build --no-cache --rm=false . -t webstep-nodejs-example-image:latest`
2. Visa intermediate containers:
   `docker ps -a --no-trunc --format "table {{.ID}}\t{{.Command}}"`
