# Tarea1Dist

- Cliente (10.10.28.155) listo
- Todo está dentro de la carpeta grpc, tanto el cliente como logística

# Flujo de trabajo

1) Se abre logística en la máquina 10.10.28.154
2) Se abre cliente en la máquina 10.10.28.155
3) Si se selecciona la opción 3 (consultar estado de una orden) antes de ingresar alguna orden, la respuesta será "no hay órdenes ingresadas"
4) Las órdenes se leen desde la carpeta archivos. Estas van generando un archivo llamado "registro.csv" que tiene la estructura de la priemra tabla del PDF. Al mismo tiempo que ingresa cada órden al archivo, se va agregando esta a su cola correspondiente, con estado "En bodega".
