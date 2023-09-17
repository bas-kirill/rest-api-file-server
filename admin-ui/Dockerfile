FROM node:17-alpine
WORKDIR /app
COPY tsconfig.json /app
COPY package.json /app
COPY src/ /app/src
COPY public/ /app/public
RUN npm install
CMD ["npm", "start"]
