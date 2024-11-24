**Objetivo:** Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em
graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

# Setup

Copie o arquivo **.env.example** para **.env** e preencha as variáveis de ambiente.

```shell
cp .env.example .env
```

Para executar a aplicação localmente, basta executar o comando:

```shell
daocker-compose up
```
# Requisições

Para facilitar os testes, foi criado um arquivo **api/requests.http** que contém as requisições HTTP para testar a
aplicação localmente e no CloudRun.

# Aplicação

A aplicação foi estruturada baseada no **DDD**.

A camada lógica pode ser encontrada no diretório **internal**.

A camada interface/presentation pode ser encontrada no pacote `presentation`. Nessa camada encontra-se o **handler**
responsável por lidar com as requisições HTTP.

A camada services pode ser encontrada no pacote `application`. Nessa camada encontram-se o **ZipcodeService** e o
**WeatherService**.

A camada domain pode ser encontrada no pacote `domain` que engloba também os pacotes `entity` e `repository`. No
pacote `repository` encontra-se a interface `ZipcodeRepository` e a interface `WeatherRepository`.

A camada infrastructure pode ser encontrada no pacote `infrastructure`. Nessa camada encontra-se a implementação dos
repositórios `ViaCep` e `WeatherApi`.

# Testes

Os testes validam o comportamento do **handler** da requisição HTTP. Para executar os testes, basta executar o comando:

```shell
go test ./internal/presentation
```

# CodeRun

A aplicação pode ser conferida funcionando no endereço:
https://fullcycle-go-labs-1-cloudrun-866933005138.us-central1.run.app