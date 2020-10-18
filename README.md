# Tarea 1

## WinduCloveer

## Maquinas Viruales:

El usuario de las máquinas es: sd

* Máquina 1: Logistica
	ip/hostname: 10.10.28.63
	contraseña: RQykXsIOZSDOuzd
 
* Máquina 2: Camiones
	ip/hostname: 10.10.28.64
	contraseña: FcPkvnGbEWEAlie
 
* Máquina 3: Clientes
	ip/hostname: 10.10.28.65
	contraseña: uNzXQlZUsGbKgND
 
* Máquina 4: Finanzas
	ip/hostname: 10.10.28.66
	contraseña: xysmRmDVuHkoWLk

# Consideraciones:

## logistica:
	* Se ignora el campo ID que envien los clientes, pues los clientes no deben definir el id, es labor del sistema.


## Camiones:
	* Para considerar si reintentar un paquete de retail se siguio la siguiente idea:
		1 intento -> costo 0
	  	2 intentos -> costo 10 
	  		Prioritarios 
	  			0.8*valor + 0.2*0.3*valor ganancia estimada 
	  				0.8*valor + 0.2*0.3*valor > 10 Condicion de 2 intentos
	  		Normales
	  			0.8*valor ganancia estimada 
	  				0.8*valor > 10 Condicion de 2 intentos
	  	3 intentos -> costo 20
	  		Prioritarios 
	  			0.8*valor + 0.2*0.3*valor ganancia estimada 
	  				0.8*valor + 0.2*0.3*valor > 20 Condicion de 3 intentos
	  		Normales
	  			0.8*valor ganancia estimada 
	  				0.8*valor > 20 Condicion de 3 intentos


## Clientes:
	* Los csv tendran formatos (entiendase como esquema o columnas) y nombres de archivos identicos a los de ejemplo.
	* Para las acciones del cliente se trabajarán mediante probabilidades definidas al inicio de la ejecución.