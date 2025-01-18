FROM node:23
WORKDIR /app

COPY client .

RUN npm install
CMD [ "npm", "run", "dev", "--", "--host"]
