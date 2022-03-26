package data

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"time"
)

//Параметры подключения к БД

var Db *sql.DB
var dbUser string
var dbUserPassword string
var dbName string

//DTOшки

//BlogPost Структура для хранения информации о посте
type BlogPost struct {
	Id        int64
	Title     string
	Content   string
	Tags      []*Tag
	CreatedOn time.Time
	Author    string
}

//Tag Структура для хранения информации о тэге поста
type Tag struct {
	Id   string
	Name string
}

func init() {
	var err error

	//получим параматры подключения к БД из переменных окружения

	dbUser = os.Getenv("pguser")
	dbUserPassword = os.Getenv("pguserpassword")
	dbName = os.Getenv("pgname")
	connstring := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbUserPassword)

	//Подключимся к БД

	Db, err = sql.Open("postgres", connstring)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
}

//Методы CRUD для поста

//GetBlogPosts метод для получения всех постов из БД
func GetBlogPosts() (posts []BlogPost, err error) {

	rows, err := Db.Query("SELECT id, title, content, author, createdon FROM public.posts")
	if err != nil {
		return
	}

	for rows.Next() {
		post := BlogPost{}
		err = rows.Scan(&post.Id, &post.Title, &post.Content, &post.Author, &post.CreatedOn)
		if err != nil {
			return nil, err
		}

		tagrows, err := Db.Query("SELECT t.id, t.name FROM public.tags t JOIN public.i_post_tags i "+
			"ON t.id = i.id_tag WHERE i.id_post = $1", post.Id)

		if err != nil {
			return nil, err
		}

		tags := make([]*Tag, 0)
		for tagrows.Next() {
			tag := Tag{}
			err = tagrows.Scan(&tag.Id, &tag.Name)
			if err != nil {
				return nil, err
			}
			tags = append(tags, &tag)
		}

		err = tagrows.Close()
		if err != nil {
			return nil, err
		}

		post.Tags = tags

		posts = append(posts, post)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return
}

func GetBlogPostByID(id int) (post BlogPost, err error) {

	row := Db.QueryRow("SELECT id, title, content, author, createdon FROM public.posts WHERE id = $1", id)
	err = row.Scan(&post.Id, &post.Title, &post.Content, &post.Author, &post.CreatedOn)
	if err != nil {
		return
	}

	tagrows, err := Db.Query("SELECT t.id, t.name FROM public.tags t JOIN public.i_post_tags i "+
		"ON t.id = i.id_tag WHERE i.id_post = $1", post.Id)

	if err != nil {
		return BlogPost{}, err
	}

	tags := make([]*Tag, 0)
	for tagrows.Next() {
		tag := Tag{}
		err = tagrows.Scan(&tag.Id, &tag.Name)
		if err != nil {
			return BlogPost{}, err
		}
		tags = append(tags, &tag)
	}

	err = tagrows.Close()
	if err != nil {
		return BlogPost{}, err
	}

	post.Tags = tags

	return
}

func (post *BlogPost) Save() (err error) {
	if post.Id != 0 {
		err = fmt.Errorf("пост уже сохранен в БД")
	}

	tx, err := Db.BeginTx(context.Background(), nil)

	//Сохраним теги
	for _, tag := range post.Tags {
		err = tag.Save()
		if err != nil {
			terr := tx.Rollback()
			if terr != nil {
				return terr
			}
			return err
		}
	}

	//Сохраним пост
	post.CreatedOn = time.Now()
	row := Db.QueryRow("insert into posts(title, content, author, createdon) values ($1,$2,$3,$4) returning id",
		post.Title, post.Content, post.Author, post.CreatedOn)

	err = row.Scan(&post.Id)
	if err != nil {
		terr := tx.Rollback()
		if terr != nil {
			return terr
		}
		return err
	}

	//Сохраним связь между постом и тегами
	for _, tag := range post.Tags {
		var i int
		row = Db.QueryRow("insert into i_post_tags(id_post, id_tag) values($1, $2) returning id", post.Id, tag.Id)
		err = row.Scan(&i)
		if err != nil {
			terr := tx.Rollback()
			if terr != nil {
				return terr
			}
			return err
		}
	}
	terr := tx.Commit()
	if terr != nil {
		return terr
	}
	return
}

//Методы CRUD для тега

//Save метод сохранения тега. Идемпотентный.
func (t *Tag) Save() error {
	row := Db.QueryRow("select id from public.tags where name = $1", t.Name)
	err := row.Scan(&t.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			insertrow := Db.QueryRow("insert into public.tags (name) values ($1) returning id", t.Name)
			err = insertrow.Scan(&t.Id)
			if err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	}
	return nil
}
