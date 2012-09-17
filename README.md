Next Steps:
 - Simple echo server via telnet
 - Change rooms and get messages
 - Change rooms with other players
 - Rabid Squirrel attacking

	// handler should create a "PlayerConnection"
	// It constructs from a buffered reader and writer and a username/password
	// if I call .SendFrom(sender, message) it should filter out
	// if I send it commands, it processes them and calls other objects
	// To get a username/password, it should use a LoginContext to handle back and forth retries
	// username/password stored in some key/value system
