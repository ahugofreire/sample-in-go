package main

type Client struct {
	Name      string
	Age       int
	Status    bool
	Addresses []Address
}

type Address struct {
	Number    int
	StretName string
	City      string
}

func (c Client) setStatus(status bool) {
	c.Status = status
}
func (c Client) GetName() string {
	return c.Name
}

type Show interface {
	GetName() string
}

type Company struct {
	Name string
	CNPJ int
}

func (cp Company) GetName() string {
	return cp.Name
}

func ShowName(coisa Show) {
	coisa.GetName()
}

func main() {
	john := Client{
		Name:   "john",
		Age:    30,
		Status: true,
		Addresses: []Address{
			{
				10,
				"Rua 15",
				"Lomdon",
			},
			{
				25,
				"Rua liv",
				"Lomdon",
			},
		},
	}

	john.setStatus(false)
	john.GetName()

	unbrella := Company{
		Name: "Unbrella company",
		CNPJ: 1203000103,
	}

	//Ambas as struct estão imprementado a interface, basta ter um metodo,
	//Com a mesma assinatura, não é possivel ser propriedades na interface
	ShowName(unbrella)
	ShowName(john)
}
