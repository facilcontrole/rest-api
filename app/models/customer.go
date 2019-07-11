package models

/*

"id":"aaa5ef3a-e749-4a38-b7b3-8ab03ececeb7"
"name":"FacilControle"
 "items":{
	 "phone":"000000000"
 }

*/

type Customer struct {
	ID    string        `json:"id"`
	Name  string        `json:"name"`
	Items CustomerItems `json:"items"`
}

type CustomerItems struct {
	Phone string `json:"phone"`
}
