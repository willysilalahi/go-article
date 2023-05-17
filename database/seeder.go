package database

import (
	"fmt"
	"go-article/model"

	"gorm.io/gorm"
)

func Seeding(db *gorm.DB) {
	// Truncate
	db.Exec("DELETE FROM post_tags")
	db.Exec("DELETE FROM comments")
	db.Exec("DELETE FROM posts")
	db.Exec("DELETE FROM tags")
	db.Exec("DELETE FROM admins")

	//Admin
	var admins = []model.Admin{
		{
			ID:    1,
			Name:  "Rony Saragih",
			Email: "rony@taspen.com",
		},
		{
			ID:    2,
			Name:  "Julia Huang",
			Email: "julia@bumn.com",
		},
		{
			ID:    3,
			Name:  "Dian Nugraha",
			Email: "dian@phub.com",
		},
	}
	db.Create(&admins)

	//Tag
	var tags = []model.Tag{
		{
			ID:   1,
			Name: "Politik",
		},
		{
			ID:   2,
			Name: "Keuangan",
		},
		{
			ID:   3,
			Name: "Teknologi",
		},
		{
			ID:   4,
			Name: "Olahraga",
		},
		{
			ID:   5,
			Name: "Pendidikan",
		},
	}
	db.Create(&tags)

	//Post
	var posts = []model.Post{
		{
			ID:          1,
			Title:       "Pengajian Sering Dibubarkan GP Ansor, Ustaz Hanan Gabung NU",
			Description: "Pengajian Ustaz Hanan Attaki di beberapa tempat di Provinsi Jawa Timur, termasuk Madura kerap kali dibubarkan dan diadang oleh pengurus GP Ansor. Alasannya, ustaz alumnus Universitas Al Azhar, Kairo, Mesir ini dianggap menjadi bagian Hizbut Tahrir Indonesia (HTI). Alhasil, sudah beberapa kali, agenda pengajiannya di beberapa tempat dibatalkan dan didatangi hingga urung terlaksana.",
			Status:      1,
			AdminID:     3,
		},
		{
			ID:          2,
			Title:       "5 Cara Memuaskan Istri di Ranjang, Suami Wajib Tahu!",
			Description: "Mengetahui cara memuaskan seorang wanita melibatkan mengetahui apa yang disukainya. Setiap wanita akan memiliki preferensi yang berbeda. Jika Anda ingin memuaskan seorang wanita di tempat tidur secara seksual, kamu mungkin harus membuat beberapa perubahan dalam prosesnya.",
			Status:      0,
			AdminID:     1,
		},
	}
	db.Create(&posts)

	//Comments
	var comments = []model.Comment{
		{
			ID:      1,
			Comment: "Kurang meyakinkan, ga betul ini beritanya",
			AdminID: 1,
			PostID:  1,
		},
		{
			ID:      2,
			Comment: "Kami tidak bertanggung jawab untuk meyakinkan saudara",
			AdminID: 3,
			PostID:  1,
		},
		{
			ID:      3,
			Comment: "Wkkwwk sampahh!",
			AdminID: 1,
			PostID:  1,
		},
		{
			ID:      4,
			Comment: "Mantap mas, beritanya sangat bermanfaat",
			AdminID: 2,
			PostID:  2,
		},
		{
			ID:      5,
			Comment: "Tq y",
			AdminID: 1,
			PostID:  2,
		},
	}
	db.Create(&comments)

	//PostTags
	var post_tags = []model.PostTag{
		{
			PostID: 1,
			TagID:  3,
		},
		{
			PostID: 1,
			TagID:  5,
		},
		{
			PostID: 1,
			TagID:  1,
		},
		{
			PostID: 2,
			TagID:  1,
		},
	}
	db.Create(&post_tags)

	fmt.Println("Sedeed successfully")
}
