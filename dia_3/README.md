# Día 3 – Concurrencia y manejo de errores

## 👉 Objetivo: usar las características únicas de Go.

### Teoría (30 min):

- Error como tipo de retorno.
- Múltiples valores de retorno en funciones.
- Goroutines (`go` keyword).
- Canales (`chan`) y comunicación entre goroutines.
- select para multiplexar canales.

### Ejercicios (1h):

- Escribir una función que devuelva (resultado, error) al dividir dos números (manejar división por cero).
- Lanzar 3 goroutines que impriman mensajes en paralelo.
- Enviar números a través de un canal y procesarlos en otra goroutine.

### Reto (30 min):

👉 Implementa un programa concurrente que calcule la suma de cuadrados de una lista de números, distribuyendo el cálculo entre varias goroutines.
