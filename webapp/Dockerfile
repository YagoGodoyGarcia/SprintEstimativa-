# Etapa de build
FROM node:18-alpine AS build

WORKDIR /app

# Declara os argumentos de build que serão passados pelo docker-compose
ARG NEXT_PUBLIC_BASE_URL
ARG NEXT_PUBLIC_BASE_WS_URL

# Torna os argumentos de build disponíveis como variáveis de ambiente para o processo de build
ENV NEXT_PUBLIC_BASE_URL=$NEXT_PUBLIC_BASE_URL
ENV NEXT_PUBLIC_BASE_WS_URL=$NEXT_PUBLIC_BASE_WS_URL

COPY package*.json ./
RUN npm install

COPY . .

# Gera o build da aplicação Next.js, utilizando as variáveis de ambiente definidas acima
RUN npm run build

# Etapa de produção
FROM node:18-alpine

WORKDIR /app

COPY --from=build /app ./

# As variáveis NEXT_PUBLIC_* também precisam estar disponíveis em tempo de execução
# Elas já são passadas pelo docker-compose na seção 'environment' do serviço webapp
ENV NODE_ENV=production

# Expõe a porta padrão do Next.js
EXPOSE 3000

# Inicia o Next.js
CMD ["npm", "run", "start"]

