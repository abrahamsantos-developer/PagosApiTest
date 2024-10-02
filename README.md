# PagosApiTest

PagosApiTest es una API REST desarrollada en Golang utilizando el framework Gin. La API permite gestionar comercios, transacciones, calcular comisiones y obtener ganancias. El proyecto está completamente dockerizado y usa PostgreSQL como base de datos.

## Requisitos
Asegúrate de tener Docker y Docker Compose instalados en tu máquina.

### Instalación de Docker y Docker Compose

Si no tienes Docker y Docker Compose instalados, sigue estos pasos:

#### Instalar Docker:

En sistemas Linux (Ubuntu/Debian):

```bash
sudo apt update
sudo apt install docker.io
```



En macOS y Windows, descarga Docker Desktop desde la página oficial de Docker.

Instalar Docker Compose:

En sistemas Linux (Ubuntu/Debian):
```bash
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```
En macOS y Windows, Docker Compose viene incluido con Docker Desktop.

Verificar la instalación:
Asegúrate de que Docker y Docker Compose están correctamente instalados ejecutando los siguientes comandos:

```bash
docker --version
docker-compose --version
```


## Instrucciones para levantar el proyecto con Docker

#### 1. Clonar el repositorio

```bash
git clone https://github.com/abrahamsantos-developer/PagosApiTest.git
```
```bash
cd PagosApiTest
```

#### 2. Levantar los contenedores
Ejecuta el siguiente comando para construir la imagen y levantar la aplicación junto con PostgreSQL:

```bash
docker-compose up --build
```

Si es necesario, puedes agregar sudo antes de los comandos docker y docker-compose.
Esto levantará tanto la API como la base de datos en contenedores Docker.


#### 3. Acceder a la API
La API estará corriendo en http://localhost:3000.


#### 4. Documentación Swagger
Puedes acceder a la documentación interactiva de la API en http://localhost:3000/swagger/index.html.


Comandos útiles

Parar los contenedores:
```bash
docker-compose down
```

Reiniciar los contenedores:
```bash

docker-compose up --build
```
