# Imagem oficial do golang com suporte a go modules
FROM golang:1.14 AS development

# Acessando o diretório de trabalho
WORKDIR /app

# Copiando o projeto do host para o container
COPY . .

# Instalando o air para live-reload da API
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    | sh -s -- -b $(go env GOPATH)/bin

# Habilitando o modo live-reload
CMD $(go env GOPATH)/bin/air

# Build multi-stage para ambientes de produção
FROM development AS production

# Compilando o binário
RUN make build

# Habilitando a API
CMD ./banking
