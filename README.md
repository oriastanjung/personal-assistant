
# **ORIVA Chatbot**

ORIVA is a personal assistant chatbot that uses gRPC with bidirectional streaming to communicate with users. It leverages OpenAI's GPT-4 to provide intelligent, context-aware responses in real time.

---

## **Features**
- Bidirectional gRPC streaming for real-time chat.
- Persistent conversation memory for contextual responses.
- Modular structure for scalability.
- Uses OpenAI's GPT-4 for natural language processing.

---

## **Prerequisites**
1. **Go**: Install Go version 1.20 or later.
2. **Protocol Buffers**: Install `protoc` compiler.
   - [Protocol Buffers Installation Guide](https://grpc.io/docs/protoc-installation/)
3. **gRPC Plugins for Go**: Install required Go plugins.
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```
4. **OpenAI API Key**: Get your API key from [OpenAI](https://platform.openai.com/).

---

## **Installation**

1. Clone the repository:
   ```bash
   git clone https://github.com/oriastanjung/personal-assistant.git
   cd personal-assistant
   ```

2. Install dependencies:
   ```bash
   go mod download
   go mod tidy
   ```

3. Generate gRPC code from `proto/chat.proto`:
   ```bash
   make proto
   ```

   If `make` is not installed, you can manually run:
   ```bash
   protoc --go_out=. --go-grpc_out=. proto/chat.proto
   ```

4. Create a `.env` file in the root directory and add your OpenAI API key:
   ```
   OPENAI_API_KEY=your-openai-api-key
   PORT=50027
   ```

---

## **Project Structure**

```plaintext
cmd/
├── main.go               # Entry point for the server
internal/
├── config/
│   ├── config.go         # Configuration loader
├── services/
│   ├── oriva_services.go # gRPC services implementation
├── usecase/
│   ├── oriva_usecase.go  # Chat logic and OpenAI integration
proto/
├── chat.proto            # gRPC definitions
```

---

## **Running the Project**

### 1. Run the Server
Start the gRPC server:
```bash
cd cmd
go run main.go
```

### 2. Run the Client
In a new terminal, start the gRPC client:
```bash
cd cmd/client
go run client.go
```

---

## **Usage**

- Type messages into the client terminal to interact with ORIVA.
- Exit the chat by typing `exit`.

---

## **Available Commands**

| Command            | Description                            |
|--------------------|----------------------------------------|
| `go mod download`  | Download Go module dependencies.       |
| `go mod tidy`      | Clean up and verify dependencies.      |
| `make proto`       | Generate gRPC code from proto file.    |
| `go run main.go`   | Start the gRPC server.                 |
| `go run client.go` | Start the gRPC client.                 |

---

## **Environment Variables**

| Variable        | Description              |
|-----------------|--------------------------|
| `OPENAI_API_KEY`| API key for OpenAI GPT-4 |
| `PORT`          | gRPC Server running PORT |

---

## **Contributing**
Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a feature branch: `git checkout -b feature-name`.
3. Commit your changes: `git commit -m "Add feature-name"`.
4. Push to your branch: `git push origin feature-name`.
5. Submit a pull request.

---

## **License**
This project is licensed under the [MIT License](LICENSE).

---

## **Contact**
For questions or feedback, feel free to reach out:
- Email: oriastan999@gmail.com
- GitHub: [oriastanjung](https://github.com/oriastanjung)
