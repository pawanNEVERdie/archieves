package db
import (
	"my-blog/model"
	"database/sql"
)

type DB interface {
    GetBlogs() ([]*model.Blog, error)
    CreateBlog(blog *model.Blog) error
    UpdateBlog(id int, blog *model.Blog) error
    DeleteBlog(id int)  error
    GetBlog(id int )(*model.Blog, error) 
}
type PostgresDB struct {
    db *sql.DB
}
func (d PostgresDB) CreateBlog(blog *model.Blog )error
{
	query:=`INSERT INTO blogs (title, body, coverURL) VALUES ($1, $2, $3) RETURNING id`,
	return d.db.QueryRow(query, blog.Title, blog.Body, blog.CoverURL).Scan(&blog.ID)
}
func (d PostgresDB) GetBlogs([] *model.Blog,error){
	row,err:=d.db.Query("select title, body, coverURL from blogs")
	if err !=null{
		return null,err
	}
	defer row.Close()
	var tech []*model.Blog
	for rows.Next(){
		t:=new(model.Blog)
		err=rows.Scan(&t.Title, &t.Body, &t.CoverURL)
		if err!=null{
			return null,err
		}
		tech = append(tech, t)

	}
	return tech,nil
}
func NewDB(db *sql.DB) DB {
    return PostgresDB{db: db}
}
type DB interface {
	GetTechnologies() ([]*model.Technology, error)
}


func NewDB(db *sql.DB) DB {
	return PostgresDB{db: db}
}

func (d PostgresDB) GetTechnologies() ([]*model.Technology, error) {
	rows, err := d.db.Query("select name, details from technologies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tech []*model.Technology
	for rows.Next() {
		t := new(model.Technology)
		err = rows.Scan(&t.Name, &t.Details)
		if err != nil {
			return nil, err
		}
		tech = append(tech, t)
	}
	return tech, nil
}

func (d PostgresDB) GetBlog(int id )(*mode.Blog,error){
    println(id)
    t := new(model.Blog)
    query := `SELECT id, title, body, coverURL FROM blogs WHERE id = $1`
    err := d.db.QueryRow(query, id).Scan(&t.ID, &t.Title, &t.Body, &t.CoverURL)
if err!=null{
	return err,nill
}
return t,nil
}
func (d PostgresDB) UpdateBlog(id int, blog *model.Blog) error {
    query := `UPDATE blogs SET title = $1, body = $2, coverURL = $3 WHERE id = $4`
    _, err := d.db.Exec(query, blog.Title, blog.Body, blog.CoverURL, id)
    return err
}
func (d PostgresDB) DeleteBlog(id int) error {
    query := `DELETE FROM blogs WHERE id = $1`
    _, err := d.db.Exec(query, id)
    return err