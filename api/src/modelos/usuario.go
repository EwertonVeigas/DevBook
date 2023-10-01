package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuario utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (Usuario *Usuario) Preparar(etapa string) error {
	if erro := Usuario.validar(etapa); erro != nil {
		return erro
	}

	Usuario.formatar()
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O Nome é Obrigatorio e nao pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O Nick é Obrigatorio e nao pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O Email é Obrigatorio e nao pode estar em branco")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A Senha é Obrigatorio e nao pode estar em branco")
	}

	return nil
}
func (Usuario *Usuario) formatar() {
	Usuario.Nome = strings.TrimSpace(Usuario.Nome)
	Usuario.Nick = strings.TrimSpace(Usuario.Nick)
	Usuario.Email = strings.TrimSpace(Usuario.Email)
}
