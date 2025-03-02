# **Go Project Skeleton** 🚀  

A **scalable, modular, and reusable** Go project template designed to serve as a foundation for real-world applications. This skeleton provides essential functionalities such as **database connections, RabbitMQ, Kafka integration, logging, and configuration management**—allowing you to focus on **business logic** instead of setting up boilerplate code.  

## **📌 Features**  
✅ **Modular and reusable structure** – Well-organized project layout.  
✅ **Configuration loader** – YAML-based configuration in `internal/config`.  
✅ **Database connectivity** – Supports **PostgreSQL, Oracle**, and other databases.  
✅ **RabbitMQ & Kafka integration** – Seamless messaging setup.  
✅ **Centralized logging** – Structured logging for better debugging.  
✅ **Reusable packages** – `pkg/` directory contains shared libraries usable across multiple projects.  
✅ **Easily extendable** – Just add your custom business logic.  

---

## **📂 Project Structure**  

```
GO_PROJECT_SKELETON/
├── cmd/                         # Main entry point for different applications
│   ├── config_check_app/        # Check config loading functionality
│   │   └── main.go
│   ├── database_check_app/      # Check database connection functionality
│   │   └── main.go
│   ├── kafka_check_app/         # Check Kafka consumer/producer functionality
│   │   └── main.go
│   └── rabbitmq_check_app/      # Check RabbitMQ functionality
│       └── main.go
├── internal/config/             # Configuration files and loader logic
│   ├── config.go
│   ├── dev.yml                  # Development environment configuration
│   ├── prod.yml                 # Production environment configuration
│   └── secrets.yml              # Secrets (passwords, tokens) for different environments
├── pkg/                         # Library packages for reuse
│   ├── database/                # Database connection logic
│   │   └── database.go
│   ├── kafka/                   # Kafka producer/consumer logic
│   │   └── kafka.go
│   ├── rabbitmq/                # RabbitMQ publisher/consumer logic
│   │   └── rabbitmq.go
│   └── utils/                   # Utility functions (e.g., logging, conversions)
│       ├── conversion.go
│       └── logger.go
├── test/                        # Unit tests
│   └── testex_test.go           # Example test file
├── go.mod                       # Go module file
├── go.sum                       # Dependency management file
└── README.md                    # Project documentation
```

---

## **🛠 Installation & Setup**  

### **1️⃣ Clone the Repository**
```bash
git clone https://github.com/tombro27/go_project_skeleton.git
cd go_project_skeleton
```

### **2️⃣ Install Dependencies**
```bash
go mod tidy
```

### **3️⃣ Setup Configuration**
Modify `internal/config/dev.yml`, `internal/config/prod.yml`, and `internal/config/secrets.yml` based on your environment.

### **4️⃣ Run the Application**
```bash
go run cmd/main.go
```

---

## **⚙️ Configuration**
This project follows a **YAML-based configuration** for managing credentials, service URLs, and environment-specific settings.  

### **Example `internal/config/dev.yml`**
```yaml
database:
  pg1:
    driver: "postgres"
    host: "localhost"
    port: "5432"
    username: "user1"
    dbname: "testdb"
rabbitmq:
  ser1:
    host: "localhost"
    port: "5672"
    username: "guest"
```

### **Example `internal/config/secrets.yml` (Not to be committed)**
```yaml
database:
  pg1:
    password: "supersecurepassword"
rabbitmq:
  ser1:
    password: "rabbitmqpassword"
```

---

## **📡 Database Support**
Supports **PostgreSQL and Oracle**. The database connection function **automatically selects the correct database based on the config**.

#### **Get Database Connection**
```go
db, err := database.GetDB("pg1", config, secrets)
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

---

## **📩 Message Queues**
### **RabbitMQ**
```go
conn, err := rabbitmq.GetRabbitMQConnection(cfg.RabbitMQ[key])
if err != nil {
	log.Fatal(err)
}
```

### **Kafka**
```go
producer, _ := kafka.GetKafkaProducer(cfg.Kafka["serv1"].Broker)
if err != nil {
	log.Fatal(err)
}
kafka.PublishMessage(producer, "test-topic", "key1", "Valeu1")
```

---

## **📜 Logging**
Structured logging using **Zap** or **Logrus**.
```go
logger := logger.NewLogger()
logger.Info("Application started successfully")
```

---

## **🧪 Testing**
Test files are located in the `/test` directory.
```bash
go test ./...
```

---

## **🆕 Future Enhancements & Suggestions** 🚀
Some additional features that can enhance this skeleton further:  
✅ **gRPC support** – Implement microservices with gRPC.  
✅ **REST API boilerplate** – Add basic CRUD APIs using `gin-gonic` or `echo`.  
✅ **Authentication module** – JWT-based authentication.  
✅ **Monitoring & Metrics** – Integrate **Prometheus & Grafana** for observability.  
✅ **Docker & Kubernetes** – Add `Dockerfile` and `k8s` configs for containerization.  
✅ **Rate Limiting & Caching** – Use **Redis** for caching.  
✅ **CI/CD pipeline** – Automate testing & deployments via **GitHub Actions**.  

---

## **📜 License**
This project is open-source and available under the **MIT License**.

---

This **Go project skeleton** helps you get started quickly and build scalable, real-world applications.  
Feel free to **contribute** and make it even better! 🚀