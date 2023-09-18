**Todo App**

1. Pendekatan yang saya gunakan ketika mengembangkan aplikasi adalah sesuai dengan skala aplikasi tersebut. Jika Aplikasi sederhana/monolitik saya menggunakan arsitektur MVC (Model View Controller) sehingga proses development bisa lebih cepat. Akan tetapi untuk skala aplikasi yang lebih besar saya menggunakan pendekatan dengan Arsitektur Domain Driven Design. Metode ini membuat pengembangan aplikasi lebih lama untuk satu fungsi yang sederhana, akan tetapi dia lebih fleksibel secara skalabilitas.
2. Kode untuk akses database

   ```
   func newDB(cfg dbConfig) (*gorm.DB, error) {
   // Split host and port.
   split := strings.Split(cfg.Address, ":")
   if len(split) != 2 {
   	return nil, errors.ErrInvalidDBFormat
   }

   // Prepare dsn and open connection.
   dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Address, cfg.Name)
   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
   	AllowGlobalUpdate: true,
   	Logger:            logger.Default.LogMode(logger.Silent),
   	NamingStrategy: schema.NamingStrategy{
   		SingularTable: true,
   	},
   })
   if err != nil {
   	return nil, err
   }

   tmp, err := db.DB()
   if err != nil {
   	return nil, err
   }

   // Set basic config.
   tmp.SetMaxIdleConns(cfg.MaxConnIdle)
   tmp.SetMaxOpenConns(cfg.MaxConnOpen)
   tmp.SetConnMaxLifetime(time.Duration(cfg.MaxConnLifetime) * time.Second)

   database.NewGORM(cfg.Name)
   return db, nil
   }
   ```

3. Abstraksi error bisa digunakan menggunakan pustaka error message sehingga pesan error bisa seragam baik kode error maupun keterangan. Saya sering menggunakan struktur ini ketika handling error
   ```
   //
   {
   	Code: 500,
   	Message: "Bad Request",
   	Constraint: [{
   	}]
   }
   ```
   Code dan message adalah pesan standart dari library error message sedangkan contraint adalah array yang berisi pesan error dari sistem. Constrain bisa berisi lebih dari satu pesan error
4. Seperti dijelaskan di point pertama. Untuk abstraksi proyek app yang kompleks/besar saya menggunakan pendekatan `Domain Driven Design`. Jadi `bussines Logic` ditaruh di `domain`. Sedangkan `Application Logic` di taruh di `service`. Selanjutnya Untuk `logika database` ditaruh di `repository`. Setelah ketiganya tersedia, aplikasi bisa diakses melalui `layer delivery`. `Delivery` ini yang menghubungkan layer aplikasi dengan entitas di luar sistem. `Delivery` bisa berbentuk `REST`, `GRPC`, `CRON`, dll.

Bersama penjelasan ini saya buat struktur aplikasi TODO dengan menggunakan domain driven design based on past project
