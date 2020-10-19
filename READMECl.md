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

## Clientes:
* Los csv con paquetes a entregar estarán en la carpeta *files* bajos los nombres *retail.csv* para paquetes de retail y *pymes.csv* para paquetes de pymes.
* Los csv tendran formatos (esquema de columnas) identicos a los de ejemplo.
* Para correcta lectura de archivos se deben tener dentro de una carpeta *files* el directorio desde el cual se ejecute.
	> * Si se ejecuta desde la carpeta raiz del sistema (\~/Tarea1) se utilizarán los archivos de ejemplo que se encuentran en "\~/Tarea1/files".   
	> * Si se ejecuta desde la carpeta "\~/Tarea1/bin" se requerirán archivos en "\~/Tarea1/bin/files".    
	> * **Se recomienda seguir las instrucciones de ejecucion y situarse en la carpeta raiz del proyecto (\~/Tarea1) para la ejecución.**
* Para las acciones del cliente se trabajarán mediante probabilidades definidas al inicio de la ejecución.

# Ejecución


## Clientes:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *Clientes* que creará un binario del mismo nombre en el directorio *bin*.
	> make Clientes
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea1) ejecutar:  
	> ./bin/Clientes

