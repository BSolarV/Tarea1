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

## logistica:
* El sistema debe terminar su ejecución luego de un tiempo definido de inactividad ingresado al iniciar la ejecución.
	> Tiempo de inactividad es considerado como el *gap* de tiempo entre acciones críticas (recibir orden de cliente; Enviar paquete a camión; Recibir paquete de camión)
	> Tiempo máximo de inactividad débe ser de al menos 7 veces el tiempo de viaje de los camiones. (ej: si el tiempo de viajes de camiones es de 60 segundos, el mínimo tiempo máximo de inactividad para logística sería 7 minutos)
* El codigo de seguimiento no requiere un formato definido, por lo cual se utilizará el ID del Paquete en el sistema. En caso de ser Retail el codigo de seguimiento será 0 pues no peden ser seguidos.
* Solamente se puede hacer seguimiento de paquetes recibidos en la instancia actual del sistema y no de paquetes escritos en el archivo registro por instancias previas.
* Se ignora el campo ID que envien los clientes, pues los clientes no deben definir el id, es labor del sistema.
* El archivo de registro se hara en un arhivo llamado *registroLogistica.csv*, el cual no llevará encabezado pues se considera sabido que el orden de columnas será:
	> timestamp | ID Paquete | Tipo | Producto(Descripcion) | Valor | Origen | Destino | CodigoDeSeguimiento

# Ejecución

## Registros:
* Para limpiar los registros se incluye el comando *clearRegisters* en el Makefile.
	> make clearRegisters

## Logística:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *Logistica* que creará un binario del mismo nombre en el directorio *bin*.
	> make Logistica
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea1) ejecutar:  
	> ./bin/Logistica

