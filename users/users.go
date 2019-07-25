package users

/*User type to store the users */
type User struct {
	ID           int     `json:"id"`
	Name         string  `json:"name,omitempty"`
	TicketSize   float32 `json:"ticketSize,omitempty"`
	TotalAccount float32 `json:"totalAccount,omitempty"`
}
