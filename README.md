# 🔥 Pokedex CLI

<div align="center">
  
  ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
  ![Pokemon](https://img.shields.io/badge/Pokemon-API-FFCB05?style=for-the-badge&logo=pokemon&logoColor=blue)
  ![CLI](https://img.shields.io/badge/CLI-Tool-green?style=for-the-badge)
  
  **A lightning-fast, interactive command-line Pokedex built with Go** ⚡
  
  *Explore the Pokemon world, catch creatures, and build your collection directly from your terminal*
  
</div>

---

## 🌟 Features

✨ **Interactive REPL Interface** - Smooth command-line experience  
🚀 **Intelligent Caching System** - Lightning-fast repeated queries  
🗺️ **Location Area Navigation** - Explore Pokemon world locations with pagination  
🎯 **Pokemon Catching System** - Probability-based catching mechanics  
📋 **Pokedex Management** - Inspect individual Pokemon and view your collection  
🔄 **Bidirectional Pagination** - Navigate forward and backward seamlessly  
🛡️ **Concurrent-Safe Operations** - Built with Go's goroutines and mutexes  
🏗️ **Clean Architecture** - Modular design with separation of concerns  

## 🚀 Quick Start

### Prerequisites

- Go 1.19+ installed on your system
- Internet connection (for PokeAPI access)

### Installation

```bash
# Clone the repository
git clone https://github.com/0xJeanmi/pokedex-cli.git
cd pokedex-cli

# Run the application
go run .
```

### Alternative: Build Binary

```bash
# Build executable
go build -o pokedex

# Run the binary
./pokedex
```

## 🎮 Usage

Once you start the application, you'll enter an interactive REPL mode:

```
Welcome to the Pokedex!
Pokedex > help
```

### Available Commands

| Command | Description | Example |
|---------|-------------|---------|
| `help` | Display all available commands | `Pokedex > help` |
| `map` | Display next 20 location areas | `Pokedex > map` |
| `mapb` | Display previous 20 location areas | `Pokedex > mapb` |
| `explore <area>` | Show Pokemon in a specific area | `Pokedex > explore canalave-city-area` |
| `catch <pokemon>` | Try to catch a Pokemon | `Pokedex > catch pikachu` |
| `inspect <pokemon>` | Show details of a caught Pokemon | `Pokedex > inspect pikachu` |
| `pokedex` | Display all caught Pokemon | `Pokedex > pokedex` |
| `exit` | Exit the Pokedex | `Pokedex > exit` |

### Example Session

```bash
Pokedex > map
🌏 canalave-city-area
🌏 iron-island-area
🌏 lake-valor-area
🌏 lake-verity-area
# ... more locations

Pokedex > explore canalave-city-area
🐾 tentacool
🐾 tentacruel
🐾 staryu
🐾 magikarp
# ... more Pokemon

Pokedex > catch tentacool
Throwing a Poke Ball at tentacool...
Congratulations! You caught: tentacool

Pokedex > inspect tentacool
Name: tentacool
Experience: 67

Pokedex > pokedex
Name: tentacool

Pokedex > exit
Closing the Pokedex... Goodbye! 👋
```

## 🏗️ Architecture

### Project Structure

```
📦 pokedex-cli/
├── 📁 internal/
│   ├── 🗃️ pokecache.go      # Intelligent caching system & Pokedex management
│   └── 🧪 cache_test.go     # Cache unit tests
├── ⚙️ main.go               # Application entry point & REPL
├── 🔧 utils.go              # HTTP utilities and commands setup
├── 🎯 commands.go           # Command implementations
└── 📚 README.md             # This file
```

### Core Components

#### 🗃️ **Cache & Pokedex System** (`internal/pokecache.go`)
- **Thread-safe operations** with RWMutex for both cache and Pokedex
- **Automatic cleanup** with configurable intervals
- **Memory efficient** with time-based expiration
- **Pokemon storage** with experience tracking
- **High performance** with concurrent read access

#### 🌐 **HTTP Client** (`utils.go`)
- **Smart caching integration** for repeated requests
- **Clean API abstraction** over PokeAPI
- **Automatic URL construction** with pagination support
- **Error handling** with user-friendly messages

#### 🎯 **Command System** (`commands.go`)
- **Modular command structure** for easy extension
- **State management** for pagination and Pokemon collection
- **Probability-based catching** using Pokemon base experience
- **Input validation** with helpful error messages

## 🔧 Technical Details

### Pokemon Catching Algorithm

The catching system uses a sophisticated probability calculation:

```go
// Probability range based on Pokemon's base experience
min := baseExperience / 3
max := baseExperience
probabilityToCatch := min + rand.Float64()*(max-min)

// Success if probability exceeds half the base experience
if probabilityToCatch >= baseExperience/2 {
    // Pokemon caught!
}
```

### Cache Implementation

The caching system uses a sophisticated approach for both API calls and Pokemon storage:

```go
type Cache struct {
    data map[string]cacheEntry
    lock sync.RWMutex
}

type Pokedex struct {
    pokemons map[string]Pokemon
    lock     sync.RWMutex
}
```

**Features:**
- 🔒 **Concurrent-safe** with read-write mutex
- ⏰ **Time-based expiration** (5-minute default)
- 🧹 **Automatic cleanup** via background goroutine
- 💾 **Byte-level storage** for memory efficiency
- 🎯 **Pokemon persistence** throughout session

### API Integration

Built on top of the excellent [PokeAPI](https://pokeapi.co/):
- **Location areas** endpoint for exploration
- **Pokemon species** endpoint for catching
- **Pagination support** with offset/limit parameters
- **JSON response parsing** with type-safe handling

## 🧪 Testing

Run the comprehensive test suite:

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run cache-specific tests
go test -v ./internal/
```

### Test Coverage

- ✅ Cache add/get operations
- ✅ Automatic cache expiration
- ✅ Concurrent access safety
- ✅ Error handling scenarios


### Coding Standards

- Follow Go conventions and `gofmt` formatting
- Write tests for new functionality
- Update documentation as needed
- Use descriptive commit messages


## 🙏 Acknowledgments

- **[PokeAPI](https://pokeapi.co/)** - Amazing free Pokemon API
- **[Go Team](https://golang.org/)** - For the fantastic programming language
- **Pokemon Company** - For creating this incredible universe

## 📬 Contact

Created with ❤️ by **Jeanmi** ([0xJeanmi](https://github.com/0xJeanmi))

---