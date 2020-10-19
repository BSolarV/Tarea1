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

## logistica:
* Se ignora el campo ID que envien los clientes, pues los clientes no deben definir el id, es labor del sistema.

# Ejecución

## Logística:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *Logistica* que creará un binario del mismo nombre en el directorio *bin*.
	> make Logistica
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea1) ejecutar:  
	> ./bin/Logistica
