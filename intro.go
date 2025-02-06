package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos-file.json")
	storage.Load(&todos)
	CmdFlags := NewCmdFlags()
	CmdFlags.Execute(&todos)
	//todos.add("Compra Leite de soja")
	//todos.add("comprar Pão")
	//todos.add("Construir Projecto de em golang")
	//todos.toggle(0)
	//todos.toggle(1)
	//todos.print()
	storage.Save(todos)
}
