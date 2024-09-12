# Conceitos

go mod init heranca

go env para saber as varaiveis de ambiente

GOROOT=C:\Program Files\Go -- essa é a pasta do go (nao é aconselhado mudar)
GOPATH=C:\Users\razie\go -- meu workspace

pasta go
    - bin - guarda os executaveis compilado que criei (o load da aplicação)
    - pkg - tem as libs padroes compiladas
    - src - onde salva os repositorios do seu codigo - esse e o lugar ideal para colocar os projetos    
        - voce cria pacotes
        - é como um mini repositorio que fica hospedado em algum serviço por exemplo no github

# modulos

sao como uma bilblioteca de musica por exemplo vocer organiza pelas pastas o genero de cada musica

- o codigo sempre sera iniciado pela funcao main!
- go tem um linter embutido!
    - go doc
    - go run
    - go build
    - go get -u github.com/go-sql-driver/mysql (instala pacote)
- sempre que declarar uma variavel é obrigado a usar
- := forma mais comum de atribuir variavel
- nao precis apor ponto e virgula
- usar _ quando nao for usar o driver na hora eu digo no import olhar import primrio
     := declara e atribui // mais usada
     = declara quando ja existe
     	var e, f bool = true, false mu ltipla
        const ()
        var ()
        const 
        var
         reflect.TypeOf(g)
         float64 é padrao
         fortemente tipada
         so string com ""
         