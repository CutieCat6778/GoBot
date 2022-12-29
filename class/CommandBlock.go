package class

type (
	CommandBlockStruct struct {
		id      string
		blocked bool
	}

	CommandBlockMap map[string]CommandBlockStruct

	CommandBlock struct {
		commandblock CommandBlockMap
	}
)

func NewCommandBlock() *CommandBlock {
	return &CommandBlock{make(CommandBlockMap)}
}

func (handler CommandBlock) GetCommandBlocks() CommandBlockMap {
	return handler.commandblock
}

func (handler CommandBlock) Get(id string) (*CommandBlockStruct, bool) {
	commandblock, found := handler.commandblock[id]

	return &commandblock, found
}

func (handler CommandBlock) Write(id string, blocked bool) (*CommandBlockStruct, bool) {
	commandblock, found := handler.commandblock[id]
	if !found {
		return &commandblock, found
	} else {
		commandblock.blocked = blocked
		return &commandblock, found
	}
}

func (handler CommandBlock) Register(id string, blocked bool) {
	commandblockstruct := CommandBlockStruct{id: id, blocked: blocked}
	handler.commandblock[id] = commandblockstruct
	if len(id) > 1 {
		handler.commandblock[id[:1]] = commandblockstruct
	}
}

func (commandblock CommandBlockStruct) GetStatus() bool {
	return commandblock.blocked
}
