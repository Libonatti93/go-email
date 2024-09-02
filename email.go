package main

import (
    "fmt"
    "net"
    "regexp"
    "strings"
)

// Função para verificar a sintaxe do e-mail usando uma expressão regular
func isValidEmailSyntax(email string) bool {
    // Expressão regular para validar o formato do e-mail
    var re = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return re.MatchString(email)
}

// Função para verificar se o domínio do e-mail tem registros MX
func hasMXRecord(domain string) bool {
    mxRecords, err := net.LookupMX(domain)
    if err != nil || len(mxRecords) == 0 {
        return false
    }
    return true
}

// Função principal para verificar o e-mail
func verifyEmail(email string) {
    // Verifica a sintaxe do e-mail
    if !isValidEmailSyntax(email) {
        fmt.Printf("E-mail inválido: %s (Sintaxe incorreta)\n", email)
        return
    }

    // Separa o domínio do e-mail
    parts := strings.Split(email, "@")
    if len(parts) != 2 {
        fmt.Printf("E-mail inválido: %s (Formato incorreto)\n", email)
        return
    }
    domain := parts[1]

    // Verifica se o domínio tem registros MX
    if !hasMXRecord(domain) {
        fmt.Printf("E-mail inválido: %s (Domínio sem registros MX)\n", email)
        return
    }

    fmt.Printf("E-mail válido: %s\n", email)
}

func main() {
    // Lista de e-mails para verificação
    emails := []string{
        "teste@exemplo.com",
        "usuario@dominioinexistente.xyz",
        "email@site.com.br",
        "email_invalido@dominio",
        "email@dominiosemx.com",
    }

    // Verifica cada e-mail na lista
    for _, email := range emails {
        verifyEmail(email)
    }
}
