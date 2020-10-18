# Tarea1Dist

Conexion Cliente (Máquina 10.10.28.155) y Logística (máquina 10.10.28.154) lista.

Conexión Camion (Máquina 10.10.28.156) y Logística lista.

Conexion Finanzas (Máquina 10.10.28.156) y Logística lista.

Todos los archivos están en la carpeta grpc, a excepcion de finanzas:
  - logistica.go
  - cliente.go
  - camion.go

Finanzas se encuentra en RabbitMQ:
  - finanzas.go

 Todos los archivos csv están en la carpeta archivos. El cliente lee desde retail.csv y pymes.csv. El archivo indexAct lo usa para ver el ID que le da a cada pedido. Al correr el Cliente se crea 'registro.csv' en el servidor, que indica las órdenes que han sido ingresadas por el cliente. Estos archivos se pueden ver con vercsv.go, el cual está en la carpeta 'Pruebas'.
 
 # Flujo del programa
 
 1) Se abre logistica que está en /grpc/, en la máquina 10.10.28.154.
 2) Se abre el cliente que está en /grpc/, en la máquina 10.10.28.155, este se conecta a través del puerto :50051.
 3) Se abre camion.go que esta en /grpc/ en la maquina 10.10.28.156.
 4) Se abre finanzas.go que se encuentra en /RabbitMQ/ en la maquina 10.10.28.157.
 5) Desde el cliente ya se puede usar el programa, si se ingresa un código de seguimiento, logística responderá que no hay órdenes ingresadas.
 6) Desde el cliente, con la opción 1 o 2 se ingresan órdenes. Luego de ingresar cada cuántos segundos estas se envían, estas llegan a logística y son guardadas en el archcivo results.csv (este archivo tiene la estructura de la primera tabla que aparece en el pdf). Al mismo tiempo que se escriben en este archivo, son agregadas a su cola correspondiente.
 7) Al ingresar un código de seguimiento, el servidor busca en cada una de las colas el código.
 8) Los camiones se conectan a través del puerto :50052. El servidor usa 'go routines' para abrir ambos al mismo tiempo (cosa que los camiones puedan despachar mientras se ingresan órdenes).
 9) Los camiones solicitan una órden a logística. Si las colas están vacías, espera 5 segundos antes de volver a pedir.
 10) Si el camion recibe una órden, espera el tiempo indicado antes de irse con una única órden. Estos la despachan, comunican a logística el resultado del despacho y escriben su csv correspondiente.
 11) Finanzas se queda recibiendo informacion de logistica constantemente, la cual va registrando en finanzas.csv
 12) Finanzas puede ser finalizada ingresando cualquier caracter, arrojando asi al final el balance en DigniPesos.

NOTA: No se logró implementar el que los camiones despacharan la órden de mayor valor primero.

NOTA2: Cualquier tipo de cliente puede solicitar el seguimiento de las ordenes, no solo pymes.
 
