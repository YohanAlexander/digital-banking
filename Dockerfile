# Imagem oficial do golang com suporte a go modules
FROM golang:1.14 AS development

# Acessando o diretório de trabalho
WORKDIR /app

# Copiando o projeto do host para o container
COPY . .

# Instalando compile daemon para live-reload da API
RUN GO111MODULE=off go get github.com/githubnemo/CompileDaemon

# Habilitando o modo live-reload
CMD CompileDaemon --build="make build" --command=./main

# Build multi-stage para ambientes de produção
FROM development AS production

# Compilando o binário
RUN make build

# Habilitando a API
CMD ./main
