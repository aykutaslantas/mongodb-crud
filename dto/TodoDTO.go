package dto

type TodoDTO struct { // Diğer yerlerden erişebilmek için struct büyük harfle başlamalıdır
	Status bool `json:"status,omitempty"`
}