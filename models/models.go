package models

type SecretRDSJson struct {
	Username            string `json:"username"` //Alt izq + 96
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

type SignUp struct {
	UserEmail string `json:"UserEmail"` //Alt izq + 96
	UserUUID  string `json:"UserUUID"`
}

//Para la parte de Auth

type TokenJSON struct {
	Sub       string
	Event_Id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

//Estructuras para category

type Category struct {
	CategID   int    `json:"categID"`
	CategName string `json:"categName"`
	CategPath string `json:"categPath"`
}

//Estructura para products

type Product struct {
	ProdId         int     `json:"prodID"`
	ProdTitle      string  `json:"prodTitle"`
	ProdDescrition string  `json:"prodDescription"`
	ProdCreateAt   string  `json:"prodCreateAt"`
	ProdUpdated    string  `json:"prodUpdated"`
	ProdPrice      float64 `json:"prodPrice,omitempty"`
	ProdStock      int     `json:"prodStock"`
	ProdCategId    int     `json:"prodCategId"`
	ProdPath       string  `json:"prodPath"`
	ProdSearch     string  `json:"search,omitempty"`
	ProdCategPath  string  `json:"categPath,omitempty"`
}

//Nuevo modelo para Select Productos

type ProductResp struct {
	TotalItems int       `json:"totalItems"`
	Data       []Product `json:"data"` //Nuestro slice (arreglo) de productos
}
