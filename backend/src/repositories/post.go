package repositories

import (
	"backend/src/models"
	"database/sql"
)

type posts struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *posts {
	return &posts{db}
}

func (repo posts) Create(post models.Post) (uint64, error) {
	statement, err := repo.db.Prepare("insert into posts (title, content, author_id) values ($1, $2, $3) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var id int
	err = statement.QueryRow(post.Title, post.Content, post.AuthorId).Scan(&id)
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (repo posts) FindById(ID uint64) (models.Post, error) {
	lines, err := repo.db.Query(`
		SELECT p.*, u.nick
		FROM posts p
		INNER JOIN users u ON (u.id = p.author_id)
		WHERE p.id = $1`,
		ID,
	)

	if err != nil {
		return models.Post{}, err
	}

	defer lines.Close()

	var post models.Post

	if lines.Next() {
		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

func (repo posts) FindPosts(userId uint64) ([]models.Post, error) {
	lines, err := repo.db.Query(`
		SELECT DISTINCT p.*, u.nick
		FROM posts p
		INNER JOIN users u ON (u.id = p.author_id)
		LEFT JOIN followers f ON (p.author_id = f.user_id)
		WHERE u.id = $1
		OR f.follower_id = $1
		ORDER BY p.created_at desc`,
		userId,
	)

	if err != nil {
		return []models.Post{}, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repo posts) Update(postID uint64, post models.Post) error {
	statement, err := repo.db.Prepare(`
		UPDATE posts 
		SET title = $1, content = $2
		WHERE id = $3
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Query(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

func (repo posts) Delete(postID uint64) error {
	statement, err := repo.db.Prepare(`
		DELETE 
		FROM posts 
		WHERE id = $1
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Query(postID); err != nil {
		return err
	}

	return nil

}

func (repo posts) FindUserPosts(userID uint64) ([]models.Post, error) {
	lines, err := repo.db.Query(`
		SELECT p.*, u.nick
		FROM posts p
		INNER JOIN users u ON (u.id = p.author_id)
		WHERE u.id = $1
		ORDER BY p.created_at desc`,
		userID,
	)

	if err != nil {
		return []models.Post{}, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repo posts) LikePost(postID uint64) error {
	statement, err := repo.db.Prepare(`
		UPDATE posts 
		SET likes = likes + 1
		WHERE id = $1
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Query(postID); err != nil {
		return err
	}

	return nil
}

func (repo posts) DislikePost(postID uint64) error {
	statement, err := repo.db.Prepare(`
		UPDATE posts 
		SET likes = CASE 
			WHEN likes > 0 THEN
				likes - 1
			ELSE 0
		END
		WHERE id = $1
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Query(postID); err != nil {
		return err
	}

	return nil
}
