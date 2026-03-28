# Go Task Manager 🚀

Bu proje, **Go (Golang)** öğrenme sürecinde geliştirilen, **Docker** ile konteynerize edilen ve **Clean Architecture** prensiplerini uygulayan bir **Görev Yönetimi (Task Management) REST API** projesidir.

Amacımız sadece kod yazmak değil; bir backend projesinin **planlama, geliştirme, test ve dağıtım** süreçlerini uçtan uca deneyimlemektir.

---

## 🛠 Kullanılan Teknolojiler

- **Dil:** Go (Golang)
- **Mimari:** Clean Architecture (Domain, Usecase, Repository, Delivery)
- **Veritabanı:** PostgreSQL
- **Konteynerizasyon:** Docker & Docker Compose
- **Kimlik Doğrulama:** JWT (JSON Web Token)
- **Konfigürasyon:** Viper
- **Versiyon Kontrol:** Git & GitHub

---

## 📂 Proje Yapısı ve Branch Stratejisi

Bu projede **Feature Branch Workflow** kullanılacaktır. Kodlar doğrudan `main` branch'ine atılmaz.

### Branch'ler:
1.  **`main`**: 🟢 Canlı ortam kodu (Production). Her zaman kararlı ve çalışır durumda olmalı.
2.  **`develop`**: 🟡 Geliştirme ortamı. Yeni feature'lar burada birleşir.
3.  **`feature/*`**: 🔵 Yeni özellikler (örn: `feature/auth`, `feature/task-crud`).

**Örnek Akış:**
`Issue #1` oluşturulur -> `feature/database-setup` dalı açılır -> Kodlanır -> `develop`'a Pull Request (PR) atılır.

---

## 🗺 Proje Yol Haritası (Roadmap)

Adım adım ilerleyerek geliştireceğimiz modüller:

- [x] **Adım 1:** GitHub Proje Kurulumu ve Planlama (Şu an buradayız)
- [ ] **Adım 2:** Go Proje İskeleti ve Clean Architecture Yapısı
- [ ] **Adım 3:** Config Yönetimi ve PostgreSQL Bağlantısı
- [ ] **Adım 4:** Kullanıcı Kayıt (Register) ve Giriş (Login)
- [ ] **Adım 5:** JWT Authentication ve Middleware
- [ ] **Adım 6:** Task CRUD İşlemleri
- [ ] **Adım 7:** Dockerfile ve Docker Compose Entegrasyonu
- [ ] **Adım 8:** GitHub Actions (CI) ve Dokümantasyon

---

## 🚀 Kurulum ve Çalıştırma

### Gereksinimler
- Docker
- Docker Compose

### Adımlar

1. **Projeyi Klonla:**
```bash
git clone https://github.com/hamidkrml/go-task.git
cd go-task
```

2. **Docker ile Başlat:**
```bash
docker-compose up --build
```

Uygulama `http://localhost:8080` adresinde çalışmaya başlayacak.

---

## 📡 API Kullanımı

### 1. Kullanıcı Kaydı
```bash
curl -X POST http://localhost:8080/register \
     -H "Content-Type: application/json" \
     -d '{"name": "Hamid", "email": "hamid@example.com", "password": "securepassword"}'
```

### 2. Giriş Yapma (Token Alma)
```bash
curl -X POST http://localhost:8080/login \
     -H "Content-Type: application/json" \
     -d '{"email": "hamid@example.com", "password": "securepassword"}'
```

**Yanıt:**
```json
{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."}
```

### 3. Görev Oluşturma (Token Gerekli)
```bash
curl -X POST http://localhost:8080/tasks \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer YOUR_TOKEN" \
     -d '{"title": "İlk Görevim", "description": "Bu bir test görevi"}'
```

### 4. Görevleri Listeleme
```bash
curl -X GET http://localhost:8080/tasks \
     -H "Authorization: Bearer YOUR_TOKEN"
```

### 5. Görev Durumu Güncelleme
```bash
curl -X PUT "http://localhost:8080/tasks/update?id=1" \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer YOUR_TOKEN" \
     -d '{"status": "done"}'
```

### 6. Görev Silme
```bash
curl -X DELETE "http://localhost:8080/tasks/delete?id=1" \
     -H "Authorization: Bearer YOUR_TOKEN"
```

---

## 🎓 Öğrenilen Konular

Bu proje geliştirme sürecinde şunlar öğrenildi:
- ✅ Go dilinde Clean Architecture uygulama
- ✅ PostgreSQL veritabanı entegrasyonu
- ✅ JWT tabanlı kimlik doğrulama
- ✅ Docker ve Docker Compose kullanımı
- ✅ Git ve GitHub workflow'ları
- ✅ REST API tasarımı ve güvenliği

---

## 📝 Lisans

Bu bir eğitim projesidir.
