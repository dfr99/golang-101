package makeacall

import (
	"fmt"
	"log"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// Configuración de Twilio (puedes pasarla desde contacto.go si quieres)
var AccountSID string
var AuthToken string
var FromNumber string
var Url string

func Llamar(numero string) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: AccountSID,
		Password: AuthToken,
	})

	params := &openapi.CreateCallParams{}
	params.SetTo(numero)
	params.SetFrom(FromNumber)
	params.SetUrl(Url) // Demo Twilio

	resp, err := client.Api.CreateCall(params)
	if err != nil {
		errMsg := err.Error()
		switch {
		case contains(errMsg, "not verified") || contains(errMsg, "unverified"):
			log.Printf("❌ Error: El número %s no está verificado en Twilio.\n", numero)
		case contains(errMsg, "URL") && contains(errMsg, "invalid"):
			log.Printf("❌ Error: La URL proporcionada es incorrecta o no accesible.\n")
		case contains(errMsg, "not a valid phone number") || contains(errMsg, "not registered"):
			log.Printf("❌ Error: El número %s no está registrado o es inválido.\n", numero)
		case contains(errMsg, "Authenticate") || contains(errMsg, "401") || contains(errMsg, "AccountSid or AuthToken"):
			log.Printf("❌ Error: Credenciales de Twilio incorrectas.\n")
		default:
			log.Printf("❌ Error al hacer la llamada: %s\n", err)
		}
		return
	}
	fmt.Printf("✅ Llamada iniciada a %s, SID: %s\n", numero, *resp.Sid)
}

// contains es un helper para comprobar si un substring está en un string
func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && (stringIndex(s, substr) >= 0)
}

// stringIndex devuelve el índice de substr en s, o -1 si no está
func stringIndex(s, substr string) int {
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
