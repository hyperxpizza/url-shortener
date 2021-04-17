FROM node:alpine
WORKDIR /app
COPY package*.json /app/
RUN npm install
COPY . ./app/
RUN npm run build

FROM nginx:alpine as prod
EXPOSE 3000
COPY --from=build /root/build /usr/share/nginx/html
COPY nginx.conf /etv/nginx/nginx.conf
