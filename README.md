# Sobre

Essa implementação tem como objetivo fixar conhecimento adquirido durante o módulo de Comunicação via gRPC da Formação Full Cycle by Code.Edu.

Dessa forma, serão apresentados exemplos de comunicação gRPC (Client/Server, Server Stream, Client Stream e Stream bi-direcioinal) em Goland.

---

# Preparação de ambiente

- https://developers.google.com/protocol-buffers/docs/reference/go-generated

- https://grpc.io/docs/languages/go/quickstart/

---

# Geração dos pacotes Protocol Buffers e gRPC

```

protoc --proto_path=proto proto/*.proto --go_out=pb
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

```

---

# Simulação de execução

## Server - Terminal 1

```

go run cmd/server/server.go

```


Output:
```

---------------------------- Server: Client Stream -----------------------------
Inserindo Usuário 1
Inserindo Usuário 2
Inserindo Usuário 3
Inserindo Usuário 4
Inserindo Usuário 5

```


## Client - Terminal 2

```

go run cmd/client/client.go

```


Output:
```

---------------------------- Client/Server -----------------------------
id:"1234567" name:"Rodrigo Aires" email:"roaires@gmail.com" 


---------------------------- Server Stream -----------------------------
Status: Instância de um novo objeto User
id:"null" name:"null" email:"null" 


Status: Objeto user com valores preenchidos
id:"0" name:"Rodrigo Aires" email:"roaires@gmail.com" 


Status: Simulando registro inserido
id:"1234" name:"Rodrigo Aires" email:"roaires@gmail.com" 


Status: Simulando retorno do registro inserido
id:"1234" name:"Rodrigo Aires" email:"roaires@gmail.com" 




---------------------------- Client Stream -----------------------------
Enviando...  Usuário 1
Enviando...  Usuário 2
Enviando...  Usuário 3
Enviando...  Usuário 4
Enviando...  Usuário 5

Lista de usuários - Retorno do server:
user:<id:"1" name:"Usu\303\241rio 1" email:"user1@user.com" > user:<id:"2" name:"Usu\303\241rio 2" email:"user2@user.com" > user:<id:"3" name:"Usu\303\241rio 3" email:"user3@user.com" > user:<id:"4" name:"Usu\303\241rio 4" email:"user4@user.com" > user:<id:"5" name:"Usu\303\241rio 5" email:"user5@user.com" > 


---------------------------- Stream bi-direcional ----------------------
Enviando...  Usuário 1
Recebendo user Usuário 1 com status: Added
 Enviando...  Usuário 2
Recebendo user Usuário 2 com status: Added
 Enviando...  Usuário 3
Recebendo user Usuário 3 com status: Added
 Enviando...  Usuário 4
Recebendo user Usuário 4 com status: Added
 Enviando...  Usuário 5
Recebendo user Usuário 5 com status: Added


```

