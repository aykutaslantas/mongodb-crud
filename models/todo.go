package models
import "go.mongodb.org/mongo-driver/bson/primitive"
type Todo struct { // Diğer yerlerden erişebilmek için struct büyük harfle başlamalıdır
	Title string `json:"title,omitempty"`
	Id primitive.ObjectID `json:"id,omitempty"`
	Content string `json:"content,omitempty"`

}