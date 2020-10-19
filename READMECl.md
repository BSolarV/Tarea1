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
	> Finanzas -> Maquina 4: 10.10.28.66  
* Los sistemas requieren un criterio de cierre que se definiria a criterio del grupo para cada cual.
* Se muestra en pantalla información que puede resultar útil para el usuario.
* Se supone todas las entradas a los sistemas serán validas (sin *typos* o dentro de valores realistas y no negativos)

## Clientes:
* El sistema pregunta al inicia que tipo de cliente simulará (cliente Retail o cliente Pyme).
* El sistema debe terminar su ejecucion cuando termine de enviar todos los paquetes que tenga en su csv respectivo (csv para *Retail* y *Pymes*)
* Se pueden tener multiples clientes enviando paquetes, pero solo se leerá el archivo correspondiente al tipo de cliente seleccionado al inicio.
* Tanto el codigo de seguimiento como el estado del paquete se mostrarán en pantalla debido a que no se requiere un archivo de registros para la instancia del sistema.
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
