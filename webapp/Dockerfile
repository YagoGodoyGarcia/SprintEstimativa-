# Etapa de build
FROM node:18-alpine AS build

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

# Gera o build da aplicação Next.js
RUN npm run build

# Etapa de produção
FROM node:18-alpine

WORKDIR /app

COPY --from=build /app ./

ENV NODE_ENV=production

# Expõe a porta padrão do Next.js
EXPOSE 3000

# Inicia o Next.js
CMD ["npm", "run", "start"]
