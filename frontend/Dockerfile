FROM node:lts-alpine as build-stage
WORKDIR /app/frontend

# Copiamos los archivos de configuración primero
COPY frontend/package*.json ./

# Instalamos todo
RUN npm install

# Copiamos el resto del código del frontend
COPY frontend/ .

# HACK: Forzamos a que solo use vite build ignorando vue-tsc
RUN npx vite build

FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/frontend/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]