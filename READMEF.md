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

## Finanzas:
* La maquina tendrá montado correctamente un servidor de RabbitMQ, con usuario *WinduCloveer* y clave *secret* con permisos en *"/"* y administrador.
* El sistema debe terminar su ejecución luego de un tiempo definido de inactividad definido al iniciar la ejecución.
* El registro se mostrará cada una cantidad definida de segundos que se solicitará al inicio de la ejecución.
* El archivo de registro se hara en un arhivo llamado *registroFinanzas.csv*, el cual no llevará encabezado pues se considera sabido que el orden de columnas será:
	> Id Paquete | Descripcion | Tipo | Intentos | Estado | ValorOriginal | Ganancia/Costo 

# Ejecución

## Registros:
* Para limpiar los registros se incluye el comando *clearRegisters* en el Makefile.
	> make clearRegisters

## Finanzas:
* Para el complilado se incluye un Makefile. Del cual se requiere compilar *Finanzas* que creará un binario del mismo nombre en el directorio *bin*.
	> make Finanzas
* La forma de ejecucion es por linea de comandos. Estando en la carpeta raiz del sistema (\~/Tarea1) ejecutar:
	> ./bin/Finanzas