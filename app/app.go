package app

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"example.com/StructProject/note"
	"example.com/StructProject/storage"
	"example.com/StructProject/user"
)

var loggedUser string

// Funcionamento completo do App de notas
func App() {
	var choice int

	for {
		_appHome()
		fmt.Print("Escolha: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			_userLoggon()

		case 2:
			_createUser()

		case 3:
			fmt.Println("Encerrando aplicativo...")
			return
		default:
			fmt.Println("Opção inválida")
		}
	}
}

func _appHome() {
	fmt.Println("*************************************")
	fmt.Println("BEM VINDO AO GERADOR DE NOTAS DE JV!")
	fmt.Println("*************************************")
	fmt.Println("Digite uma opção: ")
	fmt.Println("1. Logar")
	fmt.Println("2. Criar conta")
	fmt.Println("3. Encerrar aplicativo")
}

func __userMenu() {
	var choice int
	for {
		fmt.Println("*************************************")
		fmt.Printf("Usuário: %s\n", loggedUser)
		fmt.Println("1. Criar nota")
		fmt.Println("2. Ver notas")
		fmt.Println("3. Logout")
		fmt.Println("Escolhar: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			_createNote()
		case 2:
			_notesViewer()
		case 3:
			fmt.Println("LOGOUT REALIZADO!")
			loggedUser = ""
			return
		default:
			fmt.Println("Opção inválida")
		}
	}
}

func _createUser() {
	var login string
	var password string

	fmt.Println("Digite seu login:")
	fmt.Scan(&login)

	fmt.Println("Digite sua senha:")
	fmt.Scan(&password)

	u := user.New(login, password)
	if err := os.MkdirAll(filepath.Join("data", "users"), 0o755); err != nil {
		fmt.Println("Erro:", err)
		return
	}
	userNotesDir := filepath.Join("data", "notes", login)
	if err := os.MkdirAll(userNotesDir, 0o755); err != nil {
		fmt.Println("Error", err)
		return
	}
	uJson := filepath.Join("data", "users", login+".json")

	if _, err := os.Stat(uJson); err == nil {
		fmt.Println("Usuário já existe")
		return
	} else if !os.IsNotExist(err) {
		fmt.Println("Error ao acessar o arquivo:", err)
		return
	}

	if err := storage.SaveJSON(uJson, u); err != nil {
		fmt.Println("Error ao salvar usuário:", err)
		return
	}

	fmt.Println("Usuário cadastrado com sucesso!")
}

func _userLoggon() {
	var psw string

	fmt.Println("Digite seu login:")
	fmt.Scan(&loggedUser)
	fmt.Println("Digite sua senha:")
	fmt.Scan(&psw)

	uJson := filepath.Join("data/users/" + loggedUser + ".json")

	u, err := storage.LoadJSON[user.User](uJson)

	fmt.Println("Validando...")

	if err != nil {
		fmt.Println(err)
	}

	if user.CheckPassword(psw, u.Password) {
		fmt.Println("LOGIN FEITO!")
		__userMenu()
	} else {
		fmt.Println("SENHA INCORRETA")
	}
}

func _createNote() {
	var noteTitle string
	var noteContent string

	// recebe de argumento a souce que iremos loggedUser
	// no caso abaixo os.Stdin que é o input padrão do usuario na CLI
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("---- CRIAÇÃO DE NOTA ----")

	fmt.Println("Digite o titulo da sua nota: ")
	fmt.Scan(&noteTitle)

	fmt.Println("Digite o texto da sua nota: ")

	// Aqui você vai atribuir um caracter para ele parar de ler a entrada
	// no caso abaixo \n
	// deve-se usar single quotes
	// ''
	noteContent, _ = reader.ReadString('\n')

	// P
	noteContent = strings.TrimSpace(noteContent)

	timestamp := time.Now().Unix()

	noteFileName := fmt.Sprintf("%d_%s.json", timestamp, noteTitle)
	noteFileName = strings.ReplaceAll(noteFileName, " ", "_")
	noteFileName = strings.ToLower(noteFileName)
	notePath := filepath.Join("data", "notes", loggedUser, noteFileName)

	userNote := note.NewNote(loggedUser, noteTitle, noteContent)

	storage.SaveJSON(notePath, userNote)
}

func _notesViewer() {
	var choice int
	notesPath := filepath.Join("data", "notes", loggedUser)
	files, err := os.ReadDir(notesPath)
	if err != nil {
		fmt.Println("Error:", err)
	}

	listFiles := []string{}
	fmt.Println("Escolha a nota a visualizar:")

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		path := filepath.Join(notesPath, file.Name())
		listFiles = append(listFiles, path)
		index := len(listFiles)
		fmt.Printf("[%d] %s\n", index, file.Name())
	}
	if len(listFiles) == 0 {
		fmt.Println("O usuário não possui notas")
		return
	}
	fmt.Println("Digite sua escolha: ")
	fmt.Scan(&choice)
	choice--
	if choice < 0 || choice >= len(listFiles) {
		fmt.Println("Escolha inválida")
		return
	}
	selectPath := listFiles[choice]

	n, err := storage.LoadJSON[note.Note](selectPath)
	if err != nil {
		fmt.Println("Erro ao carregar nota:", err)
		return
	}
	n.DisplayNote()
}
