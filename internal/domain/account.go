package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	ApiKey  string  `json:"api_key"`
	Balance float64 `json:"balance"`

	mu        sync.RWMutex //concorrencia, problema de condicao de corrida, bloquear o acesso a variavel balance
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

func GenerateApiKey() string {
	b := make([]byte, 16)
	rand.Read(b)

	return hex.EncodeToString(b)
}

func NewAccount(name, email string) *Account {
	account := &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		ApiKey:    GenerateApiKey(),
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}

func (a *Account) AddBalance(amount float64) {
	//bloqueia o acesso a variavel balance
	a.mu.Lock()

	//espera a funcao terminar e libera o acesso a variavel balance
	defer a.mu.Unlock()

	a.Balance += amount
	a.UpdatedAt = time.Now()
}
