FROM node:latest

EXPOSE 3000

WORKDIR /usr/src/battleship-client-web

COPY package.json package-lock.json ./
RUN npm install

COPY . .

CMD ["/usr/local/bin/npm", "start"]
