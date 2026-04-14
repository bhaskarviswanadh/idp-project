# 🚀 Internal Developer Platform (IDP) CLI

A lightweight **Internal Developer Platform (IDP)** built using Go that automates the complete application deployment workflow, including containerization, dynamic Kubernetes resource generation, and service exposure with a **single command**.

---

## 📌 Overview

This project provides a CLI tool (`idp`) that automates the complete application deployment workflow:

* 📦 Builds Docker images
* ☸️ Deploys applications to Kubernetes (Minikube)
* ⚙️ Dynamically generates Kubernetes manifests
* 🌐 Exposes services and opens the application in a browser

---

## 🏛️ Architecture Diagram

<p align="center">
  <img width="2816" height="1536" alt="IDP_CLI_Architecture diagram" src="https://github.com/user-attachments/assets/a4994def-e1d2-4e30-89aa-deb98197a066" />
</p>

---

## ⚡ Features

* 🔹 One-command deployment (`idp deploy`)
* 🔹 Config-driven using `idp.yaml`
* 🔹 Dynamic Kubernetes YAML generation
* 🔹 Automatic Docker image build
* 🔹 Seamless Minikube integration
* 🔹 Auto-resolves service URL and opens browser
* 🔹 Clean CLI logs and structured output

---

## 🏗️ Project Structure

```id="g0m6vn"
idp-project/
├── cli/              # Go CLI tool
├── sample-app/       # Example Flask application
├── idp.yaml          # Configuration file
└── README.md
```

---

## ⚙️ Configuration (`idp.yaml`)

```yaml
app_name: idp-app
app_path: sample-app
image_name: idp-app
tag: latest
port: 5000
replicas: 2
```

---

## 🚀 How It Works

```id="y8r4od"
idp deploy
   ↓
Load configuration (idp.yaml)
   ↓
Build Docker image
   ↓
Load image into Minikube
   ↓
Generate Kubernetes manifests (dynamic)
   ↓
Deploy to Kubernetes cluster
   ↓
Fetch service URL
   ↓
Open application in browser
```

---

## 🧪 Usage

### 1️⃣ Build CLI

```bash
cd cli
go build -o idp.exe
```

---

### 2️⃣ Run Deployment

```bash
.\idp.exe deploy
```

---

### 3️⃣ Output Example

```id="a0n5zd"
[SUCCESS] Deployment pipeline completed successfully!
[INFO] Application URL: http://127.0.0.1:XXXXX
```

👉 Browser opens automatically 🎉

---

## 🛠️ Tech Stack

* **Go** — CLI development
* **Docker** — Containerization
* **Kubernetes (Minikube)** — Orchestration
* **YAML** — Configuration management

---

## 🧠 Key Concepts Demonstrated

* Infrastructure as Code (IaC)
* Kubernetes resource management
* CLI-based automation
* Dynamic configuration-driven systems
* DevOps pipeline design

---

## 🚧 Future Enhancements

* 🔹 Multi-application support
* 🔹 CI/CD integration (GitHub Actions)
* 🔹 Helm chart support
* 🔹 Cloud deployment (AWS EKS / GKE)
* 🔹 Web dashboard UI

---

## ⭐ Conclusion

This project demonstrates how developer experience can be improved by abstracting infrastructure complexity into a simple CLI tool — a foundational concept in modern DevOps.

---
