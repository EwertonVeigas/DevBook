package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

	if erro := Usuario.formatar(etapa); erro != nil {
		return erro
	}
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

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O email é invalido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A Senha é Obrigatorio e nao pode estar em branco")
	}

	return nil
}
func (Usuario *Usuario) formatar(etapa string) error {
	Usuario.Nome = strings.TrimSpace(Usuario.Nome)
	Usuario.Nick = strings.TrimSpace(Usuario.Nick)
	Usuario.Email = strings.TrimSpace(Usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(Usuario.Senha)
		if erro != nil {
			return erro
		}
		Usuario.Senha = string(senhaComHash)
	}

	return nil
}
