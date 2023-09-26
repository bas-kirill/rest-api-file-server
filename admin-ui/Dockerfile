FROM node:17-alpine as build

WORKDIR /app
COPY package.json /app
COPY package-lock.json /app
RUN npm install

COPY . .
RUN npm run build

FROM nginx:1.25.2
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /usr/share/nginx/html
