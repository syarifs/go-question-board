# Question Board
Aplikasi REST API untuk membuat survey, polling, atau melakukan evaluasi guru di sekolah. dibuat untuk memenuhi tugas mini project dalam program Kampus Merdeka bersama [alta.id](https://alta.id).

## Feature
##### Main Feature
- Email Notification.
- Evaluate Teacher
- Filtered Survey and Polling.

##### Common Feature
- Authenication.
- Evaluate Teacher.
- Create Survey or Polling.
- View Survey or Polling Status.

## Feature In Development
- [x] Authentication
	- [x] User Login
	- [x] Generate JSON Web Token
	- [ ] Refresh JSON Web Token
- [x] Basic CRUD
	- [x] Tags
	- [x] Subject
	- [x] Major
	- [x] User
	- [x] Questionnaire
- [x] Subject
	- [x] Register Teacher to Subject
- [x] Questionnaire
	- [x] Create Survey
	- [x] View Survey Response
	- [x] Answer Survey
	- [x] Filter By Tags
	- [x] Filter Completed Survey
- [x] Teacher Evaluation
	- [x] Evaluate Teacher by Subject
	- [x] View Evaluated Data from Teacher Side
- [ ] Profile Management
	- [ ] Read & Update Profile
- [ ] Email Notification
	- [ ] Notify tagged user about survey available

## Todo
- [x] ERD.
- [x] Base Project.
- [x] Menambahkan CRUD untuk manajemen `User`.
- [x] Menggunakan file `config.yaml` untuk memuat *environment variables*.
- [x] Membuat Database Seeder.
- [x] CRUD One to One Relation.
- [x] CRUD Many to Many Relation.
- [x] Menambahkan Login Path.
- [x] Implementasi `JWT Middleware`.
- [x] Implementasi `Logging Middleware`.
- [x] Menambahkan CRUD untuk manajemen `Subject`.
- [x] Menambahkan CRUD untuk manajemen `Tag`.
- [x] Menambahkan CRUD untuk manajemen `Major`.
- [x] Membuat proteksi role untuk route admin.
- [x] Membuat dokumentasi untuk setiap *route* yang dibuat.
- [x] Membuat `Unit Test` untuk modul *interface*.
- [x] Menambahkan CRUD untuk manajemen `Questionnaire`.
- [x] Menambahkan fitur `Filter by Tag` untuk `Questionnaire`.
- [x] Implementasi `Password Hashing` agar password lebih aman.
- [x] Menambahkan fitur `Answer Questionnaire`.
- [x] Mendaftarkan user dengan role `Teacher` ke `Subject`.
- [x] Menampilkan `Teacher` pada *list* `Subject`.
- [x] Membuat *route list* `Subject` yang diikuti `Student`.
- [x] Implementasi fitur `Evaluate Teacher`.
- [ ] Membuat fitur `Management Profile` untuk `User`.
- [ ] Membuat `Unit Test` untuk modul *controller*.
- [ ] Menambahkan fitur `Email Notification` untuk pemberitahuan `Questionnaire` yang dapat dikerjakan.
- [ ] Menggunakan MongoDB untuk filter JSON Web Token.
- [ ] Menggunakan MongoDB untuk *logging* sistem.
- [ ] Deploy to Docker.
- [ ] Deploy to Amazon Web Service.
