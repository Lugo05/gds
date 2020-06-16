package document_test

import (
	"github.com/Lugo05/gds/document"
)

func SetUp() *document.Document {
	doc := document.New()
	doc.Add("Cadena0", "Ejemplo0").Add("Entero0", 0).Add("Flotante0", 0.0).Add("Complejo0", 0+0i)
	doc.Add("SinSigno0", uint(0)).Add("Boolean0", true)

	sub1 := document.New()
	sub1.Add("Cadena1", "Ejemplo1").Add("Entero1", 1).Add("Flotante1", 1.1).Add("Complejo1", 1+1i)
	sub1.Add("SinSigno1", uint(1)).Add("Boolean1", true)

	sub2 := document.New()
	sub2.Add("Cadena2", "Ejemplo2").Add("Entero2", 2).Add("Flotante2", 2.2).Add("Complejo2", 2+2i)
	sub2.Add("SinSigno2", uint(2)).Add("Boolean2", false)

	sub1.Add("Documento1", sub2)
	doc.Add("Documento0", sub1)

	return doc
}

func TearDown() {

}
