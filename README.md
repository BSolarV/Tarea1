# Tarea 1

## WinduCloveer
> Carlos Jara Almendra - 201773036-5
> Bastián Solar Vargas - 201773003-k

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
	> Logistica -> Maquina 1: 10.10.28.63  
	> Clientes -> Maquina 2: 10.10.28.64  
	> Camiones -> Maquina 3: 10.10.28.65  
	> Finanzas -> Maquina 1: 10.10.28.66  


## logistica:
* Se ignora el campo ID que envien los clientes, pues los clientes no deben definir el id, es labor del sistema.


## Clientes:
* Los csv con paquetes a entregar estarán en la carpeta *files* bajos los nombres *retail.csv* para paquetes de retail y *pymes.csv* para paquetes de pymes.
* Los csv tendran formatos (esquema de columnas) identicos a los de ejemplo.
* Para correcta lectura de archivos se deben tener dentro de una carpeta *files* el directorio desde el cual se ejecute.
	> * Si se ejecuta desde la carpeta raiz del sistema (\~/Tarea1) se utilizarán los archivos de ejemplo que se encuentran en "\~/Tarea1/files".   
	> * Si se ejecuta desde la carpeta "\~/Tarea1/bin" se requerirán archivos en "\~/Tarea1/bin/files".    
	> * **Se recomienda seguir las instrucciones de ejecucion y situarse en la carpeta raiz del proyecto (\~/Tarea1) para la ejecución.**
* Para las acciones del cliente se trabajarán mediante probabilidades definidas al inicio de la ejecución.

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

## Finanzas:
* El registro se mostrará cada una cantidad definida de segundos que se solicitará al inicio de la ejecución.

# Ejecución

## Registros:
* Para limpiar los registros se incluye el comando *clearRegisters* en el Makefile.
	> make clearRegisters

## Logística:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *Logistica* que creará un binario del mismo nombre en el directorio *bin*.
	> make Logistica
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea1) ejecutar:  
	> ./bin/Logistica

## Clientes:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *Clientes* que creará un binario del mismo nombre en el directorio *bin*.
	> make Clientes
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea1) ejecutar:  
	> ./bin/Clientes

## Camiones:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *Camiones* que creará un binario del mismo nombre en el directorio *bin*.
	> make Camiones
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea1) ejecutar:  
	> ./bin/Camiones 

## Finanzas:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *Finanzas* que creará un binario del mismo nombre en el directorio *bin*.
	> make Finanzas
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea1) ejecutar:
	> ./bin/Finanzas