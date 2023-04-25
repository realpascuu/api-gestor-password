# API GESTOR PASSWORD
## Datos a tener cuenta
- Cuando se ejecuta, tener en cuenta que lo hace en el puerto 443, con lo cuál si hay algún servidor abirto rollo apache o nginx puede generar un conflicto de "puerto ya en uso".
- En sistemas Linux, se debe ejecutar como sudo ya que de forma predeterminada el sistema no deja usar puertos <= 1024 (los asignados por defecto). 
## Cómo empezar
Esta API usa el protocolo HTTPS, que junta HTTP y una comunicación segura mediante TLS. Para ello, se deben generar las claves pública y privada para que esta conexión funciones y sea segura.
Estas deben ir alojadas en la carpeta `certs`.

Para ello, ejecutamos en la terminal:
```bash
# PRIVATE KEY
# mediante el algoritmo RSA >= 2048 bits
openssl genrsa -out certs/server.key 2048

# mediante el algoritmo ECDSA
openssl ecparam -genkey -name secp384r1 -out certs/server.key

# PUBLIC KEY
# generación de la clave pública autofirmado (x509) basada en la clave privada
openssl req -new -x509 -sha256 -key certs/server.key -out certs/server.crt -days 3650
```

## Tener una base de datos
Se ha preparado el fichero Dockerfile para poder arrancar de una manera sencilla el motor de base de datos. Este es postgres.
Para crear la imagen, con el environment por defecto, ejecutamos:
```bash
docker build -t my-postgres-db ./
```

Una ver buildeada, solo quedaría crear el contenedor:
```bash
docker run -d --name gestorpassword-db -p 5432:5432 my-postgres-db
```

Con esto, ya tendríamos ejecutando una imagen de postgres en la versión indicada en el Dockerfile. 
- Los nombres tanto al buildear como al crear el contenedor como los puertos están puesto por defecto. Siempre que se quiera, se pueden cambiar.
- **IMPORTANTE**: Tener en cuenta los valores que se ponen, para tener en el archivo .env el valor de DATABASE_URI apuntando al contenedor.
- Por defecto, en el Dockerfile se asigna una contraseña root (usuario postgres) y una BD (gestorpassword). Esto también se puede cambiar.

## ¿Las tablas?
Las tablas se obtienen de manera sencilla: al ejecutar el programa. El servidor está preparada para realizar las migraciones en caso de que no estén creadas las tablas. Si se realizan cambios en los atributos de las tablas o en los valores, no se considera. Se aconseja hacer backup de los datos y resetar la BD.
 