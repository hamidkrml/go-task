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

## 🚀 Kurulum (İleride Eklenecek)

Proje tamamlandığında aşağıdaki komutlarla çalıştırılabilecek:

```bash
# Projeyi klonla
git clone https://github.com/kullaniciadi/go-task-manager.git

# Docker ile ayağa kaldır
docker-compose up --build
```
