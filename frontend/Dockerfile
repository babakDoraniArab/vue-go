FROM node:16 as dev
WORKDIR /usr/src/app/my-app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run dev

# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=dev /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]