# Golang

### Instalação (Linux Ubuntu 20.04)

* Baixar o instalador no [golang.org](https://golang.org/)

* Abrir o terminal (CTRL + ALT + T)

* Descompactar: `$ sudo cd ~/Downloads && tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz`

* Abrir o arquivo `.profile` no VS Code: `$ code ~/.profile`

* Adicionar no final do arquivo:
    
    ```
    # Golang
    export GOROOT=/usr/local/go
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin:$GOROOT/bin
    ```
  
* Aplicar as alterações: `$ source ~/.profile`

* Adicionar `export PATH=$PATH:/usr/local/go/bin` no ~/.zshrc

### Características

* Possui _Garbage Collector_ (limpar memória automaticamente)

### Go ENV

* **`GOPATH`**

    * **Definição**: diretório para armazenamentos de dados

    * **Tipos de dados armazenados**

        * Arquivos de dependências externas/projetos (`src/`)

        * Binários compilados pelo `$ go install` (`bin/`)

        * Pacotes pré-compilados (`pkg/`)

### Sintaxe

* **Função**
    
    * **Sintaxe**
      
        ```go
        func NomeDaFuncao() {
            // Corpo   
        }
        ```
      
    * **Exemplo**
      
        ```go
        func main() {
            println("Hello")    
        }
        ```

* **Declaração de variáveis**

    * **Sintaxe**

        ```go
        var <nomeDaVariavel> <tipoDaVariavel>
        ```
            
        **OU**    

        ```go
        // Inferência de tipo
        nomeDaVariavel := <valor>
        ```
                        
    * **Exemplo**
        
        ```go
        var nome string
        nome = "Wesley"
        ```
        
        ```go
        nome := "Wesley"
        nome = "Gabriel"
        ```

* **Declaração de constantes**
    
    * **Sintaxe**
      
        ```go
        const <nome> <tipo> = <valor>
        ```
        
        * **Bloco de constantes**

            ```go
            const (
                <nome> <tipo> = <valor>
                <nome> <tipo> = <valor>
                <nome> <tipo> = <valor>
            ) 
            ```
  
    * **Exemplo**
        
        ```go
        const xpto int = 1
        const nome string = "wes"
        ```

        ```go
        const (
          a string = "aaa"
          b string = "bbb"
          c string = "ccc"
        )
        ```

* **Tipos do Golang**
    
    Tipo | Exemplo
    :---:|:-------:
    int | `10` 
    string | `"Hello"`
    float64 | `10.22`
    bool | false
    int32 | `'w'`

        
    ```go
    fmt.Printf("%T \n", 10) // int
    fmt.Printf("%T \n", "Hello") // string
    fmt.Printf("%T \n", 10.44) // float64
    fmt.Printf("%T \n", false) // bool
    fmt.Printf("%T \n", 'w') // int32
    fmt.Printf(
        "%T \n", 
        `Wesley
            lindo
        `
    ) // string
    ```

* **Tratamento de erros**

    > `err` é a nomenclatura padrão para erro (convenção)
    ```go
    result, err := http.Get("https://www.google.com")
    
    if err != nil {
        panic("Erro!")
    }
    
    fmt.Println(result.Body)
    ```
  
* **Ignorar um erro**
    
    > `_` (**blank identifier**) = valor (erro) recebido é eliminado da memória
    ```go
    result, _ := http.Get("https://www.google.com")

    if error != nil {
        panic("Erro!")
    }
    
    fmt.Println(result.Body)
    ```

* **Obter a referência de uma variável no memória**

    > `&` = endereçamento de memória (em hexadecimal)

    ```go
    x := 10
	fmt.Println(&x)
    ```

* **Ponteiro (`*`)**

    * **Definição** = valor registrado no endereçamento de memória, ou seja, o valor de uma variável

    * **Exemplos**

        ```go
        x := 10
        y := &x

        fmt.Println(*y)
        ```

        ```go
        func main() {
            x := 10
            result := getPointerValue(&x)
            fmt.Println(result)
        }

        func getPointerValue(a *int) int {
            return *a
        }
        ```
        * `a *int` (parâmetro da função) = endereçamento de memória no formato inteiro

        * `int` (retorno) = inteiro

        * `return *a` = retornar o valor dentro do ponteiro (`a *int`)

* **Função**

    * **Função com 2 ou mais tipos de retorno**

        ```go
        func nome(a string, b int) (string, int) {
            return a, b
        }
        ```

        ```go
        func nome(a string, b int) (string, int, error) {
            return a, b, nil
        }
        ```

    * **Função anônima**

        ```go
        anonymousFunction := func() int {
            return 1
        }

        fmt.Println(anonymousFunction())
        ```

* **Array (array com tamanho fixo)**

    * **Sintaxe**

        ```go
        var <nome> [<tamanho>]<tipo>
        ```

        ```go
        <nome> := [<tamanho>]<tipo>{valores}
        ```

    * **Exemplo**

        ```go
        var x [5]int
        x[0] = 1
        x[1] = 2
        x[2] = 3

        fmt.Println(x) // [1 2 3 0 0]
        ```

        **OU**

        ```go
        y := [5]int{1, 2, 3}
	    fmt.Println(y) // [1 2 3 0 0]
        ```

    * **Obter um valor através do index**

        * **Sintaxe**

            ```go
            <nome>[<index>]
            ```

        * **Exemplo**

            ```go
            y := [5]int{1, 2, 3, 4, 5}
            fmt.Println(y[1]) // 2
            ```

* **Slices (array com tamanho variável)**

    * **Sintaxe**

        ```go
        make([]<tipo>)
        ```

        **OU**

        ```go
        []<tipo>
        ```


    * **Exemplo**

        ```go
        slice := make([]int)
        slice[0] = 1
        ```

        ```go
        slice := make([]int, 5)
        fmt.Println(slice)

        slice[0] = 1
        slice[1] = 1
        slice[2] = 1
        slice[3] = 1
        slice[4] = 1

        slice = append(slice, 1, 2, 3, 4) // Cópia do array anterior com os novos elementos
        fmt.Println(slice)
        ```

        ```go
        sliceString := []string {
            "Hello",
            "World",
        }
        fmt.Println(sliceString[0])
        ```

* **Map ("chave-valor")**

    * **Sintaxe**

        ```go
        make(map[<tipo chave>]<tipo valor>)
        ```

    * **Exemplo**

        ```go
        mapExample := make(map[string]int)

        mapExample["wesley"] = 10
        mapExample["mariana"] = 12

        fmt.Println(mapExample["wesley"])
        ```

* **type**

    * **Sintaxe**

        ```go
        type <Nome> <tipo>
        ```
    
    * **Exemplo**

        ```go
        type CarName string
        type CarYear int

        func main() {
            var carName CarName
            carName = "Fusion"
            fmt.Println(carName)

            var carYear CarYear
            carYear = 2020
            fmt.Println(carYear)
        }
        ```

* **struct**

    > **struct** são como "classes"

    * **Sintaxe**

        ```go
        type <nome> struct {
            // atributos
        }
        ```

    * **Exemplo**

        ```go
        type Car struct {
            CarName string
            CarYear int
        }

        func main() {

            car := Car{
                CarName: "Fusion",
                CarYear: 2020,
            }
            fmt.Println(car.CarYear)
            fmt.Println(car.CarName)

        }
        ```

    * **Função de uma struct**

        * **Sintaxe**

            ```go
            func (<nome> <struct>) nome() {
                // Corpo
            }
            ```

        * **Exemplo**

            ```go
            func (car Car) drive() {
                fmt.Println(car.CarName, "andou!")
            }

            func main() {

                car := Car{
                    CarName: "Fusion",
                    CarYear: 2020,
                }

                car.drive() // Fusion andou!

            }
            ```

            ```go
            func (car *Car) changeCarName()  {
                car.CarName = "Fusion2"
                fmt.Println(car.CarName)
            }

            func main() {

                car := Car{
                    CarName: "Fusion",
                    CarYear: 2020,
                }

                car.changeCarName() // Fusion2

            }
            ```

    * **Converter uma `struct` em JSON**

        * **Sintaxe**

            ```go
            json.Marshal(<struct>)
            ``` 
            
            > A função `Marshal()` transforma `struct` em slice de byte (`[]byte`)

        * **Exemplo**

            ```go
            type Car struct {
                CarName string
                CarYear int
            }

            func main() {
                result, _ := json.Marshal(car)
                fmt.Println(string(result)) // {"CarName":"Fusion","CarYear":2020}
            }
            ```

    * **Converter um JSON em `struct`**

        * **Sintaxe**

            ```go
            json.Unmarshal(<slice de byte>)
            ``` 

        * **Exemplo**

            ```go
            type Car struct {
                CarName string
                CarYear int
            }

            func main() {
            carJson := []byte(`{"car": "BMW", "year": 2020}`)

            var car Car
            json.Unmarshal(carJson, &car)
            fmt.Println(car) // {BMW 2020}
            }
            ```

    * **Adicionar tags aos atributos de uma `struct`**

        * **Sintaxe**

            ```go
            type Car struct {
                CarName string `json:"<nome do campo>"`
                CarYear int    `json:"<nome do campo>"`
            }
            ```

            > **OBS**: colocar um `-` no "nome do campo" para que esse campo seja ignorado na conversão para JSON.

        * **Exemplo**

            ```go
            type Car struct {
                CarName string `json:"car"`
                CarYear int    `json:"year"`
            }


            func main() {

                car := Car{
                    CarName: "Fusion",
                    CarYear: 2020,
                }

                result, _ := json.Marshal(car)
                fmt.Println(string(result)) // {"car":"Fusion","year":2020}

            }
            ```

            ```go
            type Car struct {
                CarName string `json:"car"`
                CarYear int    `json:"-"`
            }


            func main() {

                car := Car{
                    CarName: "Fusion",
                    CarYear: 2020,
                }

                result, _ := json.Marshal(car)
                fmt.Println(string(result)) // {"car":"Fusion"}

            }
            ```

* **Interfaces**

    * **Sintaxe**

        ```go
        type <nome> interface {
            // funções
            <nome>(<parâmetros>) <retorno>
        }
        ```

    * **Exemplo**

        ```go
        type vehicle interface {
            start() string
        }
        ```

        * **Implementação de uma `interface` vehicle**

            ```go
            // Como struct Car implementou a função start(), logo Car é um vehicle
            func (car Car) start() string {
                return "Iniciou"
            }
            ```

        ```go
        func acceptOnlyVehicleImplementation(car vehicle) {
            fmt.Println("É um veículo")
        }

        func main() {

            car := Car{
                CarName: "Fusion",
                CarYear: 2020,
            }

            acceptOnlyVehicleImplementation(car) // É um veículo

        }
        ```

### Comandos

* **Executar um arquivo `.go`**
  
    * **Sintaxe**: `$ go run <arquivo>`
    
    * **Exemplo**: `$ go run teste.go`
    
* **Criar um arquivo executável**

    * **Sintaxe**: `$ go build <arquivo>`
    
    * **Exemplo**: `$ go build teste.go`

* **Criar um arquivo executável para um OS específico**

    * **Sintaxe**: `$ GOOS=windows go build <arquivo>`

    * **Exemplo**: `$ GOOS=windows go build ./teste.go`