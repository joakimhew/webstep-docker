# base image
FROM node:alpine

WORKDIR /app
COPY package.json /app/package.json
RUN npm install --silent
COPY . .

CMD ["npm", "start"]