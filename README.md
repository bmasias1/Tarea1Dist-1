# Tarea1Dist

Conexion Cliente (Máquina 10.10.28.155) y Logística (máquina 10.10.28.154) lista
Conexión Camion (aún no he usado ninguna máquina) y Logística lista
  - Todo probado localmente

Todos loss archivos están en la carpeta grpc:
  - logistica.go
  - cliente.go
  - camion.go
 
 Todos los archivos csv están en la carpeta archivos. El cliente lee desde retail.csv y pymes.csv. El archivo indexAct lo usa para ver el ID que le da a cada pedido. Al correr el Cliente se crea 'registro.csv' en el servidor, que indica las órdenes que han sido ingresadas por el cliente. Estos archivos se pueden ver con vercsv.go, el cual está en la carpeta 'Pruebas'.
 
 # Flujo del progrmama hasta ahora
 
 1) Se abre logistica (que está en grpc/server) en la máquina 10.10.28.154
 2) Se abre el cliente (en carpeta principal) en la máquina 10.10.28.155, este se conecta a través del puerto :50051
 3) Desde el cliente ya se puede usar el programa, si se ingresa un código de seguimiento, logística responderá que no hay órdenes ingresadas.
 4) Desde el cliente, con la opción 1 o 2 se ingresan órdenes. Luego de ingresar cada cuántos segundos estas se envían, estas llegan a logística y son guardadas en el archcivo results.csv (este archivo tiene la estructura de la primera tabla que aparece en el pdf). Al mismo tiempo que se escriben en este archivo, son agregadas a su cola correspondiente.
 5) Al ingresar un código de seguimiento, el servidor busca en cada una de las colas el código.
 6) Los camiones se conectan a través del puerto :50052. El servidor usa 'go routines' para abrir ambos al mismo tiempo (cosa que los camiones puedan despachar mientras se ingresan órdenes)
 7) Los camiones solicitan una órden a logística. Si las colas están vacías, espera 5 segundos antes de volver a pedir.
 8) Si el camion recibe una órden, espera el tiempo indicado antes de irse con una única órden. Estos la despachan, comunican a logística el resultado del despacho y escriben su csv correspondiente.

NOTA: No se logró implementar el que los camiones despacharan la órden de mayor valor primero.

 
