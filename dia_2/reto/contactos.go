// Los ejecutables usan siempre package main;
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dfr99/golang-101/dia_2/reto/make_a_call" // importa tu librería
	"github.com/joho/godotenv"
)

func main() {
		// Cargar variables de entorno desde .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el .env: %v", err)
	}

	// Asignar variables a la librería
	make_a_call.AccountSID = os.Getenv("TWILIO_ACCOUNT_SID")
	make_a_call.AuthToken = os.Getenv("TWILIO_AUTH_TOKEN")
	make_a_call.FromNumber = os.Getenv("TWILIO_FROM_NUMBER")
	make_a_call.Url = os.Getenv("TWILIO_TWIML_URL")

	if make_a_call.AccountSID == "" || make_a_call.AuthToken == "" || make_a_call.FromNumber == "" {
		log.Fatal("❌ Faltan variables de Twilio en .env")
	}

	// Mapa de contactos: nombre -> teléfono
	contactos := make(map[string]string)

	reader := bufio.NewReader(os.Stdin)

	// Bucle infinito
	for {
		fmt.Println("\n--- Gestor de Contactos ---")
		fmt.Println("1. Agregar contacto")
		fmt.Println("2. Buscar contacto")
		fmt.Println("3. Listar contactos")
		fmt.Println("4. Llamar contacto")
		fmt.Println("5. Salir")
		fmt.Print("Elige una opción: ")

		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			fmt.Print("Nombre: ")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)

			fmt.Print("Teléfono: ")
			telefono, _ := reader.ReadString('\n')
			telefono = strings.TrimSpace(telefono)

			contactos[nombre] = telefono
			fmt.Println("✅ Contacto agregado.")

		case "2":
			fmt.Print("Nombre a buscar: ")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)

			if telefono, ok := contactos[nombre]; ok {
				fmt.Printf("📞 %s: %s\n", nombre, telefono)
			} else {
				fmt.Println("❌ Contacto no encontrado.")
			}

		case "3":
			fmt.Println("\n📋 Lista de contactos:")
			if len(contactos) == 0 {
				fmt.Println("No hay contactos guardados.")
			}
			for nombre, telefono := range contactos {
				fmt.Printf("- %s: %s\n", nombre, telefono)
			}

		case "4":
			fmt.Print("Nombre del contacto a llamar: ")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)

			if numero, ok := contactos[nombre]; ok {
				make_a_call.Llamar(numero)
			} else {
				fmt.Println("❌ Contacto no encontrado.")
			}

		case "5":
			fmt.Println("👋 Saliendo...")
			return

		default:
			fmt.Println("Opción no válida.")
		}
	}
}
