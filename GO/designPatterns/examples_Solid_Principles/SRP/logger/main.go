package main

//FOLLOWING THE PRINCIPLE OF SRP

// ConsoleLogger and FileLogger have single responsibilities:
// logging to the console
type ConsoleLogger struct {
}

func (l *ConsoleLogger) Log(message string) {
	// Log to the console
}

// logging to a file
type FileLogger struct {
}

func (l *FileLogger) Log(message string) {
	// Log to a file
}

func main() {

}
