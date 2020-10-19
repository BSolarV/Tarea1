# Tarea 1

## WinduCloveer
	> Carlos Jara Almendra - 201773036-5
	> Bastián Solar Vargas - 201773003-k


# Consideraciones:
## Generales  
* Las conexiones se definen en base a las maquinas virtuales asignadas, de forma que los ejecutables funcionarán efectivamente en su máquina correspondiente.  
	> Logistica -> Maquina 1: 10.10.28.63  
	> Clientes -> Maquina 2: 10.10.28.64  
	> Camiones -> Maquina 3: 10.10.28.65  
	> Finanzas -> Maquina 1: 10.10.28.66  


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


# Ejecución

## Camiones:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *Camiones* que creará un binario del mismo nombre en el directorio *bin*.
	> make Camiones
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea1) ejecutar:  
	> ./bin/Camiones 

