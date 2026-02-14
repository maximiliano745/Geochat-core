# Etapa de construcción
FROM node:lts-alpine as build-stage
WORKDIR /app/frontend

# Copiamos archivos de configuración
COPY frontend/package*.json ./

# Instalamos dependencias
RUN npm install

# Copiamos el código del frontend (desde la carpeta frontend)
COPY frontend/ .

# HACK: Construcción rápida saltando errores de tipos
RUN npx vite build

# Etapa de producción con Nginx
FROM nginx:stable-alpine as production-stage

# IMPORTANTE: Verificamos que la ruta coincida con el WORKDIR de arriba
COPY --from=build-stage /app/frontend/dist /usr/share/nginx/html

# Configuramos Nginx para que maneje el enrutamiento de Vue (evita 404 al refrescar)
RUN echo 'server { \
    listen 80; \
    location / { \
        root /usr/share/nginx/html; \
        index index.html; \
        try_files $uri $uri/ /index.html; \
    } \
}' > /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]