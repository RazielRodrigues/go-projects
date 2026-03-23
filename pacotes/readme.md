# PACOTES

from: https://go.dev/doc/tutorial/create-module

- quando usamos go mod init example.com/hello estamos adicionando trackeamento de dependencia no nosso codigo
- como o meu modulo nao foi publicado de verdade no github ou em outro lugar eu preciso rodar
    - go mod edit -replace example.com/greetings=../greetings
        - isso vai fazer apontar para o meu codigo local
    - depois rodar go mod tidy para sincronizar a mudanca
- depois de rodar o go build
    - eu posso rodar
        - assim vai me mostrar onde os meus arquivos go ficam salvo 
        - go list -f '{{.Target}}'
            - no meu caso: /Users/raziel.rodrigues/go/bin/hello
        - dai eu posso exportar a variavel de ambiente export PATH=$PATH:/Users/raziel.rodrigues/go/bin/hello
- go env -w GOBIN=/path/to/your/bin
    - com esse comando eu consigo mudar esse caminho default das aplicacoes go
- depois de atualizar o go path eu posso rodar "go install" e ai o meu binario vai ser um executavel dentro do sistema operacional
    - hello




- GOROOT: fica o codigo guardado do go em si o codigo do compilador
- GOPATH: fica guardado o codigo no usuario
- package: subpastas dentro do projeto ou seja o package no inicio da pasta
- modules: sao um conjunto de pacotes e fica salvo dentro do go.mod
    contem:
        - module kv-store
        - go 1.25.1
        - require github.com/google/wire v0.7.0 // indirect
        - replace: voce faz o replace do pacote que precisa por exmeplo apontando pra uma versao do pacote local
        - exclude: quando voce nao quer que o pacote seja utilizado dentro do seu projeto
        - o que é o //indirect ? alguma das depencias precisa disso inderetamente ou voce nao ta usando isso no projeto mesmo depois de instalado
- go.sum
    - ele aponta os hashes do commit do go.mod
- como fazer as versoes do module?
    - cria uma nova pasta dentro do projeto v2 e joha o codigo la
    - via branches usando tag
    - ou com git tag
    - pseudo versao:
        - ele cria versoes baseado no hash do commit e adiciona no go.sum
- se usar o go mod tidy
    - limpa os codigos que nao sendo usando eig // indirect
- go get install vs go get -u -v
    - go get -u -v // vai baixar o codigo mais atualizado
    - go get install // vai baixar a versao que esta no projeto
