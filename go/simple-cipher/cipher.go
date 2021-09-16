package cipher

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

func (*Cipher) Encode(string) string {
	return ""
}

func (*Cipher) Decode(string) string {

}

func NewCaesar() Cipher {}

func NewShift(in int) Cipher {}

func NewVigenere(str string) Cipher {}
