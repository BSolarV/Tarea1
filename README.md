# Tarea 1

## WinduCloveer

## Maquinas Viruales:

El usuario de las máquinas es: sd

* Máquina 1: Logistica  
	> ip/hostname: 10.10.28.63  
	> contraseña: RQykXsIOZSDOuzd
 
* Máquina 2: Camiones
	> ip/hostname: 10.10.28.64  
	> contraseña: FcPkvnGbEWEAlie
 
* Máquina 3: Clientes  
	> ip/hostname: 10.10.28.65  
	> contraseña: uNzXQlZUsGbKgND
 
* Máquina 4: Finanzas
	> ip/hostname: 10.10.28.66  
	> contraseña: xysmRmDVuHkoWLk

# Consideraciones:
## Generales  
* Las conexiones se definen en base a las maquinas virtuales asignadas, de forma que los ejecutables funcionarán efectivamente en su máquina correspondiente.  

	> Logistica.exe -> Maquina 1: 10.10.28.63  
	> Clientes.exe -> Maquina 2: 10.10.28.64  
	> Camiones.exe -> Maquina 3: 10.10.28.65  
	> Finanzas.exe -> Maquina 1: 10.10.28.66  

## logistica:
* Se ignora el campo ID que envien los clientes, pues los clientes no deben definir el id, es labor del sistema.


## Camiones:
* Para considerar si reintentar un paquete de retail se siguio la siguiente idea:  
	* 1 intento -> costo 0  
	* 2 intentos -> costo 10   
		* Prioritarios
			* *0.8\*valor + 0.2\*0.3\*valor* = ganancia estimada   
			* Condicion de 2 intentos: 
				> *0.8\*valor + 0.2\*0.3\*valor* > 10  
		* Normales
			* *0.8\*valor* = ganancia estimada   
			* Condicion de 2 intentos:   
				> *0.8\*valor > 10*
	* 3 intentos -> costo 20  
		* Prioritarios
			* *0.8\*valor + 0.2\*0.3\*valor* = ganancia estimada
			* Condicion de 3 intentos: 
				> 0.8\*valor + 0.2\*0.3\*valor > 20 
		* Normales
			* 0.8*valor = ganancia estimada
			* Condicion de 3 intentos: 
				> 0.8*valor > 20 


## Clientes:
* Los csv con paquetes a entregar estarán en la carpeta *files* bajos los nombres *retail.csv* para paquetes de retail y *pymes.csv* para paquetes de pymes.
* Los csv tendran formatos (esquema de columnas) identicos a los de ejemplo.
* Para las acciones del cliente se trabajarán mediante probabilidades definidas al inicio de la ejecución.