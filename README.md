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