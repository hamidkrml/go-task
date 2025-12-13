# 📌 Örnek GitHub Issue'ları

Bu dosya, projenin geliştirilmesi sırasında açılması gereken örnek issue'ları ve formatlarını içerir.

## Issue Şablonu
**Başlık:** [Kategori] Kısa ve net başlık
**Açıklama:** Ne yapılacak? Neden gerekli?
**Kabul Kriterleri (Acceptance Criteria):**
- [ ] Şart 1
- [ ] Şart 2

---

## 1️⃣ Issue: Proje İskeletinin Oluşturulması
**Başlık:** `[CHORE] Go Proje Yapısı ve Clean Architecture Kurulumu`
**Açıklama:**
Projenin temel klasör yapısının (Clean Architecture'a uygun olarak) oluşturulması gerekiyor. Bağımlılık yönetimi için `go mod init` çalıştırılmalı.
**Kabul Kriterleri:**
- [ ] `cmd`, `internal`, `pkg`, `api` klasörleri oluşturulmalı.
- [ ] `internal` altında `domain`, `usecase`, `repository`, `delivery` paketleri tanımlanmalı.
- [ ] `go.mod` dosyası oluşturulmalı.

## 2️⃣ Issue: Konfigürasyon ve Veritabanı Bağlantısı
**Başlık:** `[FEAT] Viper ile Config Yönetimi ve PostgreSQL Bağlantısı`
**Açıklama:**
Uygulamanın ayarlarını (port, db host, db user vb.) yönetmek için bir yapı kurulmalı ve veritabanı bağlantısı sağlanmalı.
**Kabul Kriterleri:**
- [ ] `.env` dosyası okunabilmeli.
- [ ] PostgreSQL bağlantı fonksiyonu yazılmalı (Gorm veya sqlx ile).
- [ ] Bağlantı hatası varsa uygulama başlamamalı.

## 3️⃣ Issue: Kullanıcı Kayıt ve Giriş (Auth)
**Başlık:** `[FEAT] User Registration ve Login Endpointleri`
**Açıklama:**
Kullanıcıların sisteme kayıt olabilmesi ve giriş yapabilmesi gerekiyor. Şifreler hashlenerek saklanmalı.
**Kabul Kriterleri:**
- [ ] `POST /register`: Ad, email, şifre ile kayıt.
- [ ] `POST /login`: Email ve şifre ile giriş.
- [ ] Şifreler `bcrypt` ile hashlenmeli.

## 4️⃣ Issue: JWT Middleware Entegrasyonu
**Başlık:** `[SEC] JWT Authentication Middleware Geliştirilmesi`
**Açıklama:**
Task işlemleri gibi korunması gereken endpoint'ler için JWT (JSON Web Token) doğrulama mekanizması gereklidir.
**Kabul Kriterleri:**
- [ ] Login olduğunda Token dönülmeli.
- [ ] Token süresi (örn: 1 saat) olmalı.
- [ ] Korunan rotalara tokensiz istek atıldığında `401 Unauthorized` dönülmeli.

## 5️⃣ Issue: Docker Ortamının Hazırlanması
**Başlık:** `[DEVOPS] Dockerfile ve Docker Compose Hazırlığı`
**Açıklama:**
Uygulamanın her ortamda aynı şekilde çalışması için konteyner yapısına geçilmeli.
**Kabul Kriterleri:**
- [ ] Go uygulaması için `Dockerfile` (Multi-stage build).
- [ ] PostgreSQL ve App servisi için `docker-compose.yml`.
- [ ] `docker-compose up` denildiğinde proje ve DB ayağa kalkmalı.
