package app

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode"

	"example.com/StructProject/note"
	"example.com/StructProject/storage"
	"example.com/StructProject/todo"
	"example.com/StructProject/user"
	"golang.org/x/text/unicode/norm"
)

var loggedUser string

// Funcionamento completo do App de notas
func App() {
	var choice int

	for {
		appHome()
		fmt.Print("Escolha: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			userLoggon()

		case 2:
			createUser()

		case 3:
			fmt.Println("Encerrando aplicativo...")
			return
		default:
			fmt.Println("Opção inválida")
		}
	}
}

func appHome() {
	fmt.Println("*************************************")
	fmt.Println("BEM VINDO AO GERADOR DE NOTAS DE JV!")
	fmt.Println("*************************************")
	fmt.Println("Digite uma opção: ")
	fmt.Println("1. Logar")
	fmt.Println("2. Criar conta")
	fmt.Println("3. Encerrar aplicativo")
}

func userMenu() {
	var choice int
	for {
		fmt.Println("*************************************")
		fmt.Printf("Usuário: %s\n", loggedUser)
		fmt.Println("1. Criar nota")
		fmt.Println("2. Ver notas")
		fmt.Println("3. Criar tarefa")
		fmt.Println("4. Ver tarefas")
		fmt.Println("5. Logout")
		fmt.Println("Escolha: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			createNote()
		case 2:
			filesViewer("data", "notes")
		case 3:
			createTodo()
		case 4:
			filesViewer("data", "todos")
		case 5:
			return
		default:
			fmt.Println("Opção inválida")
		}
	}
}

func createUser() {
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

	userTodoDir := filepath.Join("data", "todos", login)
	if err := os.MkdirAll(userTodoDir, 0o755); err != nil {
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

func userLoggon() {
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
		return
	}

	if user.CheckPassword(u.HashPassword, psw) {
		fmt.Println("LOGIN FEITO!")
		userMenu()
	} else {
		fmt.Println("SENHA INCORRETA")
		return
	}
}

func createNote() {
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

	timestamp := time.Now().Unix()

	noteFileName := fmt.Sprintf("%d_%s.json", timestamp, sanitizeFilename(noteTitle))
	notePath := filepath.Join("data", "notes", loggedUser, noteFileName)

	userNote := note.NewNote(loggedUser, noteTitle, noteContent)

	storage.SaveJSON(notePath, userNote)
}

func filesViewer(dir, dir2 string) {
	var choice int
	notesText, notesTextError := "Escolha a nota para visualizar", "Error ao carregar notas:"
	todoText, todoTextoError := "Suas tarefas:", "Error ao carregar tarefas:"
	whichDir := dir2 == "notes"
	filesPath := filepath.Join(dir, dir2, loggedUser)
	files, err := os.ReadDir(filesPath)
	if err != nil {
		fmt.Println("Error:", err)
	}

	listFiles := []string{}

	if dir2 == "todos" {
		fmt.Println(todoText)
	} else {
		fmt.Println(notesText)
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		path := filepath.Join(filesPath, file.Name())
		listFiles = append(listFiles, path)
		index := len(listFiles)
		fmt.Printf("[%d] %s\n", index, file.Name())
	}
	if len(listFiles) == 0 {
		fmt.Println("O usuário não possui arquivos")
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

	if whichDir {
		n, err := storage.LoadJSON[note.Note](selectPath)
		if err != nil {
			fmt.Println(notesTextError, err)
			return
		}
		n.Display()
	} else {
		n, err := storage.LoadJSON[todo.Todo](selectPath)
		if err != nil {
			fmt.Println(todoTextoError, err)
			return
		}
		n.Display()

	}

	// teste
}

func createTodo() {
	var todoContent string
	fmt.Println("Digite a tarefa:")
	reader := bufio.NewReader(os.Stdin)
	todoContent, _ = reader.ReadString('\n')

	timestamp := time.Now().Unix()

	todoFileName := fmt.Sprintf("%d_%s.json", timestamp, sanitizeFilename(todoContent))

	todoPath := filepath.Join("data", "todos", loggedUser, todoFileName)

	userTodo, _ := todo.New(todoContent)

	storage.SaveJSON(todoPath, userTodo)
}

func removeAccents(s string) string {
	t := norm.NFD.String(s)
	result := make([]rune, 0, len(t))
	for _, r := range t {
		if unicode.Is(unicode.Mn, r) { // Mn = Mark Nonspacing (acentos)
			continue
		}
		result = append(result, r)
	}
	return string(result)
}

func sanitizeFilename(name string) string {
	name = strings.TrimSpace(name)
	name = removeAccents(name)

	// troca espaços por underscore
	name = strings.ReplaceAll(name, " ", "_")

	// remove qualquer caractere que não seja letra, número, hífen ou underscore
	re := regexp.MustCompile(`[^a-zA-Z0-9._-]+`)
	name = re.ReplaceAllString(name, "")

	name = strings.ToLower(name)
	return name
}
