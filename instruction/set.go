package instruction

func InstructionByName(name string) Instruction {
	instructionSet := map[string]Instruction{
		"join":    Join{},
		"message": Message{},
	}

	return instructionSet[name]
}
